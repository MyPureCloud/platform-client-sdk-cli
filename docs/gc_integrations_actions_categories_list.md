## gc integrations actions categories list

Retrieves all categories of available Actions

### Synopsis

Retrieves all categories of available Actions

```
gc integrations actions categories list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --nextPage string          next page token
      --pageNumber int           The page number requested (default 1)
      --pageSize int             The total page size requested (default 25)
      --previousPage string      Previous page token
      --secure string            Filter to only include secure actions. True will only include actions marked secured. False will include only unsecure actions. Do not use filter if you want all Actions. Valid values: true, false
      --sortBy name              Root level field name to sort on.  Only name is supported on this endpoint.
      --sortOrder sortBy         Direction to sort sortBy field. Valid values: ASC, DESC
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
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

* [gc integrations actions categories](gc_integrations_actions_categories.html)	 - /api/v2/integrations/actions/categories


