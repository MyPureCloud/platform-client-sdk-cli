## gc quality forms surveys list

Get the list of survey forms. If you set expand=publishHistory, then you will be able to get published versions for each corresponding survey form.

### Synopsis

Get the list of survey forms. If you set expand=publishHistory, then you will be able to get published versions for each corresponding survey form.

```
gc quality forms surveys list [flags]
```

### Options

```
  -a, --autopaginate                   Automatically paginate through the results stripping page information
      --expand expand=publishHistory   If expand=publishHistory, then each unpublished evaluation form includes a listing of its published versions Valid values: publishHistory
      --filtercondition string         Filter list command output based on a given condition or regular expression
  -h, --help                           help for list
      --name string                    Name
      --nextPage string                next page token
      --pageNumber int                 The page number requested (default 1)
      --pageSize int                   The total page size requested (default 25)
      --previousPage string            Previous page token
      --sortBy string                  variable name requested to sort by
      --sortOrder string               Order to sort results, either asc or desc
  -s, --stream                         Paginate and stream data as it is being processed leaving page information intact
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

* [gc quality forms surveys](gc_quality_forms_surveys.html)	 - /api/v2/quality/forms/surveys

