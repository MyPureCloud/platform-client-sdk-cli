## gc telephony siptraces get

Fetch SIP metadata

### Synopsis

Fetch SIP metadata

```
gc telephony siptraces get [flags]
```

### Options

```
      --callId string           unique identification of the placed call
      --conversationId string   Unique identification of the conversation
      --dateEnd string          End date of the search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z - REQUIRED
      --dateStart string        Start date of the search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z - REQUIRED
      --fromUser string         user who placed the call
  -h, --help                    help for get
      --toUser string           User to who the call was placed
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

* [gc telephony siptraces](gc_telephony_siptraces.html)	 - /api/v2/telephony/siptraces


