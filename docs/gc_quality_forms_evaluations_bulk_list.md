## gc quality forms evaluations bulk list

Retrieve a list of evaluation forms by their ids

### Synopsis

Retrieve a list of evaluation forms by their ids

```
gc quality forms evaluations bulk list [flags]
```

### Options

```
  -a, --autopaginate                          Automatically paginate through the results stripping page information
      --filtercondition string                Filter list command output based on a given condition or regular expression
  -h, --help                                  help for list
      --id strings                            A comma-delimited list of valid evaluation form ids. The maximum number of ids allowed in this list is 100 - REQUIRED
      --includeLatestVersionFormName string   Whether to include the name of the form`s most recently published version Valid values: true, false
  -s, --stream                                Paginate and stream data as it is being processed leaving page information intact
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

* [gc quality forms evaluations bulk](gc_quality_forms_evaluations_bulk.html)	 - /api/v2/quality/forms/evaluations/bulk


