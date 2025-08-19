package gamification_contests_uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_contests_uploads_prizeimages"
)

func init() {
	gamification_contests_uploadsCmd.AddCommand(gamification_contests_uploads_prizeimages.Cmdgamification_contests_uploads_prizeimages())
	gamification_contests_uploadsCmd.Short = utils.GenerateCustomDescription(gamification_contests_uploadsCmd.Short, gamification_contests_uploads_prizeimages.Description, )
	gamification_contests_uploadsCmd.Long = gamification_contests_uploadsCmd.Short
}
