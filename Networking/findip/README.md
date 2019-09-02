# Finding your own IP
Finding the ip of your computer address on the network in Go can be surprisingly hard. Many of the top results on google/stackoverflow are platform specific, or will return unpredictable depending on what network adapters, vm's, vpn's etc you have on your computer.

The rather hacky code stump given here, found on stackoverflow, is the most reliable way I've found so far of finding an ip address which your machine is currently listening on. And should work regardless of setup. This is especially useful if you're writing code on a different OS than your instructor is using to test it.

It works by "initialising" a udp connection, then checking the outbound IP address on that connection. This works even if there is no server to connect to, or you do not have an internet connection, as the UDP protocol doesn't require a handshake for the connection to be "live."

This is a very useful for cases where programs need to be aware of their own external ip address, and is used several places in these networking examples.

Note that in rare cases, if you select an IP address and port which are specifically blocked for some reason, then the dial method could theoretically return a non null error, causing a fatal on line 14. Personally I've never seen this happen, but in theory it's presumably possible.
