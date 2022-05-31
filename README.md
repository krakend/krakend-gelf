# KrakenD-gelf

[![Go Report Card](https://goreportcard.com/badge/github.com/krakendio/krakend-gelf)](https://goreportcard.com/report/github.com/krakendio/krakend-gelf)

A gelf (Graylog Extended Log Format) Writer for [KrakenD](https://krakend.io) loggers.

## How to use it

This package just return a gelf writer with the configuration provided via KrakenD ExtraConfig.

You need to add the Writer to the logger writers.
This example uses [KrakenD-gologging](https://github.com/devopsfaith/krakend-gologging).

Import the package

```
import "github.com/krakendio/krakend-gelf"
```

Create a new Writer:

```
gelfWriter, err := gelf.NewWriter(cfg.ExtraConfig)
```

And add it to the logger:

```
gologging.NewLogger(cfg.ExtraConfig, gelfWriter...)
```

## Configuration

Add the `github_com/devopsfaith/krakend-gelf` section to the service extra config.

There's 2 parameters:

- address (This parameter is **required**)

  The address (including the port) of your graylog server (or any service that receives gelf inputs).

- enable_tcp

  By default uses UDP but you can enable TCP by setting this option to true (not recommended, your performance may suffer).

Example:

```
"extra_config": {
  "github_com/devopsfaith/krakend-gelf": {
    "address": "myGraylogInstance:12201",
    "enable_tcp": false
  }
}
```
