---
sidebar_position: 1
---

# Introduction

**DocChain** started with the sole purpose of easing the trouble of keeping hard records of documents.  
**Documents** can be anything, from an agreement to a birth certificate.
DocChain aims to prove originality to our digitalized documents using the blockchain technology.

Understand more about **[Blockchain Technology](blockchain.md)**

The way we seek to solve the problem is to create immutable **Changes**(Transactions) to the blockchain.
These changes will contain documents that are already examined and proved to be legit.

These changes go through many inspections like signing and verifying.
To make them truly immutable the blocks are chained by adding the hash of the previous block to the current block.
The block is further mined as the **[PoW](https://en.wikipedia.org/wiki/Proof_of_work)**.
We plan to make it so that the stamp value is paid by computer resources to mine the block.

DocChain is being implemented in **[Go](https://go.dev/)**.
This Go implementation will try to abstract the implementation into many levels.
These individual levels will be available for other blockchain based technologies to rely on.
Currently, only the very basic base implementation of the blockchain and related features has been started.
This documentation tries to explain this implementation thoroughly.
