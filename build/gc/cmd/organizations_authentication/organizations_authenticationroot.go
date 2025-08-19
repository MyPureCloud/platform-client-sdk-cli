package organizations_authentication

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_authentication_settings"
)

func init() {
	organizations_authenticationCmd.AddCommand(organizations_authentication_settings.Cmdorganizations_authentication_settings())
	organizations_authenticationCmd.Short = utils.GenerateCustomDescription(organizations_authenticationCmd.Short, organizations_authentication_settings.Description, )
	organizations_authenticationCmd.Long = organizations_authenticationCmd.Short
}
