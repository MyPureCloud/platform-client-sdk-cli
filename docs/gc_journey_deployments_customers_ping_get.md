## gc journey deployments customers ping get

Send a ping.

### Synopsis

Send a ping.

```
gc journey deployments customers ping get [deploymentId] [customerCookieId] [flags]
```

### Options

```
      --appNamespace string               Namespace of the application (e.g. com.genesys.bancodinero). Used for domain filtering in application sessions
      --dl home                           Document Location: 1) Web Page URL if overridden or URL fragment identifier (window.location.hash). OR  2) Application screen name that the ping request was sent from in the app. e.g. home or `help. Pings without this parameter will not return actions.
      --dt string                         Document Title.  A human readable name for the page or screen
  -h, --help                              help for get
      --sessionId string                  UUID of the customer session. Use the same Session Id for all pings, AppEvents and ActionEvents in the session
      --sinceLastBeaconMilliseconds int   How long (milliseconds) since the last app event or beacon was sent. The response may return a pollInternvalMilliseconds to reduce the frequency of pings.
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

* [gc journey deployments customers ping](gc_journey_deployments_customers_ping.html)	 - /api/v2/journey/deployments/{deploymentId}/customers/{customerCookieId}/ping


