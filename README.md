# Goknock
A very basic, extensible and insecure TCP port knocker.
The purpose was only to learn Go.

## Features

### Extensible
Want more actions than the (very) basic ones defined? Easy!
Just implement the `Action` interface (defined in `actions.go`), add a sequence with that action to the `sequences` array in `main.go`, and ta-da!

### Multiclient
Go's goroutines enable this cute port knocker to support multiclient out of the box.

### Timeouts
You can define a timeout until the next port knock when defining each sequence.

### Relatively discreet
Sequences can be set to silent and the code won't print any message.
In addition, go binaries aren't very big so it's convinent to move around.

## FAQ

### Why is this insecure?
Yes, ports will be opened one after the other in sequence (supposing the sequence is valid and no timeout occurs). This means it's discoverable (like most port knockers anyway).
This also means that a firewall may prevent port opening and an IDS may easily notice something out of the ordinary happening.

### Enhancements and bugs
Make an issue, a pull request, or if/when I abandon this go demo, fork this repo (and make something a bit better).
