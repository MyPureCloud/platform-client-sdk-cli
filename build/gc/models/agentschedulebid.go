package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentschedulebidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentschedulebidDud struct { 
    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentschedulebid
type Agentschedulebid struct { 
    // Id - The ID of the schedule bid
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // BidWindowStartDate - The bid start date when agents can start participating in schedule bidding relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowStartDate time.Time `json:"bidWindowStartDate"`


    // BidWindowEndDate - The bid end date relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BidWindowEndDate time.Time `json:"bidWindowEndDate"`


    // EffectiveDate - The date when schedule sets would be effective for schedule generation relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EffectiveDate time.Time `json:"effectiveDate"`


    // Status - The state of the bid
    Status string `json:"status"`


    // BidType - The type of the bid
    BidType string `json:"bidType"`


    

}

// String returns a JSON representation of the model
func (o *Agentschedulebid) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentschedulebid) MarshalJSON() ([]byte, error) {
    type Alias Agentschedulebid

    if AgentschedulebidMarshalled {
        return []byte("{}"), nil
    }
    AgentschedulebidMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        BidWindowStartDate time.Time `json:"bidWindowStartDate"`
        
        BidWindowEndDate time.Time `json:"bidWindowEndDate"`
        
        EffectiveDate time.Time `json:"effectiveDate"`
        
        Status string `json:"status"`
        
        BidType string `json:"bidType"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

