package authorization_subjects

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_subjects_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_subjects_bulkadd"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_subjects_bulkreplace"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_subjects_divisions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_subjects_bulkremove"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_subjects_rolecounts"
)

func init() {
	authorization_subjectsCmd.AddCommand(authorization_subjects_me.Cmdauthorization_subjects_me())
	authorization_subjectsCmd.AddCommand(authorization_subjects_bulkadd.Cmdauthorization_subjects_bulkadd())
	authorization_subjectsCmd.AddCommand(authorization_subjects_bulkreplace.Cmdauthorization_subjects_bulkreplace())
	authorization_subjectsCmd.AddCommand(authorization_subjects_divisions.Cmdauthorization_subjects_divisions())
	authorization_subjectsCmd.AddCommand(authorization_subjects_bulkremove.Cmdauthorization_subjects_bulkremove())
	authorization_subjectsCmd.AddCommand(authorization_subjects_rolecounts.Cmdauthorization_subjects_rolecounts())
	authorization_subjectsCmd.Short = utils.GenerateCustomDescription(authorization_subjectsCmd.Short, authorization_subjects_me.Description, authorization_subjects_bulkadd.Description, authorization_subjects_bulkreplace.Description, authorization_subjects_divisions.Description, authorization_subjects_bulkremove.Description, authorization_subjects_rolecounts.Description, )
	authorization_subjectsCmd.Long = authorization_subjectsCmd.Short
}
