package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportscheduleDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Reportschedule
type Reportschedule struct { 
    


    // Name
    Name string `json:"name"`


    // QuartzCronExpression - Quartz Cron Expression
    QuartzCronExpression string `json:"quartzCronExpression"`


    // NextFireTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    NextFireTime time.Time `json:"nextFireTime"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // Description
    Description string `json:"description"`


    // TimeZone
    TimeZone string `json:"timeZone"`


    // TimePeriod
    TimePeriod string `json:"timePeriod"`


    // Interval - Interval. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // ReportFormat
    ReportFormat string `json:"reportFormat"`


    // Locale
    Locale string `json:"locale"`


    // Enabled
    Enabled bool `json:"enabled"`


    // ReportId - Report ID
    ReportId string `json:"reportId"`


    // Parameters
    Parameters map[string]interface{} `json:"parameters"`


    // LastRun
    LastRun Reportrunentry `json:"lastRun"`


    

}

// String returns a JSON representation of the model
func (o *Reportschedule) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Parameters = map[string]interface{}{"": Interface{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportschedule) MarshalJSON() ([]byte, error) {
    type Alias Reportschedule

    if ReportscheduleMarshalled {
        return []byte("{}"), nil
    }
    ReportscheduleMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        QuartzCronExpression string `json:"quartzCronExpression"`
        
        NextFireTime time.Time `json:"nextFireTime"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        Description string `json:"description"`
        
        TimeZone string `json:"timeZone"`
        
        TimePeriod string `json:"timePeriod"`
        
        Interval string `json:"interval"`
        
        ReportFormat string `json:"reportFormat"`
        
        Locale string `json:"locale"`
        
        Enabled bool `json:"enabled"`
        
        ReportId string `json:"reportId"`
        
        Parameters map[string]interface{} `json:"parameters"`
        
        LastRun Reportrunentry `json:"lastRun"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Parameters: map[string]interface{}{"": Interface{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

