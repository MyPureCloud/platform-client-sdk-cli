## gc authorization divisionspermitted paged list

Returns which divisions the specified user has the given permission in.

### Synopsis

Returns which divisions the specified user has the given permission in.

```
gc authorization divisionspermitted paged list [subjectId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "25")
      --permission string        The permission string, including the object to access, e.g. routing:queue:view - REQUIRED
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

* [gc authorization divisionspermitted paged](gc_authorization_divisionspermitted_paged.html)	 - /api/v2/authorization/divisionspermitted/paged


