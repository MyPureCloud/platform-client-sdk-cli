package speechandtextanalytics_programs_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_programs_settings_insights"
)

func init() {
	speechandtextanalytics_programs_settingsCmd.AddCommand(speechandtextanalytics_programs_settings_insights.Cmdspeechandtextanalytics_programs_settings_insights())
	speechandtextanalytics_programs_settingsCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_programs_settingsCmd.Short, speechandtextanalytics_programs_settings_insights.Description, )
	speechandtextanalytics_programs_settingsCmd.Long = speechandtextanalytics_programs_settingsCmd.Short
}
