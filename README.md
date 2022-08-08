# Firefly III Exporter

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/kinduff/firefly_iii_exporter?)][releases]
[![Go Reference](https://pkg.go.dev/badge/github.com/kinduff/firefly_iii_exporter.svg?)][godoc]
[![Go Report Card](https://goreportcard.com/badge/github.com/kinduff/firefly_iii_exporter?)][goreport]
[![CodeFactor](https://www.codefactor.io/repository/github/kinduff/firefly_iii_exporter/badge)][codefactor]
[![Test / Build](https://github.com/kinduff/firefly_iii_exporter/actions/workflows/ci.yml/badge.svg?branch=master)][workflow-c]
[![Docker Pulls](https://img.shields.io/docker/pulls/kinduff/firefly_iii_exporter.svg?)][dockerhub]
[![GitHub all releases](https://img.shields.io/github/downloads/kinduff/firefly_iii_exporter/total?)][releases]

This is a Prometheus exporter for [Firefly III](https://www.firefly-iii.org). It provides statistics from accounts, transactions, and categories.

An example of the output of this exporter can be found [here](extra/metrics_example.txt).

## Feature list

This is a work in progress, please feel free to contribute with additional statistics by extending this exporting and submitting a Pull Request.

- [x]  Accounts
  - [x]  Count
  - [x]  Balances
- [x]  Transations
  - [x]  Count
- [x] Categories
  - [x] Count
  - [x] Transactions
- [ ] Bills
- [ ] Piggy banks
- [ ] Budgets
- [ ] Recurrences

## Prerequisites

One of the two, depending on your running method.

* [Go][go] `>= 1.16`
* [Docker][docker]

## Running this exporter

See [Configuration][configuration] in order to set the necessary params to run the exporter.

### Using a binary

You can download the latest version of the binary built for your architecture [here][releases].

### Using Docker

The exporter is also available as a Docker image in [DockerHub][dockerhub] and [Github CR][ghcr]. You can run it using the following example and pass the configuration as environment variables:

```shell
$ docker run \
  --name firefly_iii_exporter \
  -p 7355:7355 \
  -e BASE_URL=<url-to-firefly-iii> \
  -e API_KEY=<personal-access-token> \
  kinduff/firefly_iii_exporter:latest
```

Alternative, you can use `ghcr.io/kinduff/firefly_iii_exporter` if you want to use the Github Container Registry.

### Using the source

Optionally, you can download and build it from the sources. You have to retrieve the project sources by using one of the following way:

```shell
$ go get -u github.com/kinduff/firefly_iii_exporter
# or
$ git clone https://github.com/kinduff/firefly_iii_exporter.git
```

Install the needed vendors:

```shell
$ GO111MODULE=on go mod vendor
```

Then, build the binary:

```shell
$ go build -o firefly_iii_exporter .
```


## Configuration

You can use both environment variables or parameters in both the binary or the docker image. If you wish to use parameters, simply invoke the same environment variable but in downcase, or use the flag `--help` for more information.

| Environment variable | Description                                                  | Default | Required |
| -------------------- | ------------------------------------------------------------ | ------- | -------- |
| `BASE_URL`           | URL path including protocol to the Firefly III instance      |         | Yes      |
| `API_KEY`            | Your Personal Access Token                                   |         | Yes      |
| `HTTP_PORT`          | The port the exporter will be running the HTTP server        | 4002    |          |
| `SCRAPE_INTERVAL`    | Time in natural format to scrap statistics from the instance | `30s`   |          |

## Available Prometheus metrics

| Metric name                                   | Description                                          |
| --------------------------------------------- | ---------------------------------------------------- |
| `firefly_iii_account_balance_metrics`         | Current balance of each of the available accounts    |
| `firefly_iii_account_transactions_metrics`    | Total transactions of each of the available accounts |
| `firefly_iii_accounts_metrics`                | Total accounts available                             |
| `firefly_iii_transaction_by_category_metrics` | Total transactions per available category            |
| `firefly_iii_transactions_metrics`            | Total transactions available                         |

[codefactor]: https://www.codefactor.io/repository/github/kinduff/firefly_iii_exporter
[configuration]: #configuration
[docker]: https://docs.docker.com
[dockerhub]: https://hub.docker.com/r/kinduff/firefly_iii_exporter
[ghcr]: #ghcr
[go]: https://golang.org
[godoc]: https://pkg.go.dev/github.com/kinduff/firefly_iii_exporter
[goreport]: https://goreportcard.com/report/github.com/kinduff/firefly_iii_exporter
[releases]: https://github.com/kinduff/firefly_iii_exporter/releases
[steam-api]: https://steamcommunity.com/dev/apikey
[workflow-c]: https://github.com/kinduff/firefly_iii_exporter/actions/workflows/ci.yml
[metrics_example]: extra/metrics_example.txt
