## gc casemanagement cases externalcontacts list

Get a list of Cases for an External Contact.

### Synopsis

Get a list of Cases for an External Contact.

```
gc casemanagement cases externalcontacts list [externalContactId] [flags]
```

### Options

```
      --after string             Cursor pointing to the end of the previously returned page of Cases.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --divisionIds string       Filter by divisions.
      --expands strings          Fields to expand. Valid values: caseplan
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageSize string          Number of Cases to return (maximum 200).
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

* [gc casemanagement cases externalcontacts](gc_casemanagement_cases_externalcontacts.html)	 - /api/v2/casemanagement/cases/externalcontacts


