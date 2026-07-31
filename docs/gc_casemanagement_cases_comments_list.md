## gc casemanagement cases comments list

Get comments for a Case.

### Synopsis

Get comments for a Case.

```
gc casemanagement cases comments list [caseId] [flags]
```

### Options

```
      --after string             Cursor pointing to the end of the previously returned page of comments.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageSize string          Number of comments to return. Maximum is 100.
      --sortOrder string         Ascending or descending sort order. Valid values: asc, desc
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

* [gc casemanagement cases comments](gc_casemanagement_cases_comments.html)	 - /api/v2/casemanagement/cases/{caseId}/comments


