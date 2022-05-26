# Aeldra

## Installation

### Install Go 1.14 or higher
To install Go: [https://golang.org/doc/install](https://golang.org/doc/install)

Verify your `$GOPATH` is correctly set before continuing!

### Setup this repository

Go is bit picky about where you store your repositories.

The convention is to store:
- the source code inside the `$GOPATH/src`
- the compiled program binaries inside the `$GOPATH/bin`

#### Using Git
```bash
mkdir -p $GOPATH/src
cd $GOPATH/src

Clone the repository

go install ./...
```

## CLI
### Show available commands and flags
```bash
aeldra help
```

#### Show available run settings
```bash
aeldra run --help

Launches the Aeldra node and its HTTP API.

Usage:
  aeldra run [flags]

Flags:
      --datadir string             Absolute path to your node's data dir where the DB will be/is stored
  -h, --help                       help for run
      --ip string                  your node's public IP to communication with other peers (default "127.0.0.1")
      --port uint                  your node's public HTTP port for communication with other peers
```
## HTTP
### List all balances
```
curl http://localhost:8080/balances/list | jq
```
### Send a signed TX
```
curl --location --request POST 'http://localhost:8080/tx/add' \
--header 'Content-Type: application/json' \
--data-raw '{
	"from": "eshwar",
	"to": "rajib",
	"value": 100
}'
```

### Check node's status (latest block, known peers
```
curl http://localhost:8080/node/status | jq
```

### Fetch node's peers
```
curl --location --request GET 'http://localhost:8081/node/list'
```

### Get Lastest Block
```
curl --location --request GET 'http://localhost:8080/block/latest'
```
