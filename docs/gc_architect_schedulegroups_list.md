## gc architect schedulegroups list

Get a list of schedule groups.

### Synopsis

Get a list of schedule groups.

```
gc architect schedulegroups list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --divisionId strings       List of divisionIds on which to filter.
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --name string              Name of the Schedule Group to filter by.
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "25")
      --scheduleIds string       A comma-delimited list of Schedule IDs to filter by.
      --sortBy string            Sort by
      --sortOrder string         Sort order
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

* [gc architect schedulegroups](gc_architect_schedulegroups.html)	 - /api/v2/architect/schedulegroups


