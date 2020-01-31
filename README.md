# GraphQL && Postgresql for Geospatial Data


## GraphqlQuery

```
query {
  population {
    type
    name
    features {
      ...featureFragment
    }
  }
}

fragment geometryFragment on Geometry {
  type
  coordinates
}

fragment propertiesFragment on Properties {
  gmlId
  localId
  namespace
  versionId
  notCountedProportion
  localisedCharacterString
  endPosition
}

fragment featureFragment on Feature {
 type
 properties {
   ...propertiesFragment
 }
 geometry {
   ...geometryFragment
 }
}

```

### Benchmark

Benchmark the `/graphql` endpoint with the following code:

```
package main

import (
	"bytes"
	"log"
	"net/http"
	"testing"
)

func BenchmarkEndpoint(b *testing.B) {
	var requestBody = []byte(`{"query":"query{population{type name features{type properties{gmlId localId namespace versionId localisedCharacterString notCountedProportion endPosition duration} geometry{type coordinates}}}}"}`)

	for i := 0; i < b.N; i++ {
		res, e := http.Post("http://localhost:3000/graphql", "application/json", bytes.NewBuffer(requestBody))
		if e != nil {
			log.Fatalln(e)
		}
		defer res.Body.Close()
	}
}

```