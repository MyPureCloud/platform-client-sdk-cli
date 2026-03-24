package casemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_caseplans"
)

func init() {
	casemanagementCmd.AddCommand(casemanagement_cases.Cmdcasemanagement_cases())
	casemanagementCmd.AddCommand(casemanagement_caseplans.Cmdcasemanagement_caseplans())
	casemanagementCmd.Short = utils.GenerateCustomDescription(casemanagementCmd.Short, casemanagement_cases.Description, casemanagement_caseplans.Description, )
	casemanagementCmd.Long = casemanagementCmd.Short
}
