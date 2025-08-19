package routing_assessments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_assessments_jobs"
)

func init() {
	routing_assessmentsCmd.AddCommand(routing_assessments_jobs.Cmdrouting_assessments_jobs())
	routing_assessmentsCmd.Short = utils.GenerateCustomDescription(routing_assessmentsCmd.Short, routing_assessments_jobs.Description, )
	routing_assessmentsCmd.Long = routing_assessmentsCmd.Short
}
