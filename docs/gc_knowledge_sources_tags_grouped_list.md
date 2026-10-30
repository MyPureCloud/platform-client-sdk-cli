## gc knowledge sources tags grouped list

Get tags grouped by source for the organization

### Synopsis

Get tags grouped by source for the organization

```
gc knowledge sources tags grouped list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber string        Page number (default "1")
      --pageSize string          Number of items per page. Maximum is 200. (default "100")
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --tagName string           Optional substring filter for tag names.
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

* [gc knowledge sources tags grouped](gc_knowledge_sources_tags_grouped.html)	 - /api/v2/knowledge/sources/tags/grouped


