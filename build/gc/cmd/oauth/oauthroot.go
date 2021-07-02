package oauth

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/oauth_scopes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/oauth_clients"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/oauth_authorizations"
)

func init() {
	oauthCmd.AddCommand(oauth_scopes.Cmdoauth_scopes())
	oauthCmd.AddCommand(oauth_clients.Cmdoauth_clients())
	oauthCmd.AddCommand(oauth_authorizations.Cmdoauth_authorizations())
	oauthCmd.Short = utils.GenerateCustomDescription(oauthCmd.Short, oauth_scopes.Description, oauth_clients.Description, oauth_authorizations.Description, )
	oauthCmd.Long = oauthCmd.Short
}
