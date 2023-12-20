# entry

[![Main](https://github.com/raylas/nextdns-exporter/actions/workflows/main.yaml/badge.svg)](https://github.com/raylas/nextdns-exporter/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/raylas/nextdns-exporter)](https://goreportcard.com/report/github.com/raylas/nextdns-exporter)

A package and execution wrapper for AWS SSM Parameter expansion.

## Usage

### CLI

```bash
Usage: main [-g] [-p PREFIX] [COMMAND [ARGUMENTS [ARGUMENTS ...]]]

Positional arguments:
  COMMAND                Command to run
  ARGUMENTS              Command arguments

Options:
  -g                     Do not inherit environment
  -p PREFIX              SSM prefixes to source
  --help, -h             display this help and exit
  --version              display version and exit
```

### Package

```json
{
  "foo": "oof",
  "bar": "rab"
}
```

```go
var data struct {
  Foo string `json:"foo"`
  Bar string `json:"bar"`
}

awsConfig, err := config.LoadDefaultConfig(context.TODO())
if err != nil {
  panic("error loading AWS credentials")
}

p := &kv.Parameters{Client: ssm.NewFromConfig(awsConfig)}

// Get
params, err := p.Get([]string{"/dev/foobar"})
if err != nil {
  panic("error getting parameter")
}

// Unmarshal
if err := p.Unmarshal("/dev/foobar", &data); err != nil {
  panic("error unmarshalling parameter")
}
```
