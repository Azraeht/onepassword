# 1Password Client
Thin wrapper around the 1Password CLI for use in Golang.

__This is a stub and should not be used in production.__
> Hopefully 1Password creates their own client library. See
[discussion](https://discussions.agilebits.com/discussion/80802/why-a-cli-and-not-a-library).

## Usage
First [install the 1Password CLI](https://support.1password.com/command-line/).

Import the package, create a client, and retrieve an item.
```go
import (
    "os"

    op "github.com/azraeht/onepassword"
)

func main() {
    password := os.GetEnv("OP_PASSWORD")
    secretKey := os.GetEnv("OP_SECRET_KEY")
    client := op.NewClient("op", "subdomain", "test@subdomain.com", password, secretKey)
    item := client.GetItem(VaultName("test-vault"), ItemName("test-item"))
}
```

## Resources
- [terraform-provider-onepassword](https://github.com/anasinnyk/terraform-provider-1password)
- [1Password CLI](https://support.1password.com/command-line/)
