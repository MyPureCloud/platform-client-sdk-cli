package dataprivacy_maskingrules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dataprivacy_maskingrules_validate"
)

func init() {
	dataprivacy_maskingrulesCmd.AddCommand(dataprivacy_maskingrules_validate.Cmddataprivacy_maskingrules_validate())
	dataprivacy_maskingrulesCmd.Short = utils.GenerateCustomDescription(dataprivacy_maskingrulesCmd.Short, dataprivacy_maskingrules_validate.Description, )
	dataprivacy_maskingrulesCmd.Long = dataprivacy_maskingrulesCmd.Short
}
