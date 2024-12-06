## gc socialmedia topics dataingestionrules open delete

Delete a open data ingestion rule.

### Synopsis

Delete a open data ingestion rule.

```
gc socialmedia topics dataingestionrules open delete [topicId] [openId] [flags]
```

### Options

```
      --hardDelete string   Determines whether a open data ingestion rule should be soft-deleted (have it`s state set to deleted) or hard-deleted (permanently removed). Set to false (soft-delete) by default. Valid values: true, false
  -h, --help                help for delete
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

* [gc socialmedia topics dataingestionrules open](gc_socialmedia_topics_dataingestionrules_open.html)	 - /api/v2/socialmedia/topics/{topicId}/dataingestionrules/open


