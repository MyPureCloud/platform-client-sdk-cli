package agentic_virtualagents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/agentic_virtualagents_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/agentic_virtualagents_versions"
)

func init() {
	agentic_virtualagentsCmd.AddCommand(agentic_virtualagents_jobs.Cmdagentic_virtualagents_jobs())
	agentic_virtualagentsCmd.AddCommand(agentic_virtualagents_versions.Cmdagentic_virtualagents_versions())
	agentic_virtualagentsCmd.Short = utils.GenerateCustomDescription(agentic_virtualagentsCmd.Short, agentic_virtualagents_jobs.Description, agentic_virtualagents_versions.Description, )
	agentic_virtualagentsCmd.Long = agentic_virtualagentsCmd.Short
}
