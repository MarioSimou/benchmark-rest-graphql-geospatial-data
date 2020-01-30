package internal

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/mongo"
)

type controller struct {
	Db *mongo.Database
}

func New(db *mongo.Database) *controller {
	return &controller{db}
}

func (contr *controller) GraphqlEndpoint(w http.ResponseWriter, r *http.Request) {
	var body Body
	json.NewDecoder(r.Body).Decode(&body)
	schema, _ := NewSchema(contr.Db)

	var params = graphql.Params{
		Schema:         schema,
		OperationName:  body.OperationName,
		VariableValues: body.Variables,
		RequestString:  body.Query,
	}

	json.NewEncoder(w).Encode(graphql.Do(params))
}
