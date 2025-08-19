package oauth_clients_usage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/oauth_clients_usage_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/oauth_clients_usage_summary"
)

func init() {
	oauth_clients_usageCmd.AddCommand(oauth_clients_usage_query.Cmdoauth_clients_usage_query())
	oauth_clients_usageCmd.AddCommand(oauth_clients_usage_summary.Cmdoauth_clients_usage_summary())
	oauth_clients_usageCmd.Short = utils.GenerateCustomDescription(oauth_clients_usageCmd.Short, oauth_clients_usage_query.Description, oauth_clients_usage_summary.Description, )
	oauth_clients_usageCmd.Long = oauth_clients_usageCmd.Short
}
