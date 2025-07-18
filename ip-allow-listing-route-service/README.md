# IP Allow-List Route Service

A simple GoLang application that implements the [route service contract](https://docs.cloudfoundry.org/services/route-services.html)
of Cloud Foundry and checks the client IP address against an allow-list before passing on the
request. This requires the `x-cf-true-client-ip` header to be set by the foundation which is the
case on [SAP BTP, Cloud Foundry runtime](https://www.sap.com/products/technology-platform/btp-cloud-foundry-runtime.html).

## Usage

> [!CAUTION]
>
> This is a sample application that is only intended to demonstrate the general idea. It is not
> production grade.

To deploy this route-service you will first need a separate application you want to put this in
front of, as well as an allow-list. The allow-list should contain one IP-prefix per line, empty
lines and lines starting with a `#` are ignored. Example:

```
# This is a comment
10.0.0.0/8
127.0.0.0/8
# The following lines will be ignored


```

Put this file into the app directory next to the manifest file and name it `allowlist.txt`. Now you
can push the application like this (you may need to adjust the domain depending on the region):

```
cf push -f ./manifest.yml --var domain=cfapps.example.com --var prefix=foo
```

This pushes the app and makes it available at `foo-ip-allow-list-rs.cfapps.example.com`.
Now we can turn it into a route service and bind it to an already existing route, provide the
domain and host of the route you want to restrict access to:

```
cf create-user-provided-service allow-listing -r https://foo-ip-allow-listing-rs.cfapps.example.com
cf bind-route-service cfapps.example.com --hostname my-app allow-listing
```

See the help pages of the individual commands for extended options.

Now every request sent to the route you have bound the route-service to will pass through the
allow-listing route-service and is only forwarded to the target application if the client IP
address is allow-listed.
