# Go Govtrack HTTP Client

[https://godoc.org/github.com/congresso/gogovtrack](https://godoc.org/github.com/congresso/gogovtrack)

## Example

```go
  client := gogovtrack.Client{
    HTTPClient: http.DefaultClient,
  }

  // Get One
  b, err := client.GetBill(1)
  if err != nil {
    log.Fatal(err)
  }

  // Get All
  cs, err := client.GetCommittess()
  if err != nil {
    log.Fatal(err)
  }

  // Filter Get All
  rs, err := client.Filter(gogovtrack.Q{"limit": "20", "id": "1"}).GetRoles()
  if err != nil {
    log.Fatal(err)
  }
```
