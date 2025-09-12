package languageunderstanding_ignorephrases

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_ignorephrases_remove"
)

func init() {
	languageunderstanding_ignorephrasesCmd.AddCommand(languageunderstanding_ignorephrases_remove.Cmdlanguageunderstanding_ignorephrases_remove())
	languageunderstanding_ignorephrasesCmd.Short = utils.GenerateCustomDescription(languageunderstanding_ignorephrasesCmd.Short, languageunderstanding_ignorephrases_remove.Description, )
	languageunderstanding_ignorephrasesCmd.Long = languageunderstanding_ignorephrasesCmd.Short
}
