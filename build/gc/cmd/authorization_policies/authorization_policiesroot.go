package authorization_policies

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_policies_targets"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_policies_simulate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_policies_subject"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_policies_attributes"
)

func init() {
	authorization_policiesCmd.AddCommand(authorization_policies_targets.Cmdauthorization_policies_targets())
	authorization_policiesCmd.AddCommand(authorization_policies_simulate.Cmdauthorization_policies_simulate())
	authorization_policiesCmd.AddCommand(authorization_policies_subject.Cmdauthorization_policies_subject())
	authorization_policiesCmd.AddCommand(authorization_policies_attributes.Cmdauthorization_policies_attributes())
	authorization_policiesCmd.Short = utils.GenerateCustomDescription(authorization_policiesCmd.Short, authorization_policies_targets.Description, authorization_policies_simulate.Description, authorization_policies_subject.Description, authorization_policies_attributes.Description, )
	authorization_policiesCmd.Long = authorization_policiesCmd.Short
}
