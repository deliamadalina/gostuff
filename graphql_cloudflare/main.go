package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/shurcooL/graphql"
)

func main() {

	client := &http.Client{
		Transport: newRoundTripper{r: http.DefaultTransport},
	}
	graphqlClient := graphql.NewClient("https://api.cloudflare.com/client/v4/graphql", client)

	err := graphqlClient.Query(context.Background(), &WorkersAnaliticsQuery, variables)
	if err != nil {
		print(&WorkersAnaliticsQuery)
		panic(err)
	}
	print(WorkersAnaliticsQuery)
}

func print(q interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(q)
	if err != nil {
		panic(err)
	}
}
