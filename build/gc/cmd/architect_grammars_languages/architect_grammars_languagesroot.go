package architect_grammars_languages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_grammars_languages_files"
)

func init() {
	architect_grammars_languagesCmd.AddCommand(architect_grammars_languages_files.Cmdarchitect_grammars_languages_files())
	architect_grammars_languagesCmd.Short = utils.GenerateCustomDescription(architect_grammars_languagesCmd.Short, architect_grammars_languages_files.Description, )
	architect_grammars_languagesCmd.Long = architect_grammars_languagesCmd.Short
}
