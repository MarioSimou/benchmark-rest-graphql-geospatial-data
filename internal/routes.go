package internal

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

type controller struct {
	Db *sql.DB
}

func New(db *sql.DB) *controller {
	return &controller{Db: db}
}

func (c *controller) GraphqlEndpoint(w http.ResponseWriter, r *http.Request) {
	var body Body
	json.NewDecoder(r.Body).Decode(&body)
	schema, _ := NewSchema(c.Db)

	json.NewEncoder(w).Encode(graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  body.Query,
		OperationName:  body.OperationName,
		VariableValues: body.Variables,
	}))
}
