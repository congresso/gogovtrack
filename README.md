## Go Govtrack API

#Usage

```go
api := NewApi(http.DefaultClient, gogovtrack.BaseURL)

// Get all bills
api.Br().All()

// Filtered bills
billsResp, err := api.Br().Filter(gogovtrack.Q{"limit": "1", "order_by": "-current_status_date"})
if err != nil {
  return err
}

// Get one bill
bill, err := api.Br().One("id")
if err != nil {
  return err
}
```
