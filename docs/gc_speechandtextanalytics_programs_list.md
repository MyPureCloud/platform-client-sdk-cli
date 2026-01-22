## gc speechandtextanalytics programs list

Get the list of Speech and Text Analytics programs

### Synopsis

Get the list of Speech and Text Analytics programs

```
gc speechandtextanalytics programs list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --name string              Case insensitive partial name to filter by
      --nextPage string          The key for listing the next page
      --pageSize string          The page size for the listing (default "20")
      --sortBy string            Sort results by. Defaults to name Valid values: name
      --sortOrder string         Sort order. Defaults to asc Valid values: asc, desc
      --state string             Program state. Defaults to Latest Valid values: Latest, Published
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

* [gc speechandtextanalytics programs](gc_speechandtextanalytics_programs.html)	 - /api/v2/speechandtextanalytics/programs


