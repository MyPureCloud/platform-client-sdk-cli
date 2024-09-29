package speechandtextanalytics_translations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics_translations_languages"
)

func init() {
	speechandtextanalytics_translationsCmd.AddCommand(speechandtextanalytics_translations_languages.Cmdspeechandtextanalytics_translations_languages())
	speechandtextanalytics_translationsCmd.Short = utils.GenerateCustomDescription(speechandtextanalytics_translationsCmd.Short, speechandtextanalytics_translations_languages.Description, )
	speechandtextanalytics_translationsCmd.Long = speechandtextanalytics_translationsCmd.Short
}
