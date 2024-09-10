## gc integrations speech lex bot aliases list

Get a list of aliases for a bot in the customer`s AWS accounts

### Synopsis

Get a list of aliases for a bot in the customer`s AWS accounts

```
gc integrations speech lex bot aliases list [botId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --name string              Filter on alias name
      --pageNumber int           Page number (default 1)
      --pageSize int             Page size (default 25)
      --status string            Filter on alias status Valid values: READY, FAILED, BUILDING, NOT_BUILT
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

* [gc integrations speech lex bot aliases](gc_integrations_speech_lex_bot_aliases.html)	 - /api/v2/integrations/speech/lex/bot/{botId}/aliases


