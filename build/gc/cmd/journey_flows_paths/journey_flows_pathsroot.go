package journey_flows_paths

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_flows_paths_query"
)

func init() {
	journey_flows_pathsCmd.AddCommand(journey_flows_paths_query.Cmdjourney_flows_paths_query())
	journey_flows_pathsCmd.Short = utils.GenerateCustomDescription(journey_flows_pathsCmd.Short, journey_flows_paths_query.Description, )
	journey_flows_pathsCmd.Long = journey_flows_pathsCmd.Short
}
