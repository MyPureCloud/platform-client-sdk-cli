package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingexportjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingexportjobrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Reportingexportjobrequest
type Reportingexportjobrequest struct { 
    // Name - The user supplied name of the export request
    Name string `json:"name"`


    // TimeZone - The requested timezone of the exported data. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
    TimeZone string `json:"timeZone"`


    // ExportFormat - The requested format of the exported data
    ExportFormat string `json:"exportFormat"`


    // Interval - The time period used to limit the the exported data. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Period - The Period of the request in which to break down the intervals. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
    Period string `json:"period"`


    // ViewType - The type of view export job to be created
    ViewType string `json:"viewType"`


    // Filter - Filters to apply to create the view
    Filter Viewfilter `json:"filter"`


    // Read - Indicates if the request has been marked as read
    Read bool `json:"read"`


    // Locale - The locale use for localization of the exported data, i.e. en-us, es-mx  
    Locale string `json:"locale"`


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

}

// String returns a JSON representation of the model
func (o *Reportingexportjobrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.SelectedColumns = []Selectedcolumns{{}} 
    
    
    
    
    
    
    
     o.RecipientEmails = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingexportjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Reportingexportjobrequest

    if ReportingexportjobrequestMarshalled {
        return []byte("{}"), nil
    }
    ReportingexportjobrequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        TimeZone string `json:"timeZone"`
        
        ExportFormat string `json:"exportFormat"`
        
        Interval string `json:"interval"`
        
        Period string `json:"period"`
        
        ViewType string `json:"viewType"`
        
        Filter Viewfilter `json:"filter"`
        
        Read bool `json:"read"`
        
        Locale string `json:"locale"`
        
        HasFormatDurations bool `json:"hasFormatDurations"`
        
        HasSplitFilters bool `json:"hasSplitFilters"`
        
        ExcludeEmptyRows bool `json:"excludeEmptyRows"`
        
        HasSplitByMedia bool `json:"hasSplitByMedia"`
        
        HasSummaryRow bool `json:"hasSummaryRow"`
        
        CsvDelimiter string `json:"csvDelimiter"`
        
        SelectedColumns []Selectedcolumns `json:"selectedColumns"`
        
        HasCustomParticipantAttributes bool `json:"hasCustomParticipantAttributes"`
        
        RecipientEmails []string `json:"recipientEmails"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        SelectedColumns: []Selectedcolumns{{}},
        

        

        

        

        
        RecipientEmails: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

