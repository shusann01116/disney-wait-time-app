package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/disney-wait-time-app/app/collector_go/core"
	"github.com/disney-wait-time-app/app/collector_go/repository"
	"github.com/disney-wait-time-app/app/collector_go/tdr"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda/xrayconfig"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
)

var (
	TABLE_NAME_KEY = "DYNAMODB_TABLENAME"
	StandbyURLs    = []string{
		"https://www.tokyodisneyresort.jp/_/realtime/tdl_attraction.json",
		"https://www.tokyodisneyresort.jp/_/realtime/tds_attraction.json",
	}
)

func lambdaHandler(
	ctx context.Context,
) func(ctx context.Context, event events.APIGatewayV2HTTPRequest) (interface{}, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	otelaws.AppendMiddlewares(&cfg.APIOptions)
	dynamodbClient := dynamodb.NewFromConfig(cfg)
	httpClient := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	return func(ctx context.Context, event events.APIGatewayV2HTTPRequest) (interface{}, error) {
		// get table name from env value
		tableName, ok := os.LookupEnv(TABLE_NAME_KEY)
		if !ok {
			log.Error().Stack().Msgf("env value %s is not set", TABLE_NAME_KEY)
			return events.APIGatewayV2HTTPResponse{
				StatusCode: 400,
				Body:       "",
			}, fmt.Errorf("env value %s is not set", TABLE_NAME_KEY)
		}

		table := repository.Table{
			DynamoDBClient: dynamodbClient,
			TableName:      tableName,
		}

		tdrclient := tdr.TdrClient{
			Ctx:    ctx,
			Client: httpClient,
		}
		standbys, err := tdrclient.FetchStandbys(StandbyURLs)
		if err != nil {
			log.Fatal().Msgf("failed to fetch standby times: %v\n", err)
		}
		log.Info().Msgf("got %v standbys\n", len(standbys))

		// facility
		facility := core.NewFacility(0, "hoge", "ホゲ")
		err = table.AddFacility(ctx, facility)
		if err != nil {
			log.Error().Stack().Err(err).Msg("")
			return events.APIGatewayV2HTTPResponse{StatusCode: 400, Body: "\"Failed!\""}, err
		}

		for _, s := range standbys {
			// standby
			s.FacilityID = facility.ID
			err = table.AddStandby(ctx, s)
			if err != nil {
				log.Error().Stack().Err(err).Msg("")
				return events.APIGatewayV2HTTPResponse{StatusCode: 400, Body: "\"Failed!\""}, err
			}
			// latest standby
			attr, err := table.UpdateLatestStandby(ctx, s)
			if err != nil {
				log.Error().Stack().Err(err).Msg("")
				return events.APIGatewayV2HTTPResponse{StatusCode: 400, Body: "\"Failed!\""}, err
			}
			log.Info().Msgf("latest standby: %v", attr)
		}

		return events.APIGatewayV2HTTPResponse{
			StatusCode: 200,
			Body:       "\"Hello from Lambda!\"",
		}, nil
	}
}

func main() {
	ctx := context.Background()

	tp, err := xrayconfig.NewTracerProvider(ctx)
	if err != nil {
		fmt.Printf("error creating tracer provider: %v", err)
	}

	defer func(ctx context.Context) {
		err := tp.Shutdown(ctx)
		if err != nil {
			fmt.Printf("error shutting down tracer provider: %v", err)
		}
	}(ctx)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(xray.Propagator{})
	lambda.Start(
		otellambda.InstrumentHandler(lambdaHandler(ctx), xrayconfig.WithRecommendedOptions(tp)...),
	)
}
