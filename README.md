# Aws Vault Util

`awsvaultutil` is a Go utility library for working with AWS SDK v2 using credentials managed by [`aws-vault`](https://github.com/99designs/aws-vault).  
It provides functions to programmatically fetch temporary AWS credentials from an `aws-vault` profile and generate an AWS SDK v2 configuration, making it easy to use AWS services securely across projects.

---

## Features

- Fetch temporary AWS credentials from `aws-vault` profiles.
- Generate `aws.Config` for AWS SDK v2 with `StaticCredentialsProvider`.
- Compatible with session tokens and temporary credentials.
- Can be imported and used across multiple Go projects.
- Optional support for customizing `LoadOptions` when loading AWS config.

---

## Installation

```bash
go get github.com/yourusername/awsvaultutil@latest
```

---

## Usage

### Import the library

```go
import "github.com/yourusername/awsvaultutil"
```

### Get AWS credentials from aws-vault

```go
package main

import (
    "fmt"
    "log"

    "github.com/yourusername/awsvaultutil"
)

func main() {
    creds, err := awsvaultutil.GetAWSCredsFromVault("myapp-dev")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("AccessKeyId: %s\n", creds.AccessKeyId)
    fmt.Printf("SessionToken: %s\n", creds.SessionToken)
}
```

### Create an AWS SDK v2 configuration

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/yourusername/awsvaultutil"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
    cfg, err := awsvaultutil.GetAWSConfigFromVault("myapp-dev")
    if err != nil {
        log.Fatal(err)
    }

    ec2Client := ec2.NewFromConfig(cfg)

    output, err := ec2Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d reservations\n", len(output.Reservations))
}
```

---

## API

### `GetAWSCredsFromVault(profile string) (*AWSCreds, error)`

Fetch temporary AWS credentials from an `aws-vault` profile.

- **Parameters**: `profile` - The name of the AWS profile stored in `aws-vault`.
- **Returns**: `AWSCreds` containing `AccessKeyId`, `SecretAccessKey`, `SessionToken`, `Expiration`.

---

### `GetAWSConfigFromVault(profile string, opts ...func(*config.LoadOptions) error) (aws.Config, error)`

Generate an AWS SDK v2 configuration using credentials from `aws-vault`.

- **Parameters**:

  - `profile` - The AWS profile stored in `aws-vault`.
  - `opts` - Optional AWS SDK `LoadOptions`.

- **Returns**: Config ready to use with AWS services.

---

## Requirements

- Go 1.20+
- [`aws-vault`](https://github.com/99designs/aws-vault) installed and configured
- AWS SDK v2 (`github.com/aws/aws-sdk-go-v2`)

---

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

---

## License

This project is licensed under the **Apache License 2.0**. See the [LICENSE](LICENSE) file for details.
