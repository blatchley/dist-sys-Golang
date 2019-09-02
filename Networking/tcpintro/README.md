# Basic TCP connection

This is a very basic server which accepts incoming connections over TCP.

To run, enter "go run basicserver.go" into the terminal, or build then run it. It will then prompt for input, either enter "new" to create a "server", or the (ip:port) of the "server" you wish to connect to.

It will then connect, the server will note the connection, and the calling client will terminate.

If you cannot connect to a server hosted on your own machine through the ip address the server gives, try connecting with local host (127.0.0.1) as the ip address, with the same port. 

If connecting through localhost works but through the actual IP doesn't, then there is likely an issue with your network settings. Try, for example, setting your windows network profile to private if it is on public, or connecting to a mobile wifi hotspot if you were on eduroam.
