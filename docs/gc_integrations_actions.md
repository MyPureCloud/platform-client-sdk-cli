## gc integrations actions

/api/v2/integrations/actions

### Synopsis

/api/v2/integrations/actions

### Options

```
  -h, --help   help for actions
```

### Options inherited from parent commands

```
      --accesstoken string    accessToken override
      --clientid string       clientId override
      --clientsecret string   clientSecret override
      --environment string    environment override. E.g. mypurecloud.com.au or ap-southeast-2
  -i, --indicateprogress      Trace progress indicators to stderr
      --inputformat string    Data input format. Supported formats: YAML, JSON
      --outputformat string   Data output format. Supported formats: YAML, JSON
  -p, --profile string        Name of the profile to use for configuring the cli (default "DEFAULT")
      --transform string      Provide a Go template file for transforming output data
      --transformstr string   Provide a Go template string for transforming output data
```

### SEE ALSO

* [gc integrations](gc_integrations.html)	 - /api/v2/integrations
* [gc integrations actions categories](gc_integrations_actions_categories.html)	 - /api/v2/integrations/actions/categories
* [gc integrations actions certificates](gc_integrations_actions_certificates.html)	 - /api/v2/integrations/actions/certificates
* [gc integrations actions create](gc_integrations_actions_create.html)	 - Create a new Action. Not supported for `Function Integration` actions. Function integrations must be created as drafts to allow managing of uploading required ZIP function package before they may be used as a published action.
* [gc integrations actions delete](gc_integrations_actions_delete.html)	 - Delete an Action
* [gc integrations actions draft](gc_integrations_actions_draft.html)	 - /api/v2/integrations/actions/{actionId}/draft
* [gc integrations actions drafts](gc_integrations_actions_drafts.html)	 - /api/v2/integrations/actions/drafts
* [gc integrations actions execute](gc_integrations_actions_execute.html)	 - /api/v2/integrations/actions/{actionId}/execute
* [gc integrations actions function](gc_integrations_actions_function.html)	 - /api/v2/integrations/actions/{actionId}/function
* [gc integrations actions functions](gc_integrations_actions_functions.html)	 - /api/v2/integrations/actions/functions
* [gc integrations actions get](gc_integrations_actions_get.html)	 - Retrieves a single Action matching id.
* [gc integrations actions list](gc_integrations_actions_list.html)	 - Retrieves all actions associated with filters passed in via query param.
* [gc integrations actions schemas](gc_integrations_actions_schemas.html)	 - /api/v2/integrations/actions/{actionId}/schemas
* [gc integrations actions templates](gc_integrations_actions_templates.html)	 - /api/v2/integrations/actions/{actionId}/templates
* [gc integrations actions test](gc_integrations_actions_test.html)	 - /api/v2/integrations/actions/{actionId}/test
* [gc integrations actions update](gc_integrations_actions_update.html)	 - Patch an Action


