## gc quality calibrations

/api/v2/quality/calibrations

### Synopsis

/api/v2/quality/calibrations

### Options

```
  -h, --help   help for calibrations
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

* [gc quality](gc_quality.html)	 - /api/v2/quality
* [gc quality calibrations create](gc_quality_calibrations_create.html)	 - Create a calibration
* [gc quality calibrations delete](gc_quality_calibrations_delete.html)	 - Delete a calibration by id.
* [gc quality calibrations get](gc_quality_calibrations_get.html)	 - Get a calibration by id.  Requires either calibrator id or conversation id
* [gc quality calibrations list](gc_quality_calibrations_list.html)	 - Get the list of calibrations
* [gc quality calibrations update](gc_quality_calibrations_update.html)	 - Update a calibration to the specified calibration via PUT.  Editable fields include: evaluators, expertEvaluator, and scoringIndex


