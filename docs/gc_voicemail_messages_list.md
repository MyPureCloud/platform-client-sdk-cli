## gc voicemail messages list

List voicemail messages

### Synopsis

List voicemail messages

```
gc voicemail messages list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand strings           If the caller is a known user, which fields, if any, to expand Valid values: callerUser.routingStatus, callerUser.primaryPresence, callerUser.conversationSummary, callerUser.outOfOffice, callerUser.geolocation, conversations, transcription
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --ids string               An optional comma separated list of VoicemailMessage ids
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

* [gc voicemail messages](gc_voicemail_messages.html)	 - /api/v2/voicemail/messages


