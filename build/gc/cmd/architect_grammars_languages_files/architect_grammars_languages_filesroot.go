package architect_grammars_languages_files

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_grammars_languages_files_voice"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_grammars_languages_files_dtmf"
)

func init() {
	architect_grammars_languages_filesCmd.AddCommand(architect_grammars_languages_files_voice.Cmdarchitect_grammars_languages_files_voice())
	architect_grammars_languages_filesCmd.AddCommand(architect_grammars_languages_files_dtmf.Cmdarchitect_grammars_languages_files_dtmf())
	architect_grammars_languages_filesCmd.Short = utils.GenerateCustomDescription(architect_grammars_languages_filesCmd.Short, architect_grammars_languages_files_voice.Description, architect_grammars_languages_files_dtmf.Description, )
	architect_grammars_languages_filesCmd.Long = architect_grammars_languages_filesCmd.Short
}
