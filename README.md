# forex-grpc
* A gRPC API to compute currency exchange conversions from source currency amount to target currency amount.

* The below usage instructions assume you already have go installed.

## Usage:
---
### Server (from project root)
    $ ./install-dependencies.sh    # there is also a powershell version
    $ go run server/main.go
---
### Client
### Golang (from project root)
    $ go run client/main.go <source> <target> <amount>

### node.js (from dir client-js)
    $ npm install
    $ node client.js <source> <target> <amount>

### arguments
* <span style="text-decoration: underline">source</span>: three letter currency symbol from senders account (e.g. "USD")
* <span style="text-decoration: underline">target</span>: three letter currency symbol from receivers account (e.g. "EUR")
* <span style="text-decoration: underline">amount</span>: floating point number of amount in senders currency
---