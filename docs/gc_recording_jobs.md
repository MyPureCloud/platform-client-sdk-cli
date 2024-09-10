## gc recording jobs

/api/v2/recording/jobs

### Synopsis

/api/v2/recording/jobs

### Options

```
  -h, --help   help for jobs
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

* [gc recording](gc_recording.html)	 - /api/v2/recording
* [gc recording jobs create](gc_recording_jobs_create.html)	 - Create a recording bulk job.
* [gc recording jobs delete](gc_recording_jobs_delete.html)	 - Delete the recording bulk job
* [gc recording jobs failedrecordings](gc_recording_jobs_failedrecordings.html)	 - /api/v2/recording/jobs/{jobId}/failedrecordings
* [gc recording jobs get](gc_recording_jobs_get.html)	 - Get the status of the job associated with the job id.
* [gc recording jobs list](gc_recording_jobs_list.html)	 - Get the status of all jobs within the user`s organization
* [gc recording jobs update](gc_recording_jobs_update.html)	 - Execute the recording bulk job.


