package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkaddopportunitiesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkaddopportunitiesrequestDud struct { 
    


    

}

// Bulkaddopportunitiesrequest
type Bulkaddopportunitiesrequest struct { 
    // Opportunities - The opportunities to add
    Opportunities []Addopportunitybody `json:"opportunities"`


    // AgentIds - The IDs of the agents who are invited to the opportunities being added
    AgentIds []string `json:"agentIds"`

}

// String returns a JSON representation of the model
func (o *Bulkaddopportunitiesrequest) String() string {
     o.Opportunities = []Addopportunitybody{{}} 
     o.AgentIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkaddopportunitiesrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkaddopportunitiesrequest

    if BulkaddopportunitiesrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkaddopportunitiesrequestMarshalled = true

    return json.Marshal(&struct {
        
        Opportunities []Addopportunitybody `json:"opportunities"`
        
        AgentIds []string `json:"agentIds"`
        *Alias
    }{

        
        Opportunities: []Addopportunitybody{{}},
        


        
        AgentIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

