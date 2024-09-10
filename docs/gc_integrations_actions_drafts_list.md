## gc integrations actions drafts list

Retrieves all action drafts associated with the filters passed in via query param.

### Synopsis

Retrieves all action drafts associated with the filters passed in via query param.

```
gc integrations actions drafts list [flags]
```

### Options

```
  -a, --autopaginate                Automatically paginate through the results stripping page information
      --category string             Filter by category name.
      --filtercondition string      Filter list command output based on a given condition or regular expression
  -h, --help                        help for list
      --ids string                  Filter by action Id. Can be a comma separated list to request multiple actions.  Limit of 50 Ids.
      --includeAuthActions string   Whether or not to include authentication actions in the response. These actions are not directly executable. Some integrations create them and will run them as needed to refresh authentication information for other actions. Valid values: true, false
      --name string                 Filter by partial or complete action name.
      --nextPage string             next page token
      --pageNumber int              The page number requested (default 1)
      --pageSize int                The total page size requested (default 25)
      --previousPage string         Previous page token
      --secure secure               Filter based on secure configuration option. True will only return actions marked as secure. False will return only non-secure actions. Do not use filter if you want all Actions. Valid values: true, false
      --sortBy string               Root level field name to sort on.
      --sortOrder sortBy            Direction to sort sortBy field. Valid values: ASC, DESC
  -s, --stream                      Paginate and stream data as it is being processed leaving page information intact
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

* [gc integrations actions drafts](gc_integrations_actions_drafts.html)	 - /api/v2/integrations/actions/drafts


