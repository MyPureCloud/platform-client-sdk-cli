package journey_externalevents_schemas

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_externalevents_schemas_coretypes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_externalevents_schemas_limits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_externalevents_schemas_versions"
)

func init() {
	journey_externalevents_schemasCmd.AddCommand(journey_externalevents_schemas_coretypes.Cmdjourney_externalevents_schemas_coretypes())
	journey_externalevents_schemasCmd.AddCommand(journey_externalevents_schemas_limits.Cmdjourney_externalevents_schemas_limits())
	journey_externalevents_schemasCmd.AddCommand(journey_externalevents_schemas_versions.Cmdjourney_externalevents_schemas_versions())
	journey_externalevents_schemasCmd.Short = utils.GenerateCustomDescription(journey_externalevents_schemasCmd.Short, journey_externalevents_schemas_coretypes.Description, journey_externalevents_schemas_limits.Description, journey_externalevents_schemas_versions.Description, )
	journey_externalevents_schemasCmd.Long = journey_externalevents_schemasCmd.Short
}
