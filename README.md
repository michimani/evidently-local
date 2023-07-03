Evidently-Local
===

[![codecov](https://codecov.io/gh/michimani/evidently-local/branch/main/graph/badge.svg?token=YJIO3RF9ZR)](https://codecov.io/gh/michimani/evidently-local)

This is a simple http server to manage feature flags, compatible with Amazon CloudWatch Evidently.

# Supported Evidently Action

[Actions - CloudWatch Evidently](https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_Operations.html)

- EvaluateFeature
  - Only support evaluation with default variation and override rules.
  - TODO: Evaluation with some launches.

# Usage

## Simple usage

### 1. Create JSON files that defines the Project and Feature.

First, create a JSON file that defines the Project and Feature in the following configuration.

```text
data
└── projects
    └── test-project
        └── features
            ├── test-feature-1.json
            └── test-feature-2.json
```

In this case, project name is **test-project** and features that belong to the project are **test-feature-1** and **test-feature-2**.

And `test-feature-1.json` is like this.

```json
{
  "defaultVariation": "False",
  "entityOverrides": {
    "force-true": "True"
  },
  "name": "test-feature-1",
  "project": "test-project",
  "valueType": "BOOLEAN",
  "variations": [
    {
      "name": "True",
      "value": {
        "boolValue": true
      }
    },
    {
      "name": "False",
      "value": {
        "boolValue": false
      }
    }
  ]
}
```

In this case, `test-feature-1` is a Feature that returns a boolean that belongs to `test-project`, and the default Variation is `False`. Also, an override rule is set to always return `True` for the EntityID `force-true`.


This JSON file has the same structure as the JSON that can be obtained with the GetFeature API. So, if you want to reproduce the Feature that already exists on AWS locally, you can use the JSON obtained with the following AWS CLI command as it is.

```bash
aws evidently get-feature \
--project-name 'test-project' \
--feature-name 'test-feature-1' \
--query 'feature'
--output json \
> test-feature-1.json
```

### 2. Create a Dockerfile and run Evidently-Local

Second, create a `Dockerfile` to run Evidently-Local server. The following is an example of `Dockerfile`.

```dockerfile
FROM golang:1.20-alpine3.18 AS builder

WORKDIR /app

ADD https://github.com/michimani/evidently-local/archive/refs/tags/v0.0.2.zip ./

RUN unzip v0.0.2.zip \
  && cd evidently-local-0.0.2 \
  && go install \
  && go build -o evidently-local . \
  && mv evidently-local /app

# for run stage
FROM alpine:3.18.2

WORKDIR /app

COPY --from=builder /app/evidently-local .

RUN mkdir data

# from your local data directory
ADD ./testdata ./data

EXPOSE 2306

CMD [ "/app/evidently-local" ]
```

Then build image and run container.

```bash
docker build -t evidently-local . \
&& docker run  -p 2306:2306 evidently-local:latest
```

### 3. Call EvaluateFeature API

Finally, call EvaluateFeature API using AWS SDK for each language. The following is an example of calling the API using Go SDK.

```go
package main

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/gofrs/uuid"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/evidently"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
)

const (
	evidentlyEndpointURLKey = "EVIDENTLY_ENDPOINT_URL"
	project                 = "test-project"
	feature                 = "test-feature-1"
	region                  = "ap-northeast-1"
)

func main() {
	client, err := createEvidentlyClient()
	if err != nil {
		panic(err)
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	entityID := uuid.String()

	args := os.Args
	if len(args) > 1 {
		entityID = args[1]
	}

	evaluateFeature(client, project, feature, entityID)
}

func evaluateFeature(client *evidently.Client, project, feature, entityID string) {
	out, err := client.EvaluateFeature(context.Background(), &evidently.EvaluateFeatureInput{
		Project:  aws.String(project),
		Feature:  aws.String(feature),
		EntityId: aws.String(entityID),
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("EntityID: %s\n", entityID)
	fmt.Printf("Reason: %s\n", aws.ToString(out.Reason))
	fmt.Printf("Variation: %s\n", aws.ToString(out.Variation))
	fmt.Printf("Type of Value: %s\n", reflect.TypeOf(out.Value))

	var value any
	switch out.Value.(type) {
	case *types.VariableValueMemberStringValue:
		value = out.Value.(*types.VariableValueMemberStringValue).Value
	case *types.VariableValueMemberBoolValue:
		value = out.Value.(*types.VariableValueMemberBoolValue).Value
	case *types.VariableValueMemberLongValue:
		value = out.Value.(*types.VariableValueMemberLongValue).Value
	case *types.VariableValueMemberDoubleValue:
		value = out.Value.(*types.VariableValueMemberDoubleValue).Value
	default:
		// noop
	}

	fmt.Printf("Value: %+v\n", value)
}

func createEvidentlyClient() (*evidently.Client, error) {
	evidentlyEndpointURL := os.Getenv(evidentlyEndpointURLKey)
	if len(evidentlyEndpointURL) == 0 {
		cfg, err := config.LoadDefaultConfig(context.Background(),
			config.WithRegion(region),
		)
		if err != nil {
			return nil, err
		}

		c := evidently.NewFromConfig(cfg)
		return c, nil
	}

	// create client for custom endpoint
	customEndpointResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, opts ...any) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               evidentlyEndpointURL,
			HostnameImmutable: true,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithEndpointResolverWithOptions(customEndpointResolver),
	)
	if err != nil {
		return nil, err
	}

	c := evidently.NewFromConfig(cfg)
	return c, nil
}
```

Run this script with the following command.

```bash
EVIDENTLY_ENDPOINT_URL='http://localhost:2306' go run .
```

Then you will get the following result.

```text
EntityID: 4eb2759a-c96b-43b3-acf2-25c669ab280b
Reason: DEFAULT (local)
Variation: False
Type of Value: *types.VariableValueMemberBoolValue
Value: false
```

Also, if you run it as follows, you will get the result of the override rule.

```bash
EVIDENTLY_ENDPOINT_URL='http://localhost:2306' go run . 'force-true'
```

```text
EntityID: force-true
Reason: OVERRIDE_RULE (local)
Variation: True
Type of Value: *types.VariableValueMemberBoolValue
Value: true
```


# License

[MIT](./LICENSE)

# Author

[michimani210](https://twitter.com/michimani210)