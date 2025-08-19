package routing_sms_identityresolution

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_sms_identityresolution_phonenumbers"
)

func init() {
	routing_sms_identityresolutionCmd.AddCommand(routing_sms_identityresolution_phonenumbers.Cmdrouting_sms_identityresolution_phonenumbers())
	routing_sms_identityresolutionCmd.Short = utils.GenerateCustomDescription(routing_sms_identityresolutionCmd.Short, routing_sms_identityresolution_phonenumbers.Description, )
	routing_sms_identityresolutionCmd.Long = routing_sms_identityresolutionCmd.Short
}
