# forex-grpc
* A gRPC API to compute currency exchange conversions from source currency amount to target currency amount.
    - USD: US dollar
    - EUR: Euro
    - GBP: Great British Pound
    - JPY: Japanese Yen
    - CAD: Canadian Dollar
    - AUD: Australian Dollar
    - CNY: Chinese Yuan Renminbi
    - CHF: Swiss Franc
    - SGD: Singapore Dollar
    - NOK: Norwegian Krone
* Supports listing available currencies (see client usage below)
* The below usage instructions assume you already have go, node.js, and npm installed.

## Usage:
---
* visit https://fixer.io/ to get an API key
* the free subscription will default to using EUR as the base currency
* if you have the basic subscription ($10/month), the server will randomly choose from 10 of the worlds top currencies
* create file conf.yaml in the root directory and add "apikey: \<your-api-key>"
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
