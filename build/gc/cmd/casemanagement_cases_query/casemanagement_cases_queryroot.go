package casemanagement_cases_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_query_jobs"
)

func init() {
	casemanagement_cases_queryCmd.AddCommand(casemanagement_cases_query_jobs.Cmdcasemanagement_cases_query_jobs())
	casemanagement_cases_queryCmd.Short = utils.GenerateCustomDescription(casemanagement_cases_queryCmd.Short, casemanagement_cases_query_jobs.Description, )
	casemanagement_cases_queryCmd.Long = casemanagement_cases_queryCmd.Short
}
