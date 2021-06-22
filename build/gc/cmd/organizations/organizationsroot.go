package organizations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_ipaddressauthentication"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_features"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_embeddedintegration"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_limits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_whitelist"
)

func init() {
	organizationsCmd.AddCommand(organizations_ipaddressauthentication.Cmdorganizations_ipaddressauthentication())
	organizationsCmd.AddCommand(organizations_features.Cmdorganizations_features())
	organizationsCmd.AddCommand(organizations_me.Cmdorganizations_me())
	organizationsCmd.AddCommand(organizations_embeddedintegration.Cmdorganizations_embeddedintegration())
	organizationsCmd.AddCommand(organizations_limits.Cmdorganizations_limits())
	organizationsCmd.AddCommand(organizations_whitelist.Cmdorganizations_whitelist())
	organizationsCmd.Short = utils.GenerateCustomDescription(organizationsCmd.Short, organizations_ipaddressauthentication.Description, organizations_features.Description, organizations_me.Description, organizations_embeddedintegration.Description, organizations_limits.Description, organizations_whitelist.Description, )
	organizationsCmd.Long = organizationsCmd.Short
}
