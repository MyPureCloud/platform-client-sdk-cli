package casemanagement_cases_query_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_query_jobs_results"
)

func init() {
	casemanagement_cases_query_jobsCmd.AddCommand(casemanagement_cases_query_jobs_results.Cmdcasemanagement_cases_query_jobs_results())
	casemanagement_cases_query_jobsCmd.Short = utils.GenerateCustomDescription(casemanagement_cases_query_jobsCmd.Short, casemanagement_cases_query_jobs_results.Description, )
	casemanagement_cases_query_jobsCmd.Long = casemanagement_cases_query_jobsCmd.Short
}
