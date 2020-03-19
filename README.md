# forex-grpc

## Usage:
    $ ./install-dependencies.sh
    $ go run server/main.go
    $ go run client/main.go <source> <target> <amount>

## client variables
    - source: three letter currency symbol from senders account (e.g. "USD")
    - target: three letter currency symbol from receivers account (e.g. "EUR")
    - amount: floating point number of amount in senders currency
