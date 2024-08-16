# Arkeo Protocol

Arkeo Protocol - Free Market Blockchain Data Infrastructure

[![Arkeo CI](https://github.com/arkeonetwork/arkeo/actions/workflows/ci.yml/badge.svg)](https://github.com/arkeonetwork/arkeo/actions/workflows/ci.yml)
[![Release](https://github.com/arkeonetwork/arkeo/actions/workflows/release.yml/badge.svg)](https://github.com/arkeonetwork/arkeo/actions/workflows/release.yml)


The core implementation of the Arkeo Protocol is built using the Cosmos SDK and Tendermint, and was created with the help of the [Ignite CLI](https://ignite.com/cli).


# Software Prerequisites

- [The Go Programming Language](https://go.dev)
- [Git Distributed Version Control](https://git-scm.com)
- [Docker](https://www.docker.com)
- [GNU Make](https://www.gnu.org/software/make)


# Build from Source

Clone the repository 

```shell
git clone https://github.com/arkeonetwork/arkeo.git
```

Check our `master` branch 

```shell
git checkout master
```

Build 

```shell
make install
```

This will install `arkeod` binary to your `GOPATH`

# Running Localnet

There multiple ways to run a localnet

## Containerized

This solution uses docker-compose and docker on backend, To setup new localnet use:
```shell
make localnet
```

## Ignite CLI 

Install Ignite Cli 
```shell
curl https://get.ignite.com/cli! | bash
```

```shell
ignite chain serve
```

## Local

Build Binary 

```shell
make proto-gen install 
```

Run 
```shell
./scripts/genesis.sh
```
This starts the chain 


# Documentation
To learn more about `Arkeo Protocol`, [please visit the official Arkeo Documentation](https://docs.arkeo.network)4