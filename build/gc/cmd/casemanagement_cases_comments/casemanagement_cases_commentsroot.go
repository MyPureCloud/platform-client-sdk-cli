package casemanagement_cases_comments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/casemanagement_cases_comments_me"
)

func init() {
	casemanagement_cases_commentsCmd.AddCommand(casemanagement_cases_comments_me.Cmdcasemanagement_cases_comments_me())
	casemanagement_cases_commentsCmd.Short = utils.GenerateCustomDescription(casemanagement_cases_commentsCmd.Short, casemanagement_cases_comments_me.Description, )
	casemanagement_cases_commentsCmd.Long = casemanagement_cases_commentsCmd.Short
}
