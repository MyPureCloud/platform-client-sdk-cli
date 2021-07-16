package languageunderstanding_miners

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_miners_drafts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_miners_intents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_miners_execute"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_miners_uploads"
)

func init() {
	languageunderstanding_minersCmd.AddCommand(languageunderstanding_miners_drafts.Cmdlanguageunderstanding_miners_drafts())
	languageunderstanding_minersCmd.AddCommand(languageunderstanding_miners_intents.Cmdlanguageunderstanding_miners_intents())
	languageunderstanding_minersCmd.AddCommand(languageunderstanding_miners_execute.Cmdlanguageunderstanding_miners_execute())
	languageunderstanding_minersCmd.AddCommand(languageunderstanding_miners_uploads.Cmdlanguageunderstanding_miners_uploads())
	languageunderstanding_minersCmd.Short = utils.GenerateCustomDescription(languageunderstanding_minersCmd.Short, languageunderstanding_miners_drafts.Description, languageunderstanding_miners_intents.Description, languageunderstanding_miners_execute.Description, languageunderstanding_miners_uploads.Description, )
	languageunderstanding_minersCmd.Long = languageunderstanding_minersCmd.Short
}
