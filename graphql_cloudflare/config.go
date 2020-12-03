package main

import (
	"os"

	"github.com/shurcooL/graphql"
)

type env struct {
	api_key string
	email   string
}

var cloudflare_server = "https://api.cloudflare.com/client/v4/graphql"

func readEnv() env {
	return env{os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_EMAIL")}
}

var variables = map[string]interface{}{
	"Account":    graphql.ID("26df4c5080269e5fc427fef927573afd"),
	"DatetimeGT": "2020-08-09T04:40:00Z",
	"DatetimeLT": "2020-08-09T05:40:00Z",
	"Limit":      10,
}

var WorkersAnaliticsQuery struct {
	Viewer struct {
		Accounts []struct {
			WorkeersInvocationAdaptive []struct {
				Sum struct {
					Requests    graphql.Int
					Subrequests graphql.Int
					Errors      graphql.Int
				}
				Dimensions struct {
					ScriptName graphql.String
				}
			} `graphql:"workersInvocationsAdaptive(limit: $Limit, filter: { datetime_gt: $DatetimeGT, datetime_lt: $DatetimeLT})"`
		} `graphql:"accounts(filter: {accountTag: $Account})"`
	}
}
