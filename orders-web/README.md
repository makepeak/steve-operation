# Order Service

This is the Order service

Generated with

```
micro new --namespace=steve.user --alias=order --type=web github.com/makepeak/steve-operation/orders-web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: steve.user.web.order
- Type: web
- Alias: order

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
./order-web
```

Build a docker image
```
make docker
```