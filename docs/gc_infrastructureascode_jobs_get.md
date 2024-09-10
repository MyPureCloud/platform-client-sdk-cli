## gc infrastructureascode jobs get

Get job history

### Synopsis

Get job history

```
gc infrastructureascode jobs get [flags]
```

### Options

```
      --acceleratorId string   Find only jobs associated with this accelerator
  -h, --help                   help for get
      --includeErrors string   Include error messages Valid values: true, false
      --maxResults int         Number of jobs to show (default 1)
      --sortBy string          Sort by Valid values: id, dateSubmitted, submittedBy, acceleratorId, status
      --sortOrder string       Sort order Valid values: asc, desc
      --status string          Find only jobs in this state Valid values: Created, Queued, Running, Complete, Failed, Incomplete
      --submittedBy string     Find only jobs submitted by this user
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

* [gc infrastructureascode jobs](gc_infrastructureascode_jobs.html)	 - /api/v2/infrastructureascode/jobs


