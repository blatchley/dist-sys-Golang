|Apologies for the delay, had a bad case of the flu for the last week and a half :(
|Examples of gob implementations are up, which should already be overkill for the first handin. RPC implementations and architecture should be up soon.

# dist-sys-Golang
Supplementary resources for Go, based on distributed systems and security course.

As there are plenty of resources out there for getting started, I'll be focussing on some common Go architecture and networking patterns which we struggled with during the course last year, as well as working examples of some of the core libraries actually being implemented.

These are definitely not the only ways of solving these problems, so if you have a useful design pattern, code snippet, educational example or anything else you think would be worth adding, feel free to submit a pull request or email me, (201205917 postaudk). Likewise, if you find a bug or error, or feel something very fundamental is missing, please let me know. While I am not an instructor on the course, I will attempt to maintain this repo through the course.

# Useful links for getting started.
installation: https://golang.org/doc/install

Tour through most features, with basic examples: https://tour.golang.org/welcome/1

Documentation of native packages: https://golang.org/pkg/

Plenty of articles on medium, such as https://medium.com/rungo, and decent support on stack overflow.

# Basic Networking

TCP client/server example - [Here](Networking/tcpintro)

Finding your own IP address - [Here](Networking/findip)

Basic data sending using Gob - [Here](Networking/basicgob)

sending structs using Gob - [Here](Networking/structgob)

using net/RPC to send multiple struct types - WIP

# Managing connections
Two way RPC connections - WIP

Fully Connected network using Peer List - WIP

# Architecture

How to arrange packages - WIP

Example net/RPC design using channels - WIP

Concurrency with locks - WIP
