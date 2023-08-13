example
===

This is an example for using evidently-local.

## Build and run evidently-local

### Build

```bash
docker build -t evidently-local .
```

### Run

```bash
docker run -p 2306:2306 evidently-local:latest
```

## Request to evidently local

### by Python code (boto3)

Please read [python/README.md](./python/README.md).


### by Go code (AWS SDK for Go v2)

Please read [go/README.md](./go/README.md).