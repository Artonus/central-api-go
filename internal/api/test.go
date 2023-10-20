package api

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
)

func (api *api) testGraphConnection(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	query := "CREATE (n:Location {id: 'c52da186-683f-475a-987f-4645f0292b25', name:'MiddlewareBB'});"
	_, err := neo4j.ExecuteQuery(ctx, api.graphDriver, query, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(""))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
