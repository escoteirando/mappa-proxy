# MAPPA PROXY

[![CodeQL](https://github.com/escoteirando/mappa-proxy/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/escoteirando/mappa-proxy/actions/workflows/codeql-analysis.yml)
[![Go](https://github.com/escoteirando/mappa-proxy/actions/workflows/go.yml/badge.svg)](https://github.com/escoteirando/mappa-proxy/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/escoteirando/mappa-proxy)](https://goreportcard.com/report/github.com/escoteirando/mappa-proxy)

Access: https://mappa-proxy.fly.dev/

## Links

* https://medium.com/c%C3%B3digo-palavras/como-subir-sua-aplica%C3%A7%C3%A3o-golang-gratuitamente-para-a-internet-57321cfcbaa0
* https://dlintw.github.io/gobyexample/public/memory-and-sizeof.html
* https://dev.to/knowbee/how-to-setup-secure-subdomains-using-nginx-and-certbot-on-a-vps-4m8h

## Frontend Build

Javascript requirements:
* [NodeJS](https://nodejs.org/en/download/)
* [Yarn](https://classic.yarnpkg.com/lang/en/docs/install/#debian-stable)

Run makefile action

```bash
make frontend-setup
```

When some modification on frontend, run the build distribution with the command: 

```bash
make frontend
```