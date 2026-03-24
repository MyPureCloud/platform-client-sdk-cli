package casemanagement_cases_associations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_associations_query"
)

func init() {
	casemanagement_cases_associationsCmd.AddCommand(casemanagement_cases_associations_query.Cmdcasemanagement_cases_associations_query())
	casemanagement_cases_associationsCmd.Short = utils.GenerateCustomDescription(casemanagement_cases_associationsCmd.Short, casemanagement_cases_associations_query.Description, )
	casemanagement_cases_associationsCmd.Long = casemanagement_cases_associationsCmd.Short
}
