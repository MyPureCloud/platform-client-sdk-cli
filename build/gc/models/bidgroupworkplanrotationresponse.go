package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BidgroupworkplanrotationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BidgroupworkplanrotationresponseDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Bidgroupworkplanrotationresponse
type Bidgroupworkplanrotationresponse struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // ManagementUnit - The management unit to which the work plan rotation belongs
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // AgentCount - The count of agents that can be assigned to this work plan rotation
    AgentCount int `json:"agentCount"`


    

}

// String returns a JSON representation of the model
func (o *Bidgroupworkplanrotationresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bidgroupworkplanrotationresponse) MarshalJSON() ([]byte, error) {
    type Alias Bidgroupworkplanrotationresponse

    if BidgroupworkplanrotationresponseMarshalled {
        return []byte("{}"), nil
    }
    BidgroupworkplanrotationresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        AgentCount int `json:"agentCount"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

