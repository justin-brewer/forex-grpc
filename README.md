# forex-grpc
* A gRPC API to compute currency exchange conversions from source currency amount to target currency amount.
* Supports listing available currencies (see client usage below)
* The below usage instructions assume you already have go, node.js, and npm installed.

## Usage:
---
* visit https://fixer.io/ to get an API key
* create file conf.yaml in the root directory and add "apikey: <your-api-key>"
---
### Server (from project root)
    $ ./install-dependencies.sh    # there is also a powershell version
    $ go run server/main.go
---
### Client
### Golang (from project root)
    $ go run client/main.go <source> <target> <amount>
    $ go run client/main.go list

### node.js (from dir client-js)
    $ npm install
    $ node client.js <source> <target> <amount>
    $ node client.js list

### arguments
* <span style="text-decoration: underline">source</span>: three letter currency symbol from senders account (e.g. "USD")
* <span style="text-decoration: underline">target</span>: three letter currency symbol from receivers account (e.g. "EUR")
* <span style="text-decoration: underline">amount</span>: floating point number of amount in senders currency
---
