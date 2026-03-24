package casemanagement_caseplans

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_caseplans_publish"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_caseplans_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_caseplans_stageplans"
)

func init() {
	casemanagement_caseplansCmd.AddCommand(casemanagement_caseplans_publish.Cmdcasemanagement_caseplans_publish())
	casemanagement_caseplansCmd.AddCommand(casemanagement_caseplans_versions.Cmdcasemanagement_caseplans_versions())
	casemanagement_caseplansCmd.AddCommand(casemanagement_caseplans_stageplans.Cmdcasemanagement_caseplans_stageplans())
	casemanagement_caseplansCmd.Short = utils.GenerateCustomDescription(casemanagement_caseplansCmd.Short, casemanagement_caseplans_publish.Description, casemanagement_caseplans_versions.Description, casemanagement_caseplans_stageplans.Description, )
	casemanagement_caseplansCmd.Long = casemanagement_caseplansCmd.Short
}
