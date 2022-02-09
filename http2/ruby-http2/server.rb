require 'http/2'
require 'socket'

# Example H2C web app
# Based on https://github.com/igrigorik/http-2/blob/master/example/server.rb

port = ENV.fetch("PORT")
server = TCPServer.new(port)

loop do
  sock = server.accept

  conn = HTTP2::Server.new
  conn.on(:frame) do |bytes|
    sock.is_a?(TCPSocket) ? sock.sendmsg(bytes) : sock.write(bytes)
  end

  conn.on(:stream) do |stream|
    stream.on(:half_close) do
      stream.headers({
        ":status" => "200",
        "content-type" => "text/plain"
      }, end_stream: false)

      stream.data("Hello! This Ruby application is speaking plain text HTTP2 (H2C) with the CF routing layer\n", end_stream: true)
    end
  end

  while !sock.closed? && !(sock.eof? rescue true)
    data = sock.readpartial(1024)

    begin
      conn << data
    rescue StandardError => e
      puts "#{e.class} exception: #{e.message} - closing socket."
      e.backtrace.each { |l| puts "\t" + l }
      sock.close
    end
  end
end
