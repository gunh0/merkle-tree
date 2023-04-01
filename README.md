# Merkle Tree Implementation in Go

This is a simple implementation of a Merkle tree in Go, based on the SHA-256 hash function.

<br/>

### Overview

The Merkle tree is a tree structure used in cryptography and computer science to efficiently verify the contents of large data sets. It allows the efficient and secure verification of the contents of a large data set by dividing it into smaller chunks and creating a hash tree out of them.

In this implementation, the tree is created by first hashing all the data elements using SHA-256, and then building the tree recursively by hashing pairs of nodes until a single root node is produced. Each non-leaf node in the tree is the hash of its two child nodes, and the root node is the hash of all the data elements.

### Usage

The main function in the main.go file provides an example of how to use this implementation. It creates a Merkle tree from a slice of byte arrays, and then prints out the entire tree using the printTree function.

To use this implementation in your own code, you can import the merkle package and use the NewTree function to create a Merkle tree from a slice of byte arrays. The NewTree function returns the root node of the tree, which you can then use to traverse the tree and verify the contents of the data set.

### Dependencies

This implementation uses the crypto/sha256 package to compute SHA-256 hashes.
