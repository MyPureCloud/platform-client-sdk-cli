## gc signals anomalydetection queues analysis get

Get forecasted bounds and analysis for the specified queue

### Synopsis

Get forecasted bounds and analysis for the specified queue

```
gc signals anomalydetection queues analysis get [queueId] [flags]
```

### Options

```
      --endTime string       The end of the time range for the analysis, in ISO-8601 format. If not specified, defaults to 48 hours after the most recent UTC midnight
  -h, --help                 help for get
      --mediaTypes strings   The media types by which to filter the analysis. When not specified, returns analysis for Email, Message, and Voice Valid values: Email, Message, Voice
      --metrics strings      The metrics by which to filter the analysis. When not specified, returns all metrics Valid values: nOffered, tTalkCompleteAverage, oAnsweredPercent, oAbandonPercent, oServiceLevel, tAnsweredMax, tAnsweredAverage, tAcwAverage, tHeldCompleteAverage, tHandleMax, tHandleAverage, nHandled, nTransferred
      --startTime string     The start of the time range for the analysis, in ISO-8601 format. If not specified, defaults to 48 hours before the most recent UTC midnight
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

* [gc signals anomalydetection queues analysis](gc_signals_anomalydetection_queues_analysis.html)	 - /api/v2/signals/anomalydetection/queues/{queueId}/analysis


