package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulegenerationresultsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulegenerationresultsummaryDud struct { 
    


    


    


    

}

// Schedulegenerationresultsummary
type Schedulegenerationresultsummary struct { 
    // Failed - Whether the schedule generation run failed
    Failed bool `json:"failed"`


    // RunId - The ID of the schedule generation run. Reference this when requesting support
    RunId string `json:"runId"`


    // MessageCount - The number of schedule generation messages for this schedule generation run
    MessageCount int `json:"messageCount"`


    // MessageSeverityCounts - The list of schedule generation message counts by severity for this schedule generation run
    MessageSeverityCounts []Schedulermessageseveritycount `json:"messageSeverityCounts"`

}

// String returns a JSON representation of the model
func (o *Schedulegenerationresultsummary) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.MessageSeverityCounts = []Schedulermessageseveritycount{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulegenerationresultsummary) MarshalJSON() ([]byte, error) {
    type Alias Schedulegenerationresultsummary

    if SchedulegenerationresultsummaryMarshalled {
        return []byte("{}"), nil
    }
    SchedulegenerationresultsummaryMarshalled = true

    return json.Marshal(&struct { 
        Failed bool `json:"failed"`
        
        RunId string `json:"runId"`
        
        MessageCount int `json:"messageCount"`
        
        MessageSeverityCounts []Schedulermessageseveritycount `json:"messageSeverityCounts"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        MessageSeverityCounts: []Schedulermessageseveritycount{{}},
        

        
        Alias: (*Alias)(u),
    })
}

