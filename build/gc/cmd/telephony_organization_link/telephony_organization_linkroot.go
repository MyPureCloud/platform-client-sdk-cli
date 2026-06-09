package telephony_organization_link

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_organization_link_approve"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_organization_link_regions"
)

func init() {
	telephony_organization_linkCmd.AddCommand(telephony_organization_link_approve.Cmdtelephony_organization_link_approve())
	telephony_organization_linkCmd.AddCommand(telephony_organization_link_regions.Cmdtelephony_organization_link_regions())
	telephony_organization_linkCmd.Short = utils.GenerateCustomDescription(telephony_organization_linkCmd.Short, telephony_organization_link_approve.Description, telephony_organization_link_regions.Description, )
	telephony_organization_linkCmd.Long = telephony_organization_linkCmd.Short
}
