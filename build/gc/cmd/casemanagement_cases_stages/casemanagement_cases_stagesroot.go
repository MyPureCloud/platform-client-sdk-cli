package casemanagement_cases_stages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_stages_steps"
)

func init() {
	casemanagement_cases_stagesCmd.AddCommand(casemanagement_cases_stages_steps.Cmdcasemanagement_cases_stages_steps())
	casemanagement_cases_stagesCmd.Short = utils.GenerateCustomDescription(casemanagement_cases_stagesCmd.Short, casemanagement_cases_stages_steps.Description, )
	casemanagement_cases_stagesCmd.Long = casemanagement_cases_stagesCmd.Short
}
