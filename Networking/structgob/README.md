# Using gob to send strings

Example similar to basicgob, just showing how you can send more complex data structures.

Note that while this works for most default types in go, some external package types, such as big.int, or interfaces, are not handled by the default gob marshaller.

Also note that gob only encodes the exported fields in structures.
