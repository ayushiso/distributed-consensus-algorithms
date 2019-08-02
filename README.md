# Consensus Algorithms
Go implementations of basic distributed systems consensus algorithms.

## Synchronized agreement algorithm (Crash fault tolerant)
Movie agreement algorithm from Chapter 4 (p. 23) in [Distributed Computing Pearls](http://www.faculty.idc.ac.il/gadi/DCPbook.htm) (Taubenfeld)

## Synchronized agreement algorithm (Byzantine fault tolerant)
Byzantine agreement algorithm from Chapter 5 (p. 33) in [Distributed Computing Pearls](http://www.faculty.idc.ac.il/gadi/DCPbook.htm) (Taubenfeld)

## Ben-Or's randomized protocol for asynchronous message-passing systems (Crash fault tolerant)
A probabilistic workaraound for the FLP impossibility, which guarantees that as long as a majority of the processes continues to operate, a decision will be made. It is guaranteed to work with probability 1 even against an adversary scheduler who knows all about the system. [Original paper](https://allquantor.at/blockchainbib/pdf/ben1983another.pdf) (Ben-Or, 1983). Pseudocode used from [Aspnes](https://arxiv.org/pdf/cs/0209014.pdf) (2018).
