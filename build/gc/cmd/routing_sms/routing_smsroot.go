package routing_sms

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_sms_addresses"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_sms_phonenumbers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_sms_availablephonenumbers"
)

func init() {
	routing_smsCmd.AddCommand(routing_sms_addresses.Cmdrouting_sms_addresses())
	routing_smsCmd.AddCommand(routing_sms_phonenumbers.Cmdrouting_sms_phonenumbers())
	routing_smsCmd.AddCommand(routing_sms_availablephonenumbers.Cmdrouting_sms_availablephonenumbers())
	routing_smsCmd.Short = utils.GenerateCustomDescription(routing_smsCmd.Short, routing_sms_addresses.Description, routing_sms_phonenumbers.Description, routing_sms_availablephonenumbers.Description, )
	routing_smsCmd.Long = routing_smsCmd.Short
}
