# Kairos
A Framework for testing API in golang.

## How to install
`go get github.com/KummerCompany/kairos`


## Example
```go

import (
	"github.com/KummerCompany/kairos"
	kairoRepo "github.com/KummerCompany/kairos/repositories"
	)
	
func TestAPI(t *testing.T) {

	testAPI := kairoRepo.NewTestingAPI()

	testAPI.Description("Testing GET API",
		kairos.Before{
			Method: "GET",
			URL:    "/api/video/251252",
			Body:   nil,
			BeforeFunc: func(req *http.Request) {
				controllers.Repo = repository.NewRepositoryStub(models.Video{
					Name: "Juan",
				}, nil)
			},
		},
		kairos.Test{
			Controller: controllers.GetVideo,
		},
		kairos.After{
			AfterFunc:  nil,
			StatusCode: 200,
			Want: map[string]interface{}{
				"id":       "251252",
				"name":     "Juan",
				"year":     0,
				"director": map[string]interface{}{},
			},
		},
		t,
	)
}
```
