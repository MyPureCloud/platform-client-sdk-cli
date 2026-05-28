package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentqueryopportunitiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentqueryopportunitiesDud struct { 
    


    

}

// Agentqueryopportunities
type Agentqueryopportunities struct { 
    // NextStartDate - The start date to use for the next query to retrieve additional results in ISO-8601 format. Null if there are no more results
    NextStartDate time.Time `json:"nextStartDate"`


    // BusinessUnits - The opportunities for the agent grouped by business unit
    BusinessUnits []Agentbusinessunitopportunities `json:"businessUnits"`

}

// String returns a JSON representation of the model
func (o *Agentqueryopportunities) String() string {
    
     o.BusinessUnits = []Agentbusinessunitopportunities{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentqueryopportunities) MarshalJSON() ([]byte, error) {
    type Alias Agentqueryopportunities

    if AgentqueryopportunitiesMarshalled {
        return []byte("{}"), nil
    }
    AgentqueryopportunitiesMarshalled = true

    return json.Marshal(&struct {
        
        NextStartDate time.Time `json:"nextStartDate"`
        
        BusinessUnits []Agentbusinessunitopportunities `json:"businessUnits"`
        *Alias
    }{

        


        
        BusinessUnits: []Agentbusinessunitopportunities{{}},
        

        Alias: (*Alias)(u),
    })
}

