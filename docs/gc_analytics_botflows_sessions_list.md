## gc analytics botflows sessions list

Get Bot Flow Sessions.

### Synopsis

Get Bot Flow Sessions.

```
gc analytics botflows sessions list [botFlowId] [flags]
```

### Options

```
      --after string                                                        The cursor that points to the ID of the last item in the list of entities that has been returned.
  -a, --autopaginate                                                        Automatically paginate through the results stripping page information
      --botResultCategories string                                          Optional case-insensitive comma separated list of Bot Result Categories to filter sessions by. Valid values: Unknown, UserExit, BotExit, Error, RecognitionFailure, UserDisconnect, BotDisconnect, SessionExpired, Transfer
      --endLanguage string                                                  Optional case-insensitive language code to filter sessions by the language the sessions ended in.
      --filtercondition string                                              Filter list command output based on a given condition or regular expression
  -h, --help                                                                help for list
      --interval 2022-11-22T09:11:11.111+08:00/2022-11-30T07:17:44.586-07   Date range filter based on the date the individual resources were completed. UTC is the default if no TZ is supplied, however alternate timezones can be used e.g: 2022-11-22T09:11:11.111+08:00/2022-11-30T07:17:44.586-07. . Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
      --pageSize string                                                     Max number of entities to return. Maximum of 250
  -s, --stream                                                              Paginate and stream data as it is being processed leaving page information intact
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

* [gc analytics botflows sessions](gc_analytics_botflows_sessions.html)	 - /api/v2/analytics/botflows/{botFlowId}/sessions


