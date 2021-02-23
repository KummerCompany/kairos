# Kairos
A Framework for testing API in golang.

## Example
```go
func TestAPI(t *testing.T) {

	testAPI := test.NewTestingAPI()

	testAPI.Description("Testing GET API",
		test.Before{
			Method: "GET",
			URL:    "/api/video/251252",
			Body:   nil,
			BeforeFunc: func(req *http.Request) {
				controllers.Repo = repository.NewRepositoryStub(models.Video{
					Name: "Juan",
				}, nil)
			},
		},
		test.Test{
			Controller: controllers.GetVideo,
		},
		test.After{
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