## gc notifications channels

/api/v2/notifications/channels

### Synopsis

/api/v2/notifications/channels

### Options

```
  -h, --help          help for channels
      --noheartbeat   Filters out the heartbeat from the event stream when using the notification service
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

* [gc notifications](gc_notifications.html)	 - /api/v2/notifications
* [gc notifications channels create](gc_notifications_channels_create.html)	 - Create a new channel
* [gc notifications channels head](gc_notifications_channels_head.html)	 - Verify a channel still exists and is valid
* [gc notifications channels list](gc_notifications_channels_list.html)	 - The list of existing channels
* [gc notifications channels listen](gc_notifications_channels_listen.html)	 - Listens to a channel and sends messages to standard out
* [gc notifications channels subscriptions](gc_notifications_channels_subscriptions.html)	 - /api/v2/notifications/channels/{channelId}/subscriptions


