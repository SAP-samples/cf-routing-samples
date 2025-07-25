package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/netip"
	"net/url"
	"os"
	"strings"
)

var (
	ErrBadRequest                   = fmt.Errorf("bad request")
	ErrForbidden                    = fmt.Errorf("forbidden")
	ErrMissingOrInvalidForwardedURL = fmt.Errorf("%w: invalid or missing x-cf-forwarded-url header", ErrBadRequest)
	ErrMissingTrueClientIP          = fmt.Errorf("%w: missing x-cf-true-client-ip header", ErrBadRequest)
	ErrInvalidTrueClientIP          = fmt.Errorf("%w: invalid x-cf-true-client-ip header", ErrBadRequest)
)

//go:embed allowlist.txt
var allowListFile []byte

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))

	err := Main()
	if err != nil {
		slog.Error("route-service failed", "error", err)
		os.Exit(1)
	}
}

func Main() error {
	allowedPrefixes, err := loadPrefixes()
	if err != nil {
		return err
	}

	s := http.Server{
		Addr: ":" + os.Getenv("PORT"),
		Handler: &httputil.ReverseProxy{
			Director:     proxyDirector,
			ErrorHandler: proxyErrorHandler,
			Transport: &transport{
				allowedPrefixes: allowedPrefixes,
				roundTripper:    &http.Transport{},
			},
		},
	}

	slog.Info("starting server", "allow-list", allowedPrefixes, "addr", s.Addr)

	return s.ListenAndServe()
}

func loadPrefixes() (prefixes []netip.Prefix, err error) {
	s := bufio.NewScanner(bytes.NewReader(allowListFile))
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 || l[0] == '#' {
			continue
		}

		p, err := netip.ParsePrefix(l)
		if err != nil {
			return nil, err
		}

		prefixes = append(prefixes, p)
	}

	if s.Err() != nil {
		return nil, s.Err()
	}

	return prefixes, nil
}

func proxyDirector(req *http.Request) {
	log := slog.With("component", "proxy-director", "vcap-id", req.Header.Get("x-vcap-request-id"))

	log.Info("handling request", "client-ip", req.RemoteAddr)

	forwardedURL := req.Header.Get("x-cf-forwarded-url")
	if forwardedURL == "" {
		log.Error("missing x-cf-forwarded-url header")
		req.URL = nil
		req.Host = ""
		return // the transport will deal with that
	}

	u, err := url.Parse(forwardedURL)
	if err != nil {
		log.Error("invalid x-cf-forwarded-url header", "error", err)
		req.URL = nil
		req.Host = ""
		return // the transport will deal with that
	}

	log.Debug("obtained forwarded URL", "forwarded-url", u.String())

	req.URL = u
	req.Host = u.Host
}

func proxyErrorHandler(w http.ResponseWriter, r *http.Request, handleErr error) {
	log := slog.With("component", "proxy-error-handler", "vcap-id", r.Header.Get("x-vcap-request-id"))

	if handleErr == nil {
		panic("received nil error")
	}

	var status int
	switch {
	case errors.Is(handleErr, ErrBadRequest):
		status = http.StatusBadRequest
	case errors.Is(handleErr, ErrForbidden):
		status = http.StatusForbidden
	default:
		status = http.StatusBadGateway
	}

	log.Warn("handling error", "error", handleErr, "status", status)

	w.WriteHeader(status)
	_, err := fmt.Fprintf(w, "error: %s", handleErr.Error())
	if err != nil {
		log.Warn("failed to send handleError to client", "handleError", handleErr, "error", err, "status", status)
	}
}

type transport struct {
	allowedPrefixes []netip.Prefix
	roundTripper    http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// This indicates that an issue was encountered while trying to read the forwarded URL.
	if req.URL == nil {
		return nil, ErrMissingOrInvalidForwardedURL
	}

	addrString := req.Header.Get("x-cf-true-client-ip")
	if addrString == "" {
		return nil, ErrMissingTrueClientIP
	}

	addr, err := netip.ParseAddr(addrString)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidTrueClientIP, err)
	}

	// This is a bit inefficient but for a good enough for a sample. A proper implementation should
	// use some form of trie for efficient lookups.
	allowed := false
	for _, p := range t.allowedPrefixes {
		if p.Contains(addr) {
			allowed = true
			break
		}
	}

	if !allowed {
		return nil, fmt.Errorf("%w: address '%s' is not in allow-list", ErrForbidden, addr.String())
	}

	return t.roundTripper.RoundTrip(req)
}
