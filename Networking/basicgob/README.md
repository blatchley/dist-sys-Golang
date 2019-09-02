#Using gob to send strings
This is a very basic serverm which uses the gob package to decode strings sent to it by connected clients.

To run, enter "go run basicserver.go" into the terminal, or build then run it. It will then prompt for input, either enter "new" to create a "server", or the (ip:port) of the "server" you wish to connect to.

Clients can type text, which they encode into the connection. The server runs a loop on each handled connection, decoding each string and printing it to terminal.

Gob is relatively simple, but also pretty limited. One decode can only work on a single data type or data structure, and if it ever receives data which it cannot deserialise into that data type, it will error out/crash. Furthermore, we found that at high loads with large structures, weird network buffer collisions started happening, causing deserialisation errors.

If you need to send multiple different structures around you will need to either have multiple different decoders depending on type, (and either run these on different connections, or have some RELIABLE way of knowing what is coming before attempting to decode it,) or use another tool. For our blockchain bachelor project, we implemented our P2P layer using RPC. RPC is essentially a wrapper around gob which solves many of these issues, and will be covered later.

This is essentially a one way chat client, and could be extended to function as a two way client by just having both ends run a loop accepting text inputs, and sending those over the connection. 

Other extensions like Multiparty chat rooms could be supported by, for example, maintaining a list of connections at each node, and every time text is sent, sending it to every connection in the list. Or distributed systems without a central host, by having every node allow others to connect to them, and all flooding messages they haven't seen before to all other nodes. (This could require running the "listen for connections" in it's own go routine, leaving the main thread free to accept text inputs.)