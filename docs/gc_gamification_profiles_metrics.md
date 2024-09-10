## gc gamification profiles metrics

/api/v2/gamification/profiles/{sourceProfileId}/metrics /api/v2/gamification/profiles/{profileId}/metrics

### Synopsis

/api/v2/gamification/profiles/{sourceProfileId}/metrics /api/v2/gamification/profiles/{profileId}/metrics

### Options

```
  -h, --help   help for metrics
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

* [gc gamification profiles](gc_gamification_profiles.html)	 - /api/v2/gamification/profiles
* [gc gamification profiles metrics create](gc_gamification_profiles_metrics_create.html)	 - Creates a gamified metric with a given metric definition and metric objective under in a performance profile
* [gc gamification profiles metrics get](gc_gamification_profiles_metrics_get.html)	 - Performance profile gamified metric by id
* [gc gamification profiles metrics link](gc_gamification_profiles_metrics_link.html)	 - /api/v2/gamification/profiles/{sourceProfileId}/metrics/{sourceMetricId}/link
* [gc gamification profiles metrics list](gc_gamification_profiles_metrics_list.html)	 - All gamified metrics for a given performance profile
* [gc gamification profiles metrics objectivedetails](gc_gamification_profiles_metrics_objectivedetails.html)	 - /api/v2/gamification/profiles/{profileId}/metrics/objectivedetails
* [gc gamification profiles metrics update](gc_gamification_profiles_metrics_update.html)	 - Updates a metric in performance profile


