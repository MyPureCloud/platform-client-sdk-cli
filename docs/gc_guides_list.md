## gc guides list

Get all guides.

### Synopsis

Get all guides.

```
gc guides list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --name string              Filter by matching name - case insensitive.
      --nameContains string      Filter by name contains - case insensitive.
      --pageNumber string        Page number. (default "1")
      --pageSize string          Page size. The maximum page size is 100. (default "25")
      --sortBy string            Sort by. Default value dateModified. Valid values: dateModified, name, status
      --sortOrder string         Sort Order. Default value desc. Valid values: asc, desc
      --status string            Filter by status. Valid values: Published, Draft
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

* [gc guides](gc_guides.html)	 - /api/v2/guides


