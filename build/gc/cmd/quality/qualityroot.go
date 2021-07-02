package quality

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_publishedforms"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_calibrations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_evaluators"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_agents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_surveys"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_evaluations"
)

func init() {
	qualityCmd.AddCommand(quality_publishedforms.Cmdquality_publishedforms())
	qualityCmd.AddCommand(quality_calibrations.Cmdquality_calibrations())
	qualityCmd.AddCommand(quality_forms.Cmdquality_forms())
	qualityCmd.AddCommand(quality_evaluators.Cmdquality_evaluators())
	qualityCmd.AddCommand(quality_agents.Cmdquality_agents())
	qualityCmd.AddCommand(quality_conversations.Cmdquality_conversations())
	qualityCmd.AddCommand(quality_surveys.Cmdquality_surveys())
	qualityCmd.AddCommand(quality_evaluations.Cmdquality_evaluations())
	qualityCmd.Short = utils.GenerateCustomDescription(qualityCmd.Short, quality_publishedforms.Description, quality_calibrations.Description, quality_forms.Description, quality_evaluators.Description, quality_agents.Description, quality_conversations.Description, quality_surveys.Description, quality_evaluations.Description, )
	qualityCmd.Long = qualityCmd.Short
}
