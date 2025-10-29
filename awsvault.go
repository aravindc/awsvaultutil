package awsutil

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

// AWSCreds holds temporary AWS credentials from aws-vault
type AWSCreds struct {
	AccessKeyId     string    `json:"AccessKeyId"`
	SecretAccessKey string    `json:"SecretAccessKey"`
	SessionToken    string    `json:"SessionToken"`
	Expiration      time.Time `json:"Expiration"`
}

// GetAWSCredsFromVault runs aws-vault to retrieve credentials for a given profile
func GetAWSCredsFromVault(profile string) (*AWSCreds, error) {
	cmd := exec.Command("aws-vault", "exec", profile, "--json")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run aws-vault: %w", err)
	}

	var creds AWSCreds
	if err := json.Unmarshal(out, &creds); err != nil {
		return nil, fmt.Errorf("failed to parse aws-vault output: %w", err)
	}

	return &creds, nil
}

// GetAWSConfigFromVault returns an AWS SDK v2 config using aws-vault credentials
func GetAWSConfigFromVault(profile string, opts ...func(*config.LoadOptions) error) (aws.Config, error) {
	creds, err := GetAWSCredsFromVault(profile)
	if err != nil {
		return aws.Config{}, err
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		append(opts,
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					creds.AccessKeyId,
					creds.SecretAccessKey,
					creds.SessionToken,
				),
			),
		)...,
	)
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return cfg, nil
}
