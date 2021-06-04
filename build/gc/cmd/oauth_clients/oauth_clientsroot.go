package oauth_clients

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/oauth_clients_secret"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/oauth_clients_usage"
)

func init() {
	oauth_clientsCmd.AddCommand(oauth_clients_secret.Cmdoauth_clients_secret())
	oauth_clientsCmd.AddCommand(oauth_clients_usage.Cmdoauth_clients_usage())
	oauth_clientsCmd.Short = utils.GenerateCustomDescription(oauth_clientsCmd.Short, oauth_clients_secret.Description, oauth_clients_usage.Description, )
	oauth_clientsCmd.Long = oauth_clientsCmd.Short
}
