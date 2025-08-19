package routing_sms_phonenumbers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_sms_phonenumbers_import"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_sms_phonenumbers_alphanumeric"
)

func init() {
	routing_sms_phonenumbersCmd.AddCommand(routing_sms_phonenumbers_import.Cmdrouting_sms_phonenumbers_import())
	routing_sms_phonenumbersCmd.AddCommand(routing_sms_phonenumbers_alphanumeric.Cmdrouting_sms_phonenumbers_alphanumeric())
	routing_sms_phonenumbersCmd.Short = utils.GenerateCustomDescription(routing_sms_phonenumbersCmd.Short, routing_sms_phonenumbers_import.Description, routing_sms_phonenumbers_alphanumeric.Description, )
	routing_sms_phonenumbersCmd.Long = routing_sms_phonenumbersCmd.Short
}
