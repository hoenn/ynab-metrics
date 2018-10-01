# YNAB Metrics Exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/hoenn/ynab-metrics)](https://goreportcard.com/report/github.com/hoenn/ynab-metrics)

[YNAB](https://www.youneedabudget.com/) recently made an [API](https://api.youneedabudget.com/) to retrieve data for your user account. This project will allow you to poll the YNAB API with your personal access token and it will create and post [prometheus](https://github.com/prometheus/prometheus) metrics on the port of your choice.

## Use case
The YNAB browser and mobile apps have some graphs available, namely Net Worth at a monthly resolution, but you may be interested in creating your own graphs, at a much finer resolution, on the specific topics worth seeing. Prometheus will scrape the data that the exporter exposes and organize the time series data we would need.

## Grafana
![screenshot](links)

#### Premade dashboard

## Metrics Collected

## Request new metrics

## Limitations

## Installation

## Simple guide to running everything you need

### Docker compose
