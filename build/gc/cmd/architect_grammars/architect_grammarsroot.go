package architect_grammars

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_grammars_languages"
)

func init() {
	architect_grammarsCmd.AddCommand(architect_grammars_languages.Cmdarchitect_grammars_languages())
	architect_grammarsCmd.Short = utils.GenerateCustomDescription(architect_grammarsCmd.Short, architect_grammars_languages.Description, )
	architect_grammarsCmd.Long = architect_grammarsCmd.Short
}
