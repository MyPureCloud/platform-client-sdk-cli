## gc scripts published variables get

Get the published variables

### Synopsis

Get the published variables

```
gc scripts published variables get [scriptId] [flags]
```

### Options

```
  -h, --help                       help for get
      --input string               input Valid values: true, false
      --output string              output Valid values: true, false
      --scriptDataVersion string   Advanced usage - controls the data version of the script
      --varType string             type Valid values: string, number, boolean
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

* [gc scripts published variables](gc_scripts_published_variables.html)	 - /api/v2/scripts/published/{scriptId}/variables


