## gc notifications channels subscriptions subscribe

Add a list of subscriptions to the existing list of subscriptions

### Synopsis

Add a list of subscriptions to the existing list of subscriptions

```
gc notifications channels subscriptions subscribe [channelId] [flags]
```

### Options

```
  -d, --directory string      Directory path with files containing request bodies
  -f, --file string           File name containing the JSON body
  -h, --help                  help for subscribe
      --ignoreErrors string   Optionally prevent throwing of errors for failed permissions checks. Valid values: true, false
  -b, --printrequestbody      Print the request body format of the API.
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

* [gc notifications channels subscriptions](gc_notifications_channels_subscriptions.html)	 - /api/v2/notifications/channels/{channelId}/subscriptions


