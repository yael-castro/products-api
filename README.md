# Products RESTful API - Coding evaluation

REST API for product storage management
###### Quick Links
- [Required environment variables](.env.example)
- [API documentation (Swagger)](swagger.yaml)
- [API documentation (Postman)](https://documenter.getpostman.com/view/12474312/VUxLvTJR)

###### How to run (Golang)
```shell
# Load required environment variables
export $(grep -v ^# .env.example)
# Run SQL migrations
go run ./cmd/migrations/migrations.go
# Run HTTP server
go run ./cmd/server/server.go
```

###### Unit tests
#### Structure
```go
package foo

import (
	"strconv"
	"testing"
)

func TestX(t *testing.T) {
	// Table Driven Testing 
	tdt := []struct{}{
	    {},	// Test Case
    }
	
	// Setup
	x := struct{}{}
	
	// Cleanup function (Optional)
	t.Cleanup(func() {})
	
	// Subtests 
	for i, v := range tdt {
		t.Run(strconv.Itoa(i), func(t *testing.T){
			t.Log(x, v)
			// Cleanup function for each subtest (Optional)
			t.Cleanup(func() {})
		})
	}
}
```
#### Notes
- Added custom cli flags to modify unit test behavior (for example, if the -V flag is used when running unit tests, additional logs are displayed)
###### Architecture style explained
The architecture pattern used in this project is the most common form of layered architecture pattern 
with three layers (presentation, business, and data access) and some additional changes about
communication between layers, due to architecture decisions (explained below ππΌ)

```
βββ cmd            ππΌ (executable commands)
βββ internal
    βββ business   ππΌ (business logic layer)
    βββ dependency ππΌ (manage dependencies)
    βββ handler    ππΌ (presentation layer)
    βββ model      ππΌ (data transfer objects, business objects and enums)
    βββ repository ππΌ (data access layer)
```

[See more about layer architecture pattern](https://github.com/yael-castro/layered-architecture)

###### Communication between layers
π₯IMPORTANTπ₯ The layers are integrated in a top-down communication model: each layer can hold dependency only on the layer directly beneath it.
<hr>

One way to communicate the layers to share data between it is to use `Domain Objects` to communicate the `Presentation Layer`
with the `Business Layer` and `Data Transfer Objects` to communicate the `Business Layer` with the `Access Data Layer`

```
     
    |--------------------|                      |----------------------|                             |-------------------|
    | Presentation layer | ==[domain object]==> | Business logic layer | ==[data transfer object]==> | Data access layer |
    |--------------------|                      |----------------------|                             |-------------------|

```

However, many times data transfer objects and domain objects are similar or identical, which lead to redundancy...
To solve this problem I decided create the model package to be the way in that the three layers can be shared information
using data types in common.

```
     
    |--------------------|           |----------------------|           |-------------------|
    | Presentation layer | ========> | Business logic layer | ========> | Data access layer |
    |--------------------|           |----------------------|           |-------------------|
              ^                                 ^                                 ^
              |                                 |                                 |
              |=================================|=================================|
                                                |
                                     |----------------------|
                                     |     Package model    |
                                     |----------------------|
```
<hr>
