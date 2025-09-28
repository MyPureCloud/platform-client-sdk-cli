## gc speechandtextanalytics topics list

Get the list of Speech and Text Analytics topics

### Synopsis

Get the list of Speech and Text Analytics topics

```
gc speechandtextanalytics topics list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --dialects strings         Comma separated dialect strings to filter by. Maximum of 15 dialects allowed. Valid values: en-US, es-US, en-AU, en-GB, en-ZA, es-ES, en-IN, fr-FR, fr-CA, it-IT, de-DE, pt-BR, pl-PL, pt-PT, nl-NL, ko-KR
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --ids strings              Comma separated Topic IDs to filter by. Cannot be used with other filters. Maximum of 50 IDs allowed.
      --name string              Case insensitive partial name to filter by
      --nextPage string          The key for listing the next page
      --pageSize string          The page size for the listing. The max that will be returned is 500. (default "20")
      --sortBy string            Sort results by. Defaults to name Valid values: name, matchingType
      --sortOrder string         Sort order. Defaults to asc Valid values: asc, desc
      --state string             Topic state. Defaults to latest Valid values: latest, published
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

* [gc speechandtextanalytics topics](gc_speechandtextanalytics_topics.html)	 - /api/v2/speechandtextanalytics/topics


