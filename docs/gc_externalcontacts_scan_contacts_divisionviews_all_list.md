## gc externalcontacts scan contacts divisionviews all list

Scan for external contacts using paging

### Synopsis

Scan for external contacts using paging

```
gc externalcontacts scan contacts divisionviews all list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --cursor string            Indicates where to resume query results (not required for first page), each page returns a new cursor with a 24h TTL
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --limit string             The number of contacts per page; must be between 10 and 200, default is 100
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

* [gc externalcontacts scan contacts divisionviews all](gc_externalcontacts_scan_contacts_divisionviews_all.html)	 - /api/v2/externalcontacts/scan/contacts/divisionviews/all


