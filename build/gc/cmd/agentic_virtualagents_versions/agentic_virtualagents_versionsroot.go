package agentic_virtualagents_versions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/agentic_virtualagents_versions_jobs"
)

func init() {
	agentic_virtualagents_versionsCmd.AddCommand(agentic_virtualagents_versions_jobs.Cmdagentic_virtualagents_versions_jobs())
	agentic_virtualagents_versionsCmd.Short = utils.GenerateCustomDescription(agentic_virtualagents_versionsCmd.Short, agentic_virtualagents_versions_jobs.Description, )
	agentic_virtualagents_versionsCmd.Long = agentic_virtualagents_versionsCmd.Short
}
