package outbound_campaigns_all

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_all_divisionviews"
)

func init() {
	outbound_campaigns_allCmd.AddCommand(outbound_campaigns_all_divisionviews.Cmdoutbound_campaigns_all_divisionviews())
	outbound_campaigns_allCmd.Short = utils.GenerateCustomDescription(outbound_campaigns_allCmd.Short, outbound_campaigns_all_divisionviews.Description, )
	outbound_campaigns_allCmd.Long = outbound_campaigns_allCmd.Short
}
