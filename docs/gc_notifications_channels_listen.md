## gc notifications channels listen

Listens to a channel and sends messages to standard out

### Synopsis

Listens to a channel and sends messages to standard out

```
gc notifications channels listen [channelId] [flags]
```

### Options

```
  -h, --help   help for listen
```

### Options inherited from parent commands

```
      --accesstoken string    accessToken override
      --clientid string       clientId override
      --clientsecret string   clientSecret override
      --environment string    environment override. E.g. mypurecloud.com.au or ap-southeast-2
  -i, --indicateprogress      Trace progress indicators to stderr
      --inputformat string    Data input format. Supported formats: YAML, JSON
      --noheartbeat           Filters out the heartbeat from the event stream when using the notification service
      --outputformat string   Data output format. Supported formats: YAML, JSON
  -p, --profile string        Name of the profile to use for configuring the cli (default "DEFAULT")
      --transform string      Provide a Go template file for transforming output data
      --transformstr string   Provide a Go template string for transforming output data
```

### SEE ALSO

* [gc notifications channels](gc_notifications_channels.html)	 - /api/v2/notifications/channels


