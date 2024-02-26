package journey_flows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_flows_paths"
)

func init() {
	journey_flowsCmd.AddCommand(journey_flows_paths.Cmdjourney_flows_paths())
	journey_flowsCmd.Short = utils.GenerateCustomDescription(journey_flowsCmd.Short, journey_flows_paths.Description, )
	journey_flowsCmd.Long = journey_flowsCmd.Short
}
