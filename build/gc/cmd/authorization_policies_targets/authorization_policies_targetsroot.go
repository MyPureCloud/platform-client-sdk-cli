package authorization_policies_targets

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_policies_targets_subject"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_policies_targets_validate"
)

func init() {
	authorization_policies_targetsCmd.AddCommand(authorization_policies_targets_subject.Cmdauthorization_policies_targets_subject())
	authorization_policies_targetsCmd.AddCommand(authorization_policies_targets_validate.Cmdauthorization_policies_targets_validate())
	authorization_policies_targetsCmd.Short = utils.GenerateCustomDescription(authorization_policies_targetsCmd.Short, authorization_policies_targets_subject.Description, authorization_policies_targets_validate.Description, )
	authorization_policies_targetsCmd.Long = authorization_policies_targetsCmd.Short
}
