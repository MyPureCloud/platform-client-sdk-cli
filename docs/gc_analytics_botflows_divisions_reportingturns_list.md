## gc analytics botflows divisions reportingturns list

Get Reporting Turns (division aware).

### Synopsis

Get Reporting Turns (division aware).

```
gc analytics botflows divisions reportingturns list [botFlowId] [flags]
```

### Options

```
      --actionId string                                                     Optional action ID to get the reporting turns associated to a particular flow action
      --after string                                                        The cursor that points to the ID of the last item in the list of entities that has been returned.
      --askActionResults string                                             Optional case-insensitive comma separated list of ask action results to filter the reporting turns. Valid values: AgentRequestedByUser, ConfirmationRequired, DisambiguationRequired, Error, ExpressionError, NoInputCollection, NoInputConfirmation, NoInputDisambiguation, NoMatchCollection, NoMatchConfirmation, NoMatchDisambiguation, SuccessCollection, SuccessConfirmationNo, SuccessConfirmationYes, SuccessDisambiguation, SuccessDisambiguationNone
  -a, --autopaginate                                                        Automatically paginate through the results stripping page information
      --filtercondition string                                              Filter list command output based on a given condition or regular expression
  -h, --help                                                                help for list
      --interval 2022-11-22T09:11:11.111+08:00/2022-11-30T07:17:44.586-07   Date range filter based on the date the individual resources were completed. UTC is the default if no TZ is supplied, however alternate timezones can be used e.g: 2022-11-22T09:11:11.111+08:00/2022-11-30T07:17:44.586-07. . Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
      --language string                                                     Optional language code to get the reporting turns for a particular language
      --pageSize string                                                     Max number of entities to return. Maximum of 250
      --sessionId string                                                    Optional session ID to get the reporting turns for a particular session. Specifying a session ID alongside an action ID or a language or any ask action results is not allowed.
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

* [gc analytics botflows divisions reportingturns](gc_analytics_botflows_divisions_reportingturns.html)	 - /api/v2/analytics/botflows/{botFlowId}/divisions/reportingturns


