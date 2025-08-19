package speechandtextanalytics_programs_general

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_programs_general_jobs"
)

func init() {
	speechandtextanalytics_programs_generalCmd.AddCommand(speechandtextanalytics_programs_general_jobs.Cmdspeechandtextanalytics_programs_general_jobs())
	speechandtextanalytics_programs_generalCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_programs_generalCmd.Short, speechandtextanalytics_programs_general_jobs.Description, )
	speechandtextanalytics_programs_generalCmd.Long = speechandtextanalytics_programs_generalCmd.Short
}
