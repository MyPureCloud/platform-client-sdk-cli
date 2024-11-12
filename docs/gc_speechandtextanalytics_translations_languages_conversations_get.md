## gc speechandtextanalytics translations languages conversations get

Translate a single interaction recording (or an email conversation)

### Synopsis

Translate a single interaction recording (or an email conversation)

```
gc speechandtextanalytics translations languages conversations get [languageId] [conversationId] [flags]
```

### Options

```
      --communicationId string   Communication id associated with the conversation. Please provide a valid communicationId when requesting non-email interactions.
  -h, --help                     help for get
      --recordingId string       Recording id associated with the communication. Please provide a valid recordingId when requesting voice interactions.
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

* [gc speechandtextanalytics translations languages conversations](gc_speechandtextanalytics_translations_languages_conversations.html)	 - /api/v2/speechandtextanalytics/translations/languages/{languageId}/conversations


