## gc conversations recordings list

Get all of a Conversation`s Recordings.

### Synopsis

Get all of a Conversation`s Recordings.

```
gc conversations recordings list [conversationId] [flags]
```

### Options

```
      --formatId string                                     The desired media format. Valid values:WAV,WEBM,WAV_ULAW,OGG_VORBIS,OGG_OPUS,MP3,NONE. Valid values: WAV, WEBM, WAV_ULAW, OGG_VORBIS, OGG_OPUS, MP3, NONE
  -h, --help                                                help for list
      --includePauseAnnotationsForScreenRecordings string   Include applicable Secure Pause annotations from all audio recordings to all screen recordings Valid values: true, false
      --locale string                                       The locale used for redacting sensitive information in requested files, as an ISO 639-1 code
      --maxWaitMs string                                    The maximum number of milliseconds to wait for the recording to be ready. Must be a positive value. (default "5000")
      --mediaFormats strings                                All acceptable media formats. Overrides formatId. Valid values:WAV,WEBM,WAV_ULAW,OGG_VORBIS,OGG_OPUS,MP3.
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

* [gc conversations recordings](gc_conversations_recordings.html)	 - /api/v2/conversations/{conversationId}/recordings


