package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportrunentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportrunentryDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Reportrunentry
type Reportrunentry struct { 
    


    // Name
    Name string `json:"name"`


    // ReportId
    ReportId string `json:"reportId"`


    // RunTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RunTime time.Time `json:"runTime"`


    // RunStatus
    RunStatus string `json:"runStatus"`


    // ErrorMessage
    ErrorMessage string `json:"errorMessage"`


    // RunDurationMsec
    RunDurationMsec int `json:"runDurationMsec"`


    // ReportUrl
    ReportUrl string `json:"reportUrl"`


    // ReportFormat
    ReportFormat string `json:"reportFormat"`


    // ScheduleUri
    ScheduleUri string `json:"scheduleUri"`


    

}

// String returns a JSON representation of the model
func (o *Reportrunentry) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportrunentry) MarshalJSON() ([]byte, error) {
    type Alias Reportrunentry

    if ReportrunentryMarshalled {
        return []byte("{}"), nil
    }
    ReportrunentryMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        ReportId string `json:"reportId"`
        
        RunTime time.Time `json:"runTime"`
        
        RunStatus string `json:"runStatus"`
        
        ErrorMessage string `json:"errorMessage"`
        
        RunDurationMsec int `json:"runDurationMsec"`
        
        ReportUrl string `json:"reportUrl"`
        
        ReportFormat string `json:"reportFormat"`
        
        ScheduleUri string `json:"scheduleUri"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

