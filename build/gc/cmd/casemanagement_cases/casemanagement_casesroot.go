package casemanagement_cases

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_associations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_priority"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_datedue"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_summary"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_externalcontacts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_references"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_terminate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_stages"
)

func init() {
	casemanagement_casesCmd.AddCommand(casemanagement_cases_associations.Cmdcasemanagement_cases_associations())
	casemanagement_casesCmd.AddCommand(casemanagement_cases_priority.Cmdcasemanagement_cases_priority())
	casemanagement_casesCmd.AddCommand(casemanagement_cases_datedue.Cmdcasemanagement_cases_datedue())
	casemanagement_casesCmd.AddCommand(casemanagement_cases_summary.Cmdcasemanagement_cases_summary())
	casemanagement_casesCmd.AddCommand(casemanagement_cases_externalcontacts.Cmdcasemanagement_cases_externalcontacts())
	casemanagement_casesCmd.AddCommand(casemanagement_cases_references.Cmdcasemanagement_cases_references())
	casemanagement_casesCmd.AddCommand(casemanagement_cases_terminate.Cmdcasemanagement_cases_terminate())
	casemanagement_casesCmd.AddCommand(casemanagement_cases_stages.Cmdcasemanagement_cases_stages())
	casemanagement_casesCmd.Short = utils.GenerateCustomDescription(casemanagement_casesCmd.Short, casemanagement_cases_associations.Description, casemanagement_cases_priority.Description, casemanagement_cases_datedue.Description, casemanagement_cases_summary.Description, casemanagement_cases_externalcontacts.Description, casemanagement_cases_references.Description, casemanagement_cases_terminate.Description, casemanagement_cases_stages.Description, )
	casemanagement_casesCmd.Long = casemanagement_casesCmd.Short
}
