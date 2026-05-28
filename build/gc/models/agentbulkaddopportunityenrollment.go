package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentbulkaddopportunityenrollmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentbulkaddopportunityenrollmentDud struct { 
    Id string `json:"id"`


    

}

// Agentbulkaddopportunityenrollment
type Agentbulkaddopportunityenrollment struct { 
    


    // OpportunityId - The ID of the opportunity in which the agent was enrolled
    OpportunityId string `json:"opportunityId"`

}

// String returns a JSON representation of the model
func (o *Agentbulkaddopportunityenrollment) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentbulkaddopportunityenrollment) MarshalJSON() ([]byte, error) {
    type Alias Agentbulkaddopportunityenrollment

    if AgentbulkaddopportunityenrollmentMarshalled {
        return []byte("{}"), nil
    }
    AgentbulkaddopportunityenrollmentMarshalled = true

    return json.Marshal(&struct {
        
        OpportunityId string `json:"opportunityId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

