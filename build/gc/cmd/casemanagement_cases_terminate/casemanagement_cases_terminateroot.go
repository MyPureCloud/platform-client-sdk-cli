package casemanagement_cases_terminate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_terminate_jobs"
)

func init() {
	casemanagement_cases_terminateCmd.AddCommand(casemanagement_cases_terminate_jobs.Cmdcasemanagement_cases_terminate_jobs())
	casemanagement_cases_terminateCmd.Short = utils.GenerateCustomDescription(casemanagement_cases_terminateCmd.Short, casemanagement_cases_terminate_jobs.Description, )
	casemanagement_cases_terminateCmd.Long = casemanagement_cases_terminateCmd.Short
}
