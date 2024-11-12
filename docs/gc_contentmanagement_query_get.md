## gc contentmanagement query get

Query content

### Synopsis

Query content

```
gc contentmanagement query get [flags]
```

### Options

```
      --expand strings       Which fields, if any, to expand. Valid values: acl, workspace
  -h, --help                 help for get
      --pageNumber string    Page number (default "1")
      --pageSize string      Page size (default "25")
      --queryPhrase string   Phrase tokens are ANDed together over all searchable fields - REQUIRED
      --sortBy string        name or dateCreated
      --sortOrder string     ascending or descending
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

* [gc contentmanagement query](gc_contentmanagement_query.html)	 - /api/v2/contentmanagement/query


