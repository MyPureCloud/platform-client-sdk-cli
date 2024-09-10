## gc recording crossplatform mediaretentionpolicies list

Gets media retention policy list with query options to filter on name and enabled.

### Synopsis

Gets media retention policy list with query options to filter on name and enabled.

```
gc recording crossplatform mediaretentionpolicies list [flags]
```

### Options

```
  -a, --autopaginate              Automatically paginate through the results stripping page information
      --deleteDaysThreshold int   provides a way to fetch all policies with any actions having deleteDays exceeding the provided value
      --enabled string            checks to see if policy is enabled - use enabled = true or enabled = false Valid values: true, false
      --expand strings            variable name requested by expand list
      --filtercondition string    Filter list command output based on a given condition or regular expression
      --hasErrors string          provides a way to fetch all policies with errors or policies that do not have errors Valid values: true, false
  -h, --help                      help for list
      --name string               the policy name - used for filtering results in searches.
      --nextPage string           next page token
      --pageNumber int            The page number requested (default 1)
      --pageSize int              The total page size requested (default 25)
      --previousPage string       Previous page token
      --sortBy string             variable name requested to sort by
  -s, --stream                    Paginate and stream data as it is being processed leaving page information intact
      --summary string            provides a less verbose response of policy lists. Valid values: true, false
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

* [gc recording crossplatform mediaretentionpolicies](gc_recording_crossplatform_mediaretentionpolicies.html)	 - /api/v2/recording/crossplatform/mediaretentionpolicies


