package main

import (
	"fmt"
	"github.com/bndr/gopencils"
)

type respStruct struct {
	Args          map[string]string
	Headers       map[string]string
	Origin        string
	Url           string
	Authorization string
}

func main() {
	// Create Basic Auth
	auth := gopencils.BasicAuth{"username", "password"}
	// Create New Api with our auth
	api := gopencils.Api("http://your-api-url.com/api/", &auth)

	// Maybe some payload to send along with the request?
	payload := map[string]interface{}{"Key": "Value1"}

	// Perform a GET request
	// URL Requested: http://your-api-url.com/api/users
	api.Res("users", &respStruct{}).Get()

	// Get Single Item
	api.Res("users", &respStruct{}).Id(1).Get()

	// Perform a GET request with Querystring
	querystring := map[string]string{"page": "100", "per_page": "1000"}
	// URL Requested: http://your-api-url.com/api/users/123/items?page=100&per_page=1000
	api.Res("users").Id(123).Res("items", resp).Get(querystring)

	// Or perform a POST Request
	// URL Requested: http://your-api-url.com/api/items/123 with payload as json Data
	api.Res("items", resp).Id(123).Post(payload)

	// Users endpoint
	users := api.Res("users")

	for i := 0; i < 10; i++ {
		// Create a new pointer to response Struct
		user := new(respStruct)
		// Get user with id i into the newly created response struct
		users.Id(i, user).Get()
	}

}
