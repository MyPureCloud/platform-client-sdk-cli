package contentmanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_status"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_workspaces"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_documents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_auditquery"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_usage"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_shares"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_shared"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_securityprofiles"
)

func init() {
	contentmanagementCmd.AddCommand(contentmanagement_status.Cmdcontentmanagement_status())
	contentmanagementCmd.AddCommand(contentmanagement_workspaces.Cmdcontentmanagement_workspaces())
	contentmanagementCmd.AddCommand(contentmanagement_documents.Cmdcontentmanagement_documents())
	contentmanagementCmd.AddCommand(contentmanagement_auditquery.Cmdcontentmanagement_auditquery())
	contentmanagementCmd.AddCommand(contentmanagement_usage.Cmdcontentmanagement_usage())
	contentmanagementCmd.AddCommand(contentmanagement_query.Cmdcontentmanagement_query())
	contentmanagementCmd.AddCommand(contentmanagement_shares.Cmdcontentmanagement_shares())
	contentmanagementCmd.AddCommand(contentmanagement_shared.Cmdcontentmanagement_shared())
	contentmanagementCmd.AddCommand(contentmanagement_securityprofiles.Cmdcontentmanagement_securityprofiles())
	contentmanagementCmd.Short = utils.GenerateCustomDescription(contentmanagementCmd.Short, contentmanagement_status.Description, contentmanagement_workspaces.Description, contentmanagement_documents.Description, contentmanagement_auditquery.Description, contentmanagement_usage.Description, contentmanagement_query.Description, contentmanagement_shares.Description, contentmanagement_shared.Description, contentmanagement_securityprofiles.Description, )
	contentmanagementCmd.Long = contentmanagementCmd.Short
}
