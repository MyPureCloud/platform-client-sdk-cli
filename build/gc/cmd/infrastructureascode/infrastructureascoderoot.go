package infrastructureascode

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/infrastructureascode_accelerators"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/infrastructureascode_jobs"
)

func init() {
	infrastructureascodeCmd.AddCommand(infrastructureascode_accelerators.Cmdinfrastructureascode_accelerators())
	infrastructureascodeCmd.AddCommand(infrastructureascode_jobs.Cmdinfrastructureascode_jobs())
	infrastructureascodeCmd.Short = utils.GenerateCustomDescription(infrastructureascodeCmd.Short, infrastructureascode_accelerators.Description, infrastructureascode_jobs.Description, )
	infrastructureascodeCmd.Long = infrastructureascodeCmd.Short
}
