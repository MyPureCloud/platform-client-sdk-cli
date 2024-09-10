## gc speechandtextanalytics programs delete

Delete a Speech and Text Analytics program by id

### Synopsis

Delete a Speech and Text Analytics program by id

```
gc speechandtextanalytics programs delete [programId] [flags]
```

### Options

```
      --forceDelete string   Indicates whether the program is forced to be deleted or not. Required when the program to delete is the default program. Valid values: true, false Valid values: true, false
  -h, --help                 help for delete
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


