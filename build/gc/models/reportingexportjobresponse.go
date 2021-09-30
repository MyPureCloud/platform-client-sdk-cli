package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingexportjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingexportjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Reportingexportjobresponse
type Reportingexportjobresponse struct { 
    


    // Name
    Name string `json:"name"`


    // RunId - The unique run id of the export schedule execute
    RunId string `json:"runId"`


    // Status - The current status of the export request
    Status string `json:"status"`


    // TimeZone - The requested timezone of the exported data. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
    TimeZone string `json:"timeZone"`


    // ExportFormat - The requested format of the exported data
    ExportFormat string `json:"exportFormat"`


    // Interval - The time period used to limit the the exported data. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // DownloadUrl - The url to download the request if it's status is completed
    DownloadUrl string `json:"downloadUrl"`


    // ViewType - The type of view export job to be created
    ViewType string `json:"viewType"`


    // ExportErrorMessagesType - The error message in case the export request failed
    ExportErrorMessagesType string `json:"exportErrorMessagesType"`


    // Period - The Period of the request in which to break down the intervals. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
    Period string `json:"period"`


    // Filter - Filters to apply to create the view
    Filter Viewfilter `json:"filter"`


    // Read - Indicates if the request has been marked as read
    Read bool `json:"read"`


    // CreatedDateTime - The created date/time of the request. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDateTime time.Time `json:"createdDateTime"`


    // ModifiedDateTime - The last modified date/time of the request. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDateTime time.Time `json:"modifiedDateTime"`


    // Locale - The locale use for localization of the exported data, i.e. en-us, es-mx  
    Locale string `json:"locale"`


    // PercentageComplete - The percentage of the job that has completed processing
    PercentageComplete float64 `json:"percentageComplete"`


    // HasFormatDurations - Indicates if durations are formatted in hh:mm:ss format instead of ms
    HasFormatDurations bool `json:"hasFormatDurations"`


    // HasSplitFilters - Indicates if filters will be split in aggregate detail exports
    HasSplitFilters bool `json:"hasSplitFilters"`


    // ExcludeEmptyRows - Excludes empty rows from the exports
    ExcludeEmptyRows bool `json:"excludeEmptyRows"`


    // HasSplitByMedia - Indicates if media type will be split in aggregate detail exports
    HasSplitByMedia bool `json:"hasSplitByMedia"`


    // HasSummaryRow - Indicates if summary row needs to be present in exports
    HasSummaryRow bool `json:"hasSummaryRow"`


    // CsvDelimiter - The user supplied csv delimiter string value either of type 'comma' or 'semicolon' permitted for the export request
    CsvDelimiter string `json:"csvDelimiter"`


    // SelectedColumns - The list of ordered selected columns from the export view by the user
    SelectedColumns []Selectedcolumns `json:"selectedColumns"`


    // HasCustomParticipantAttributes - Indicates if custom participant attributes will be exported
    HasCustomParticipantAttributes bool `json:"hasCustomParticipantAttributes"`


    // RecipientEmails - The list of email recipients for the exports
    RecipientEmails []string `json:"recipientEmails"`


    // EmailStatuses - The status of individual email addresses as a map
    EmailStatuses map[string]string `json:"emailStatuses"`


    // EmailErrorDescription - The optional error message in case the export fail to email
    EmailErrorDescription string `json:"emailErrorDescription"`


    // Enabled
    Enabled bool `json:"enabled"`


    

}

// String returns a JSON representation of the model
func (o *Reportingexportjobresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.SelectedColumns = []Selectedcolumns{{}} 
    
    
    
    
    
    
    
     o.RecipientEmails = []string{""} 
    
    
    
     o.EmailStatuses = map[string]string{"": ""} 
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingexportjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Reportingexportjobresponse

    if ReportingexportjobresponseMarshalled {
        return []byte("{}"), nil
    }
    ReportingexportjobresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        RunId string `json:"runId"`
        
        Status string `json:"status"`
        
        TimeZone string `json:"timeZone"`
        
        ExportFormat string `json:"exportFormat"`
        
        Interval string `json:"interval"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        ViewType string `json:"viewType"`
        
        ExportErrorMessagesType string `json:"exportErrorMessagesType"`
        
        Period string `json:"period"`
        
        Filter Viewfilter `json:"filter"`
        
        Read bool `json:"read"`
        
        CreatedDateTime time.Time `json:"createdDateTime"`
        
        ModifiedDateTime time.Time `json:"modifiedDateTime"`
        
        Locale string `json:"locale"`
        
        PercentageComplete float64 `json:"percentageComplete"`
        
        HasFormatDurations bool `json:"hasFormatDurations"`
        
        HasSplitFilters bool `json:"hasSplitFilters"`
        
        ExcludeEmptyRows bool `json:"excludeEmptyRows"`
        
        HasSplitByMedia bool `json:"hasSplitByMedia"`
        
        HasSummaryRow bool `json:"hasSummaryRow"`
        
        CsvDelimiter string `json:"csvDelimiter"`
        
        SelectedColumns []Selectedcolumns `json:"selectedColumns"`
        
        HasCustomParticipantAttributes bool `json:"hasCustomParticipantAttributes"`
        
        RecipientEmails []string `json:"recipientEmails"`
        
        EmailStatuses map[string]string `json:"emailStatuses"`
        
        EmailErrorDescription string `json:"emailErrorDescription"`
        
        Enabled bool `json:"enabled"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        SelectedColumns: []Selectedcolumns{{}},
        

        

        

        

        
        RecipientEmails: []string{""},
        

        

        
        EmailStatuses: map[string]string{"": ""},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

