## gc businessrules decisiontables versions rows list

Get a list of decision table rows.

### Synopsis

Get a list of decision table rows.

```
gc businessrules decisiontables versions rows list [tableId] [tableVersion] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber string        Page number of the entities to return. Defaults to 1.
      --pageSize string          Number of entities to return. Maximum of 100. Defaults to 25.
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

* [gc businessrules decisiontables versions rows](gc_businessrules_decisiontables_versions_rows.html)	 - /api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows


