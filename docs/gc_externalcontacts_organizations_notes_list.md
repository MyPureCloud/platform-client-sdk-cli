## gc externalcontacts organizations notes list

List notes for an external organization

### Synopsis

List notes for an external organization

```
gc externalcontacts organizations notes list [externalOrganizationId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand strings           which fields, if any, to expand Valid values: author, externalDataSources, division
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber string        Page number (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000) (default "1")
      --pageSize string          Page size (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000) (default "20")
      --sortOrder string         The Note field to sort by. Any of: [createDate]. Direction: [asc, desc]. e.g. createDate:asc, createDate:desc
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

* [gc externalcontacts organizations notes](gc_externalcontacts_organizations_notes.html)	 - /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes


