package workforcemanagement_alternativeshifts_offers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_alternativeshifts_offers_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_alternativeshifts_offers_search"
)

func init() {
	workforcemanagement_alternativeshifts_offersCmd.AddCommand(workforcemanagement_alternativeshifts_offers_jobs.Cmdworkforcemanagement_alternativeshifts_offers_jobs())
	workforcemanagement_alternativeshifts_offersCmd.AddCommand(workforcemanagement_alternativeshifts_offers_search.Cmdworkforcemanagement_alternativeshifts_offers_search())
	workforcemanagement_alternativeshifts_offersCmd.Short = utils.GenerateCustomDescription(workforcemanagement_alternativeshifts_offersCmd.Short, workforcemanagement_alternativeshifts_offers_jobs.Description, workforcemanagement_alternativeshifts_offers_search.Description, )
	workforcemanagement_alternativeshifts_offersCmd.Long = workforcemanagement_alternativeshifts_offersCmd.Short
}