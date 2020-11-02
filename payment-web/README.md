# Payment Service

This is the Payment service

Generated with

```
micro new --namespace=steve.user --alias=payment --type=web github.com/makepeak/steve-operation/payment-web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: steve.user.web.payment
- Type: web
- Alias: payment

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./payment-web
```

Build a docker image
```
make docker
```