## gc greetings media get

Get media playback URI for this greeting

### Synopsis

Get media playback URI for this greeting

```
gc greetings media get [greetingId] [flags]
```

### Options

```
      --formatId string   The desired media format. Valid values: WAV, WEBM, WAV_ULAW, OGG_VORBIS, OGG_OPUS, MP3, NONE
  -h, --help              help for get
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

* [gc greetings media](gc_greetings_media.html)	 - /api/v2/greetings/{greetingId}/media

