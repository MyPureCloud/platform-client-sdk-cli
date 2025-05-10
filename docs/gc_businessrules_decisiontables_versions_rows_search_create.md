## gc businessrules decisiontables versions rows search create

Search for decision table rows

### Synopsis

Search for decision table rows

```
gc businessrules decisiontables versions rows search create [tableId] [tableVersion] [flags]
```

### Options

```
  -d, --directory string    Directory path with files containing request bodies
  -f, --file string         File name containing the JSON body
  -h, --help                help for create
      --pageNumber string   Page number of the entities to return. Defaults to 1.
      --pageSize string     Number of entities to return. Maximum of 100. Defaults to 25.
  -b, --printrequestbody    Print the request body format of the API.
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

* [gc businessrules decisiontables versions rows search](gc_businessrules_decisiontables_versions_rows_search.html)	 - /api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/search


