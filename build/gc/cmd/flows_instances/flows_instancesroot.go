package flows_instances

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_instances_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_instances_querycapabilities"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_instances_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_instances_settings"
)

func init() {
	flows_instancesCmd.AddCommand(flows_instances_jobs.Cmdflows_instances_jobs())
	flows_instancesCmd.AddCommand(flows_instances_querycapabilities.Cmdflows_instances_querycapabilities())
	flows_instancesCmd.AddCommand(flows_instances_query.Cmdflows_instances_query())
	flows_instancesCmd.AddCommand(flows_instances_settings.Cmdflows_instances_settings())
	flows_instancesCmd.Short = utils.GenerateCustomDescription(flows_instancesCmd.Short, flows_instances_jobs.Description, flows_instances_querycapabilities.Description, flows_instances_query.Description, flows_instances_settings.Description, )
	flows_instancesCmd.Long = flows_instancesCmd.Short
}
