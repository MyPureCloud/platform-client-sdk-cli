## gc learning modules list

Get all learning modules of an organization

### Synopsis

Get all learning modules of an organization

```
gc learning modules list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand strings           Fields to expand in response(case insensitive) Valid values: rule, summaryData
      --externalIds strings      Specifies the module external IDs to filter by. Only one ID is allowed
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --isArchived string        Archive status Valid values: true, false
      --isPublished string       Specifies if only the Unpublished (isPublished is False) or Published (isPublished is True) modules are returned. If isPublished is Any or omitted, both types are returned Valid values: True, False, Any
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "25")
      --searchTerm string        Search Term (searchable by name)
      --sortBy string            Sort by Valid values: name, createddate, percentpassed, averagescore
      --sortOrder string         Sort order Valid values: ascending, descending
      --statuses strings         Specifies the module statuses to filter by Valid values: Unpublished, Published, Archived
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --types strings            Specifies the module types. Informational, AssessedContent and Assessment are deprecated Valid values: Informational, AssessedContent, Assessment, External, Native
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

* [gc learning modules](gc_learning_modules.html)	 - /api/v2/learning/modules


