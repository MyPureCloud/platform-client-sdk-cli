## gc learning modules assignments list

Get all learning modules of an organization including assignments for a specific user

### Synopsis

Get all learning modules of an organization including assignments for a specific user

```
gc learning modules assignments list [flags]
```

### Options

```
      --assignmentStates strings   Specifies the assignment states to return. Valid values: NotAssigned, Assigned, InProgress, Completed, InvalidSchedule
  -a, --autopaginate               Automatically paginate through the results stripping page information
      --expand strings             Fields to expand in response(case insensitive) Valid values: coverArt
      --filtercondition string     Filter list command output based on a given condition or regular expression
  -h, --help                       help for list
      --overdue string             Specifies if only modules with overdue/not overdue (overdue is True or False) assignments are returned. If overdue is Any or omitted, both are returned and can including modules that are unassigned. Valid values: True, False, Any
      --pageNumber string          Page number (default "1")
      --pageSize string            Page size (default "25")
      --searchTerm string          Search Term (searches by name and description)
  -s, --stream                     Paginate and stream data as it is being processed leaving page information intact
      --userIds strings            The IDs of the users to include - REQUIRED
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

* [gc learning modules assignments](gc_learning_modules_assignments.html)	 - /api/v2/learning/modules/assignments


