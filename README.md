# Merkle Tree Implementation in Go

This is a simple implementation of a Merkle tree in Go, based on the SHA-256 hash function.

<br/>

### Overview

The Merkle tree is a tree structure used in cryptography and computer science to efficiently verify the contents of large data sets. It allows the efficient and secure verification of the contents of a large data set by dividing it into smaller chunks and creating a hash tree out of them.

In this implementation, the tree is created by first hashing all the data elements using SHA-256, and then building the tree recursively by hashing pairs of nodes until a single root node is produced. Each non-leaf node in the tree is the hash of its two child nodes, and the root node is the hash of all the data elements.

### Usage

The main function in the main.go file provides an example of how to use this implementation. It creates a Merkle tree from a slice of byte arrays, and then prints out the entire tree using the printTree function.

To use this implementation in your own code, you can import the merkle package and use the NewTree function to create a Merkle tree from a slice of byte arrays. The NewTree function returns the root node of the tree, which you can then use to traverse the tree and verify the contents of the data set.

<br/>

### Output

```
go run main.go
└── 2846362fb3dce03001286fe9a997206a0ade591d5bddbfed0a7a9901bac798cc
    ├── 7305db9b2abccd706c256db3d97e5ff48d677cfe4d3a5904afb7da0e3950e1e2
    │   ├── 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
    │   └── 486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7
    └── 487e8e3fb58ea5fc6855763fe7a918bda75f564dd0649d8c6b7aefb6f23bd094
        ├── 7975edd9e7393c229e744913fe0d0bb86fb4cf46906e2e51152137e20ad15590
        └── dc9c5edb8b2d479e697b4b0b8ab874f32b325138598ce9e7b759eb8292110622
```

<br/>

### Dependencies

This implementation uses the crypto/sha256 package to compute SHA-256 hashes.
