## gc languageunderstanding miners intents list

Retrieve a list of mined intents.

### Synopsis

Retrieve a list of mined intents.

```
gc languageunderstanding miners intents list [minerId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand string            Option to fetch utterances. Valid values: phrases, utterances
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
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

* [gc languageunderstanding miners intents](gc_languageunderstanding_miners_intents.html)	 - /api/v2/languageunderstanding/miners/{minerId}/intents

