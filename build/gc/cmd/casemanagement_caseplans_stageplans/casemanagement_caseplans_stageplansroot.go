package casemanagement_caseplans_stageplans

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_caseplans_stageplans_stepplans"
)

func init() {
	casemanagement_caseplans_stageplansCmd.AddCommand(casemanagement_caseplans_stageplans_stepplans.Cmdcasemanagement_caseplans_stageplans_stepplans())
	casemanagement_caseplans_stageplansCmd.Short = utils.GenerateCustomDescription(casemanagement_caseplans_stageplansCmd.Short, casemanagement_caseplans_stageplans_stepplans.Description, )
	casemanagement_caseplans_stageplansCmd.Long = casemanagement_caseplans_stageplansCmd.Short
}
