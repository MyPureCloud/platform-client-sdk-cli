package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeekschedulegenerationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeekschedulegenerationresultDud struct { 
    


    


    


    

}

// Weekschedulegenerationresult
type Weekschedulegenerationresult struct { 
    // Failed - Whether the schedule generation failed
    Failed bool `json:"failed"`


    // RunId - ID of the schedule run
    RunId string `json:"runId"`


    // AgentWarnings - Warning messages from the schedule run. This will be available only when requesting information for a single week schedule
    AgentWarnings []Schedulegenerationwarning `json:"agentWarnings"`


    // AgentWarningCount - Count of warning messages from the schedule run. This will be available only when requesting multiple week schedules
    AgentWarningCount int `json:"agentWarningCount"`

}

// String returns a JSON representation of the model
func (o *Weekschedulegenerationresult) String() string {
    
    
    
    
    
    
    
    
    
    
     o.AgentWarnings = []Schedulegenerationwarning{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weekschedulegenerationresult) MarshalJSON() ([]byte, error) {
    type Alias Weekschedulegenerationresult

    if WeekschedulegenerationresultMarshalled {
        return []byte("{}"), nil
    }
    WeekschedulegenerationresultMarshalled = true

    return json.Marshal(&struct { 
        Failed bool `json:"failed"`
        
        RunId string `json:"runId"`
        
        AgentWarnings []Schedulegenerationwarning `json:"agentWarnings"`
        
        AgentWarningCount int `json:"agentWarningCount"`
        
        *Alias
    }{
        

        

        

        

        

        
        AgentWarnings: []Schedulegenerationwarning{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

