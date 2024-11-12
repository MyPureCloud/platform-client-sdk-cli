## gc textbots bots search get

Find bots using the currently configured friendly name or ID.

### Synopsis

Find bots using the currently configured friendly name or ID.

```
gc textbots bots search get [flags]
```

### Options

```
      --botId strings     Bot IDs
      --botName string    Bot name
      --botType strings   Bot types Valid values: GenesysBotConnector, GenesysDialogEngine, AmazonLex, GoogleDialogFlowES, GoogleDialogFlowCX, NuanceDlg, GenesysBotFlow, GenesysDigitalBotFlow, GenesysVoiceSurveyFlow
  -h, --help              help for get
      --pageSize string   The maximum results to return (default "25")
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

* [gc textbots bots search](gc_textbots_bots_search.html)	 - /api/v2/textbots/bots/search


