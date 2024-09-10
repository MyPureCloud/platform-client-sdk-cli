## gc conversations recordings get

Gets a specific recording.

### Synopsis

Gets a specific recording.

```
gc conversations recordings get [conversationId] [recordingId] [flags]
```

### Options

```
      --chatFormatId string      The desired media format when downloading a chat recording. Valid values:ZIP,NONE  Valid values: ZIP, NONE
      --download string          requesting a download format of the recording. Valid values:true,false Valid values: true, false Valid values: true, false
      --emailFormatId string     The desired media format when downloading an email recording. Valid values:EML,NONE Valid values: EML, NONE
      --fileName string          the name of the downloaded fileName
      --formatId string          The desired media format. Valid values:WAV,WEBM,WAV_ULAW,OGG_VORBIS,OGG_OPUS,MP3,NONE Valid values: WAV, WEBM, WAV_ULAW, OGG_VORBIS, OGG_OPUS, MP3, NONE
  -h, --help                     help for get
      --locale string            The locale for the requested file when downloading, as an ISO 639-1 code
      --mediaFormats strings     All acceptable media formats. Overrides formatId. Valid values:WAV,WEBM,WAV_ULAW,OGG_VORBIS,OGG_OPUS,MP3
      --messageFormatId string   The desired media format when downloading a message recording. Valid values:ZIP,NONE Valid values: ZIP, NONE
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


