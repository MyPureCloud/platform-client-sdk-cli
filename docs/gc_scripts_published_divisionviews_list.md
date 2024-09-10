## gc scripts published divisionviews list

Get the published scripts metadata.

### Synopsis

Get the published scripts metadata.

```
gc scripts published divisionviews list [flags]
```

### Options

```
  -a, --autopaginate               Automatically paginate through the results stripping page information
      --divisionIds string         Filters scripts to requested divisionIds
      --expand string              Expand
      --feature string             Feature filter
      --filtercondition string     Filter list command output based on a given condition or regular expression
      --flowId string              Secure flow id filter
  -h, --help                       help for list
      --name string                Name filter
      --pageNumber int             Page number (default 1)
      --pageSize int               Page size (default 25)
      --scriptDataVersion string   Advanced usage - controls the data version of the script
  -s, --stream                     Paginate and stream data as it is being processed leaving page information intact
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

* [gc scripts published divisionviews](gc_scripts_published_divisionviews.html)	 - /api/v2/scripts/published/divisionviews


