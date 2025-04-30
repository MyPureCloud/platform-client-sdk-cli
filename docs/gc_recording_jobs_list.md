## gc recording jobs list

Get the status of all jobs within the user`s organization

### Synopsis

Get the status of all jobs within the user`s organization

```
gc recording jobs list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --cursor string            Indicates where to resume query results (not required for first page)
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --includeTotal string      If false, cursor will be used to locate the page instead of pageNumber. It is recommended to set it to false for improved performance. Valid values: true, false
      --jobType string           Job Type (Can be left empty for both) Valid values: ARCHIVE, DELETE, EXPORT
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "25")
      --showOnlyMyJobs string    Show only my jobs Valid values: true, false
      --sortBy string            Sort by Valid values: userId, dateCreated
      --state string             Filter by state Valid values: FULFILLED, PENDING, READY, PROCESSING, CANCELLED, FAILED
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

* [gc recording jobs](gc_recording_jobs.html)	 - /api/v2/recording/jobs


