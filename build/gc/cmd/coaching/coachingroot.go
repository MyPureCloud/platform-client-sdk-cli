package coaching

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_appointments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_notifications"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_scheduleslots"
)

func init() {
	coachingCmd.AddCommand(coaching_appointments.Cmdcoaching_appointments())
	coachingCmd.AddCommand(coaching_notifications.Cmdcoaching_notifications())
	coachingCmd.AddCommand(coaching_scheduleslots.Cmdcoaching_scheduleslots())
	coachingCmd.Short = utils.GenerateCustomDescription(coachingCmd.Short, coaching_appointments.Description, coaching_notifications.Description, coaching_scheduleslots.Description, )
	coachingCmd.Long = coachingCmd.Short
}
