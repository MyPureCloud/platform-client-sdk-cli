## gc diagnostics logcapture browser entries query create

Query collected log entries. It returns a limited amount of records, to get all records use download endpoint.

### Synopsis

Query collected log entries. It returns a limited amount of records, to get all records use download endpoint.

```
gc diagnostics logcapture browser entries query create [flags]
```

### Options

```
      --after string       The cursor that points to the end of the set of entities that has been returned.
      --before string      The cursor that points to the start of the set of entities that has been returned.
  -d, --directory string   Directory path with files containing request bodies
  -f, --file string        File name containing the JSON body
  -h, --help               help for create
      --pageSize string    Number of entities to return. Maximum of 200.
  -b, --printrequestbody   Print the request body format of the API.
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

* [gc diagnostics logcapture browser entries query](gc_diagnostics_logcapture_browser_entries_query.html)	 - /api/v2/diagnostics/logcapture/browser/entries/query


