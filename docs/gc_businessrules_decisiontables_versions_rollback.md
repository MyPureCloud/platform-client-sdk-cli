## gc businessrules decisiontables versions rollback

/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rollback

### Synopsis

/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rollback

### Options

```
  -h, --help   help for rollback
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

* [gc businessrules decisiontables versions](gc_businessrules_decisiontables_versions.html)	 - /api/v2/businessrules/decisiontables/{tableId}/versions
* [gc businessrules decisiontables versions rollback create](gc_businessrules_decisiontables_versions_rollback_create.html)	 - Re-publish a superseded decision table version as the current published version


