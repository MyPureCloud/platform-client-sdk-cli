## gc casemanagement caseplans list

Get a list of Caseplans.

### Synopsis

Get a list of Caseplans.

```
gc casemanagement caseplans list [flags]
```

### Options

```
      --after string              The cursor that points to the end of the set of caseplans that has been returned.
  -a, --autopaginate              Automatically paginate through the results stripping page information
      --customerIntentId string   Filter by Customer Intent.
      --divisionIds string        Filter by Divisions.
      --filtercondition string    Filter list command output based on a given condition or regular expression
  -h, --help                      help for list
      --pageSize string           Number of caseplans to return. Maximum of 200.
  -s, --stream                    Paginate and stream data as it is being processed leaving page information intact
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

* [gc casemanagement caseplans](gc_casemanagement_caseplans.html)	 - /api/v2/casemanagement/caseplans


