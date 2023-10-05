package api

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
)

func (api *api) testGraphConnection(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	session := api.graphDriver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: ""})
	defer session.Close(ctx)
	// Run index queries via implicit auto-commit transaction
	index := "create index on :Location(id)"
	index2 := "create index on :Location(name)"
	_, err := session.Run(ctx, index, nil)

	if err != nil {
		panic(err)
	}
	_, err = session.Run(ctx, index2, nil)
	if err != nil {
		panic(err)
	}
	query := "CREATE (n:Location {id: 'c52da186-683f-475a-987f-4645f0292b25', name:'MiddlewareBB'});"
	_, err = neo4j.ExecuteQuery(ctx, api.graphDriver, query, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(""))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
