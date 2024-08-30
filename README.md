# VKN Query Library

This library is used to query registered individuals or businesses in the Revenue Administration of Turkey.

## Getting Started

### Installation

To add this library to your project:

```bash
go get github.com/9ssi7/vkn
```

### Usage

```go
package main

import (
    "fmt"
    "github.com/9ssi7/vkn"
)

func main() {
    config := vkn.Config{
        Username: "your-username",
        Password: "your-password",
    }

    client := vkn.New(config)
    data, err := client.GetRecipient(context.Background(), "your-recipients-vkn-or-tck")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(data) // {FirstName: "John", LastName: "Doe", Title: "Company Name", TaxOffice: "Tax Office Name"}
}
```

## Important Notes

- When querying with identity credentials (TCKN), you will receive the firstName and lastName fields, but the title (unvan) field will not be present.
- When querying with a tax number (VKN), the title (unvan) field will be returned instead of the firstName and lastName fields.

## Documentation

For detailed documentation, you can visit the [here](https://pkg.go.dev/github.com/9ssi7/vkn).

## License

This library is licensed under the MIT license. For more information, see the LICENSE file.
