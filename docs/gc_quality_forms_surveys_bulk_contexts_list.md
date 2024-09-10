## gc quality forms surveys bulk contexts list

Retrieve a list of the latest form versions by context ids

### Synopsis

Retrieve a list of the latest form versions by context ids

```
gc quality forms surveys bulk contexts list [flags]
```

### Options

```
      --contextId strings   A comma-delimited list of valid survey form context ids. The maximum number of ids allowed in this list is 100. - REQUIRED
  -h, --help                help for list
      --published string    If true, the latest published version will be included. If false, only the unpublished version will be included. Valid values: true, false
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

* [gc quality forms surveys bulk contexts](gc_quality_forms_surveys_bulk_contexts.html)	 - /api/v2/quality/forms/surveys/bulk/contexts


