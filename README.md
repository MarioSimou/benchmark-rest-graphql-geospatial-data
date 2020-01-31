# REST && Postgresql for Geospatial Data 

### Benchmarking

Benchmark the `/api/v1/cy/population` endpoint with the following code:

```
package main

import (
	"net/http"
	"testing"
)

func BenchmarkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.Get("http://localhost:3000/api/v1/cy/population")
	}
}
``