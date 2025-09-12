package languageunderstanding_ignoretopics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_ignoretopics_remove"
)

func init() {
	languageunderstanding_ignoretopicsCmd.AddCommand(languageunderstanding_ignoretopics_remove.Cmdlanguageunderstanding_ignoretopics_remove())
	languageunderstanding_ignoretopicsCmd.Short = utils.GenerateCustomDescription(languageunderstanding_ignoretopicsCmd.Short, languageunderstanding_ignoretopics_remove.Description, )
	languageunderstanding_ignoretopicsCmd.Long = languageunderstanding_ignoretopicsCmd.Short
}
