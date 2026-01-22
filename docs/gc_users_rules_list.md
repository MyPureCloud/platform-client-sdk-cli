## gc users rules list

Get the list of users rules

### Synopsis

Get the list of users rules

```
gc users rules list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --enabled string           Whether to list enabled or disabled rules Valid values: true, false
      --expand strings           Fields to expand in response Valid values: criteria
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber string        Page number (default "1")
      --pageSize string          Number of results per page (default "25")
      --searchTerm string        a search term for finding a rule by name
      --sortOrder string         sort rules by name, ascending, descending Valid values: ascending, descending
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --types strings            The types of the rule - REQUIRED Valid values: Learning, ActivityPlan
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

* [gc users rules](gc_users_rules.html)	 - /api/v2/users/rules/{ruleId}/dependents /api/v2/users/rules


