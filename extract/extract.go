package extract

import (
	"context"
	"flag"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type SSMClient interface {
	GetParameters(ctx context.Context, params *ssm.GetParametersInput, optFns ...func(*ssm.Options)) (*ssm.GetParametersOutput, error)
}

// Parses the command line arguments and returns two slices of strings for futher interpretation.
func Argv(argv []string) (preDashArgs []string, postDashArgs []string) {
	dashFound := false
	for _, arg := range argv[1:] {
		if arg == "--" {
			dashFound = true
			continue
		}

		if dashFound {
			postDashArgs = append(postDashArgs, arg)
		} else {
			preDashArgs = append(preDashArgs, arg)
		}
	}
	return preDashArgs, postDashArgs
}

// Parses the pre-dash arguments and returns a slice of SSM parameter paths.
func Paths(preDash []string) (ssmPaths []string, err error) {
	ssmFlag := flag.NewFlagSet("Entry", flag.ExitOnError)
	ssmFlag.Func("path", "Specify SSM parameter path to fetch", func(s string) error {
		ssmPaths = append(ssmPaths, s)
		return nil
	})

	if err := ssmFlag.Parse(preDash); err != nil {
		return nil, err
	}
	return ssmPaths, nil
}

// Fetches SSM parameters and returns a slice of environment variable strings.
func SSM(ctx context.Context, ssmClient SSMClient, ssmPaths []string) (mergedParams []types.Parameter, err error) {
	resp, err := ssmClient.GetParameters(ctx, &ssm.GetParametersInput{
		Names:          ssmPaths,
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return []types.Parameter{}, err
	}

	return append(mergedParams, resp.Parameters...), nil
}
