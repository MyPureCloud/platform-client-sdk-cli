package oauth_clients_usage_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/oauth_clients_usage_query_results"
)

func init() {
	oauth_clients_usage_queryCmd.AddCommand(oauth_clients_usage_query_results.Cmdoauth_clients_usage_query_results())
	oauth_clients_usage_queryCmd.Short = utils.GenerateCustomDescription(oauth_clients_usage_queryCmd.Short, oauth_clients_usage_query_results.Description, )
	oauth_clients_usage_queryCmd.Long = oauth_clients_usage_queryCmd.Short
}
