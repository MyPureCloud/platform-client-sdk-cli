## gc conversations recordingmetadata

/api/v2/conversations/{conversationId}/recordingmetadata

### Synopsis

/api/v2/conversations/{conversationId}/recordingmetadata

### Options

```
  -h, --help   help for recordingmetadata
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

* [gc conversations](gc_conversations.html)	 - /api/v2/conversations
* [gc conversations recordingmetadata conversationmetadata](gc_conversations_recordingmetadata_conversationmetadata.html)	 - Get recording metadata for a conversation. Does not return playable media nor system annotations. Bookmark annotations will be excluded if either recording:recording:view or recording:annotation:view permission is missing.
* [gc conversations recordingmetadata recordingmetadata](gc_conversations_recordingmetadata_recordingmetadata.html)	 - Get metadata for a specific recording. Does not return playable media.


