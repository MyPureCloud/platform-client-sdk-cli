package externalcontacts_contacts_journey

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_contacts_journey_sessions"
)

func init() {
	externalcontacts_contacts_journeyCmd.AddCommand(externalcontacts_contacts_journey_sessions.Cmdexternalcontacts_contacts_journey_sessions())
	externalcontacts_contacts_journeyCmd.Short = utils.GenerateCustomDescription(externalcontacts_contacts_journeyCmd.Short, externalcontacts_contacts_journey_sessions.Description, )
	externalcontacts_contacts_journeyCmd.Long = externalcontacts_contacts_journeyCmd.Short
}
