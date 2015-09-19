# Domain Router

An example Nginx router setup, using Docker, complete with integration tests and resilience testing using [Muxy](https://github.com/mefellows/muxy).

## TOC

* [Setup](#setup)
* [Testing](#testing)

## Setup

### Install & Setup Docker

* Install [Docker](http://docker.io/) and Docker Machine
* Build and run the image

```
docker-compose build && docker-compose up
```

This should spin up Nginx, DNSMasq and a test API (echo server). You can now hit the API at `curl -v -H"Host: api.foo.com" <dockerhost>:8001/header/host`.
Note that we have to pass through the host header with the request.

# Testing

Integration test suite for Nginx.

## Running Tests

```
./scripts/run-tests.sh
```

## Test Setup

To be able to run the Nginx router locally, and also to be able to test the configuration, we use [dnsmasq](http://www.thekelleys.org.uk/dnsmasq/doc.html) in conjunction with Docker Compose.

Nginx is then configured to use dnsmasq for [host resolution](http://nginx.org/en/docs/http/ngx_http_core_module.html#resolver), allowing us to replace real systems with a [Mock Service](../mockapi/README.md) by creating host entries in the dnsmasq container. This works nicely in Production, where we don't create any host entries so the real services are resolved instead.

The components of the integration test are:

* Nginx
* dnsmasq - hijacks the real domain names, instructing Nginx to send traffic to our Mock Server so that we can
* [Mock Server](../mockapi/README.md)
* Test Case container issuing requests against the Nginx instance

This can be visually represented as:

```
[Test Container] <- (issues tests to) -> [Nginx Container] <- (proxies) -> [Mock API]
                                                 ||
                                    (resolves DNS queries from)
                                                 \/
                                             [Dnsmasq]
```