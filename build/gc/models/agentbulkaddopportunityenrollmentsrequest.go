package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentbulkaddopportunityenrollmentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentbulkaddopportunityenrollmentsrequestDud struct { 
    

}

// Agentbulkaddopportunityenrollmentsrequest
type Agentbulkaddopportunityenrollmentsrequest struct { 
    // Enrollments - The list of the enrollments to add
    Enrollments []Agentaddopportunityenrollment `json:"enrollments"`

}

// String returns a JSON representation of the model
func (o *Agentbulkaddopportunityenrollmentsrequest) String() string {
     o.Enrollments = []Agentaddopportunityenrollment{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentbulkaddopportunityenrollmentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Agentbulkaddopportunityenrollmentsrequest

    if AgentbulkaddopportunityenrollmentsrequestMarshalled {
        return []byte("{}"), nil
    }
    AgentbulkaddopportunityenrollmentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Enrollments []Agentaddopportunityenrollment `json:"enrollments"`
        *Alias
    }{

        
        Enrollments: []Agentaddopportunityenrollment{{}},
        

        Alias: (*Alias)(u),
    })
}

