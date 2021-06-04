package languages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languages_translations"
)

func init() {
	languagesCmd.AddCommand(languages_translations.Cmdlanguages_translations())
	languagesCmd.Short = utils.GenerateCustomDescription(languagesCmd.Short, languages_translations.Description, )
	languagesCmd.Long = languagesCmd.Short
}
