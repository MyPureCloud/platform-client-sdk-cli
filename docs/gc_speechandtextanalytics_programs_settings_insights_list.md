## gc speechandtextanalytics programs settings insights list

Get the list of program AI Insights settings for the organization

### Synopsis

Get the list of program AI Insights settings for the organization

```
gc speechandtextanalytics programs settings insights list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber string        The page number for the listing (default "1")
      --pageSize string          The page size for the listing. The max that will be returned is 100. (default "100")
      --programIds strings       Comma separated Program IDs to filter by. Maximum of 50 IDs allowed.
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

* [gc speechandtextanalytics programs settings insights](gc_speechandtextanalytics_programs_settings_insights.html)	 - /api/v2/speechandtextanalytics/programs/{programId}/settings/insights /api/v2/speechandtextanalytics/programs/settings/insights


