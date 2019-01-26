# GoCanvas

TODO

```go
// Create a new Provider.
p := service.Provider{
    BaseURL: "canvas.uva.nl/api/v1/",
    AccessToken: "{your-access-token}",
}

// Fetch the profile using the provider.
var profile *models.Profile 
profile, ok = p.FetchProfile()
```
