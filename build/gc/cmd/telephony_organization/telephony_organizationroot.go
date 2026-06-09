package telephony_organization

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_organization_link"
)

func init() {
	telephony_organizationCmd.AddCommand(telephony_organization_link.Cmdtelephony_organization_link())
	telephony_organizationCmd.Short = utils.GenerateCustomDescription(telephony_organizationCmd.Short, telephony_organization_link.Description, )
	telephony_organizationCmd.Long = telephony_organizationCmd.Short
}
