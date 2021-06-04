package coaching

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_scheduleslots"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_appointments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_notifications"
)

func init() {
	coachingCmd.AddCommand(coaching_scheduleslots.Cmdcoaching_scheduleslots())
	coachingCmd.AddCommand(coaching_appointments.Cmdcoaching_appointments())
	coachingCmd.AddCommand(coaching_notifications.Cmdcoaching_notifications())
	coachingCmd.Short = utils.GenerateCustomDescription(coachingCmd.Short, coaching_scheduleslots.Description, coaching_appointments.Description, coaching_notifications.Description, )
	coachingCmd.Long = coachingCmd.Short
}
