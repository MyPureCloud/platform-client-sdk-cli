package dataprivacy

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dataprivacy_maskingrules"
)

func init() {
	dataprivacyCmd.AddCommand(dataprivacy_maskingrules.Cmddataprivacy_maskingrules())
	dataprivacyCmd.Short = utils.GenerateCustomDescription(dataprivacyCmd.Short, dataprivacy_maskingrules.Description, )
	dataprivacyCmd.Long = dataprivacyCmd.Short
}
