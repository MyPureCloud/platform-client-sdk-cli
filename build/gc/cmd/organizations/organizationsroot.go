package organizations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_embeddedintegration"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_ipaddressauthentication"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_whitelist"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_authentication"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_features"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_limits"
)

func init() {
	organizationsCmd.AddCommand(organizations_embeddedintegration.Cmdorganizations_embeddedintegration())
	organizationsCmd.AddCommand(organizations_ipaddressauthentication.Cmdorganizations_ipaddressauthentication())
	organizationsCmd.AddCommand(organizations_whitelist.Cmdorganizations_whitelist())
	organizationsCmd.AddCommand(organizations_authentication.Cmdorganizations_authentication())
	organizationsCmd.AddCommand(organizations_features.Cmdorganizations_features())
	organizationsCmd.AddCommand(organizations_me.Cmdorganizations_me())
	organizationsCmd.AddCommand(organizations_limits.Cmdorganizations_limits())
	organizationsCmd.Short = utils.GenerateCustomDescription(organizationsCmd.Short, organizations_embeddedintegration.Description, organizations_ipaddressauthentication.Description, organizations_whitelist.Description, organizations_authentication.Description, organizations_features.Description, organizations_me.Description, organizations_limits.Description, )
	organizationsCmd.Long = organizationsCmd.Short
}
