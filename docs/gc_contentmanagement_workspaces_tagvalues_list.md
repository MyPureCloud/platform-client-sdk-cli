## gc contentmanagement workspaces tagvalues list

Get a list of workspace tags

### Synopsis

Get a list of workspace tags

```
gc contentmanagement workspaces tagvalues list [workspaceId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand strings           Which fields, if any, to expand. Valid values: acl
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber int           Page number (default 1)
      --pageSize int             Page size (default 25)
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --value string             filter the list of tags returned
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

* [gc contentmanagement workspaces tagvalues](gc_contentmanagement_workspaces_tagvalues.html)	 - /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues


