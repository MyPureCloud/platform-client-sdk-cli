package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentaddopportunityenrollmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentaddopportunityenrollmentDud struct { 
    

}

// Agentaddopportunityenrollment
type Agentaddopportunityenrollment struct { 
    // OpportunityId - The ID of the opportunity in which to enroll the agent
    OpportunityId string `json:"opportunityId"`

}

// String returns a JSON representation of the model
func (o *Agentaddopportunityenrollment) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentaddopportunityenrollment) MarshalJSON() ([]byte, error) {
    type Alias Agentaddopportunityenrollment

    if AgentaddopportunityenrollmentMarshalled {
        return []byte("{}"), nil
    }
    AgentaddopportunityenrollmentMarshalled = true

    return json.Marshal(&struct {
        
        OpportunityId string `json:"opportunityId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

