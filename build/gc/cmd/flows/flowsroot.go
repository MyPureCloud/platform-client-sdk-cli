package flows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_datatables"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_actions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_executions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_milestones"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_outcomes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_latestconfiguration"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_history"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_divisionviews"
)

func init() {
	flowsCmd.AddCommand(flows_datatables.Cmdflows_datatables())
	flowsCmd.AddCommand(flows_jobs.Cmdflows_jobs())
	flowsCmd.AddCommand(flows_actions.Cmdflows_actions())
	flowsCmd.AddCommand(flows_executions.Cmdflows_executions())
	flowsCmd.AddCommand(flows_milestones.Cmdflows_milestones())
	flowsCmd.AddCommand(flows_outcomes.Cmdflows_outcomes())
	flowsCmd.AddCommand(flows_latestconfiguration.Cmdflows_latestconfiguration())
	flowsCmd.AddCommand(flows_history.Cmdflows_history())
	flowsCmd.AddCommand(flows_versions.Cmdflows_versions())
	flowsCmd.AddCommand(flows_divisionviews.Cmdflows_divisionviews())
	flowsCmd.Short = utils.GenerateCustomDescription(flowsCmd.Short, flows_datatables.Description, flows_jobs.Description, flows_actions.Description, flows_executions.Description, flows_milestones.Description, flows_outcomes.Description, flows_latestconfiguration.Description, flows_history.Description, flows_versions.Description, flows_divisionviews.Description, )
	flowsCmd.Long = flowsCmd.Short
}
