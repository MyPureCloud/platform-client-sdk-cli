package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentworkplanbidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentworkplanbidDud struct { 
    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentworkplanbid - Work plan bid reference
type Agentworkplanbid struct { 
    // Id - The ID of the work plan bid
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // BidWindowStartDate - The date when agents can start participating in work plan bidding. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowStartDate time.Time `json:"bidWindowStartDate"`


    // BidWindowEndDate - The inclusive end date of a bid window. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowEndDate time.Time `json:"bidWindowEndDate"`


    // EffectiveDate - The date when agents will be assigned to the new work plan. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EffectiveDate time.Time `json:"effectiveDate"`


    // Status - The state of the bid
    Status string `json:"status"`


    // WorkPlanFieldsVisibleToAgents - The work plan fields visible to agents whenever work plan preferences are made
    WorkPlanFieldsVisibleToAgents []string `json:"workPlanFieldsVisibleToAgents"`


    

}

// String returns a JSON representation of the model
func (o *Agentworkplanbid) String() string {
    
    
    
    
    
    
     o.WorkPlanFieldsVisibleToAgents = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentworkplanbid) MarshalJSON() ([]byte, error) {
    type Alias Agentworkplanbid

    if AgentworkplanbidMarshalled {
        return []byte("{}"), nil
    }
    AgentworkplanbidMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        BidWindowStartDate time.Time `json:"bidWindowStartDate"`
        
        BidWindowEndDate time.Time `json:"bidWindowEndDate"`
        
        EffectiveDate time.Time `json:"effectiveDate"`
        
        Status string `json:"status"`
        
        WorkPlanFieldsVisibleToAgents []string `json:"workPlanFieldsVisibleToAgents"`
        *Alias
    }{

        


        


        


        


        


        


        
        WorkPlanFieldsVisibleToAgents: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

