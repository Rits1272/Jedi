# Jedi
A distributed, scalable, and consistent cache store written in Go

This is a self-learning project teaching myself how to design a distributed cache with the following salient:
- replication (chain replication?)
- consistency (quorum, master/slave nodes)
- high scalability (consistent hashing, virtual nodes)
- Jedi Client to interact with the cache store
- Zookeeper to manage nodes

## Goal
- The goal of the project would be to design a cache store with above features
- Possibly would write a script to spin up docker containers locally which will have node containers, replica containers, client containers and a zookeeper container
