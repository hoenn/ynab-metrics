# YNAB Metrics Exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/hoenn/ynab-metrics)](https://goreportcard.com/report/github.com/hoenn/ynab-metrics)

[YNAB](https://www.youneedabudget.com/) recently made an [API](https://api.youneedabudget.com/) to retrieve data for your user account. This project will allow you to poll the YNAB API with your personal access token and it will create and post [prometheus](https://github.com/prometheus/prometheus) metrics on the port of your choice.

## Use case
The YNAB browser and mobile apps have some graphs available, namely Net Worth at a monthly resolution, but you may be interested in creating your own graphs, at a much finer resolution, on the specific topics worth seeing. Prometheus will scrape the data that the exporter exposes and organize the time series data we would need.

## Grafana
![screenshot](links)

#### Premade dashboard

## Metrics Collected
The metrics currently collected are

|Metric name | Labels| Value |
|-----|:----:|----:|
|account_balance |budget_name, name| balance|
|category_activity|budet_name, name| activity|
|category_budget| budget_name, name | budget|
|budget_transactions | budget_name, category_name, payee_name, transactions_id| amount|
|rate_limit_total| | API limit|
|rate_limit_used| | requests this hour|

## Limitations

## Installation
```
$ git clone https://github.com/Hoenn/ynab-metrics
$ cd ynab-metrics/
$ go get -d ./...
```
The project is ready to be built and run, before continuing create `ynab-metrics/config.json` based on the `sample.json` file in the same directory. Replace `access_token` with your actual access token.

You can create an access token by accessing [app.youneedabudget.com/settings](app.youneedabudget.com/settings), signing in. Scroll down to Developer Settings and create a New Token. Copy the token after it's created and paste it into `config.json`

Now that `config.json` has your access token we can build and run `ynab-metrics`

```
$ make build
$ make run
```

As long as no error messages appear, navigate to `http://localhost:8080` in your browser and you should see metrics appearing.a

## Simple guide to running everything you need

### Do it yourself

Check out the install guides for [Prometheus](https://github.com/prometheus/prometheus) and [Grafana](https://github.com/grafana/grafana), all you should need is to know where `ynab-metrics` is posting metrics once it's running (default `localhost:8080`)

### Docker Compose
You can run the full prometheus, grafana, ynab-metrics stack with the docker-compose file in `ynab-metrics/docker/`

Make sure you `make build` before trying to `docker-compose up`. After that's done you can `cd docker` and `docker-compose up`. Include `-d` in your command if you want the containers to run in the background.

By default: Grafana will run on port 3000, Prometheus on 9090, and the ynab-metrics exporter on 8080.

The `ynab-metrics/config.json` file will be mounted and used to run the ynab-metrics container the `ynab-metrics/docker/prometheus/prometheus.yml` config is used for prometheus.
