package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentqueryadherenceadjustmentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentqueryadherenceadjustmentsrequestDud struct { 
    


    


    


    

}

// Agentqueryadherenceadjustmentsrequest
type Agentqueryadherenceadjustmentsrequest struct { 
    // StartDate - The start timestamp of the range to query in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - The end timestamp of the range to query in ISO-8601 format
    EndDate time.Time `json:"endDate"`


    // ReasonCodeIds - A filter for the reason codes to include. Leave empty or omit entirely for all reason codes
    ReasonCodeIds []string `json:"reasonCodeIds"`


    // Statuses - A filter for which adherence adjustment statuses to include. Leave empty or omit entirely for all statuses
    Statuses []string `json:"statuses"`

}

// String returns a JSON representation of the model
func (o *Agentqueryadherenceadjustmentsrequest) String() string {
    
    
     o.ReasonCodeIds = []string{""} 
     o.Statuses = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentqueryadherenceadjustmentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Agentqueryadherenceadjustmentsrequest

    if AgentqueryadherenceadjustmentsrequestMarshalled {
        return []byte("{}"), nil
    }
    AgentqueryadherenceadjustmentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        ReasonCodeIds []string `json:"reasonCodeIds"`
        
        Statuses []string `json:"statuses"`
        *Alias
    }{

        


        


        
        ReasonCodeIds: []string{""},
        


        
        Statuses: []string{""},
        

        Alias: (*Alias)(u),
    })
}

