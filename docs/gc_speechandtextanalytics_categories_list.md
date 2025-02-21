## gc speechandtextanalytics categories list

Get the list of Speech and Text Analytics categories

### Synopsis

Get the list of Speech and Text Analytics categories

```
gc speechandtextanalytics categories list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --ids strings              Comma separated Category IDs to filter by. Cannot be used with other filters. Maximum of 25 IDs allowed.
      --name string              The category name filter applied to the listing
      --pageNumber string        The page number for the listing (default "1")
      --pageSize string          The page size for the listing. The max that will be returned is 25. (default "25")
      --sortBy string            The field to sort by for the listing Valid values: name, description
      --sortOrder string         The sort order for the listing Valid values: asc, desc
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

* [gc speechandtextanalytics categories](gc_speechandtextanalytics_categories.html)	 - /api/v2/speechandtextanalytics/categories


