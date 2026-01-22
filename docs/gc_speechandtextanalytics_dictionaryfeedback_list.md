## gc speechandtextanalytics dictionaryfeedback list

Get the list of Speech and Text Analytics dictionary feedbacks

### Synopsis

Get the list of Speech and Text Analytics dictionary feedbacks

```
gc speechandtextanalytics dictionaryfeedback list [flags]
```

### Options

```
  -a, --autopaginate                 Automatically paginate through the results stripping page information
      --dialect string               The key for filter the listing by dialect, dialect format is {language}-{country} where language follows ISO 639-1 standard and country follows ISO 3166-1 alpha 2 standard
      --filtercondition string       Filter list command output based on a given condition or regular expression
  -h, --help                         help for list
      --nextPage string              The key for listing the next page
      --pageSize string              The page size for the listing (default "500")
  -s, --stream                       Paginate and stream data as it is being processed leaving page information intact
      --transcriptionEngine string   Filter by transcription engine Valid values: Genesys, GenesysExtended
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

* [gc speechandtextanalytics dictionaryfeedback](gc_speechandtextanalytics_dictionaryfeedback.html)	 - /api/v2/speechandtextanalytics/dictionaryfeedback


