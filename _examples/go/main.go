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
