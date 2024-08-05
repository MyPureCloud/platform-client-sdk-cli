package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BidgroupworkplanresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BidgroupworkplanresponseDud struct { 
    Id string `json:"id"`


    


    


    SuggestedAgentCount int `json:"suggestedAgentCount"`


    AgentCountRange Agentcountrange `json:"agentCountRange"`


    SelfUri string `json:"selfUri"`

}

// Bidgroupworkplanresponse
type Bidgroupworkplanresponse struct { 
    


    // ManagementUnit - The management unit to which this work plan belongs.  Nullable in some routes
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // OverrideAgentCount - The modified agent count for this work plan
    OverrideAgentCount int `json:"overrideAgentCount"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Bidgroupworkplanresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bidgroupworkplanresponse) MarshalJSON() ([]byte, error) {
    type Alias Bidgroupworkplanresponse

    if BidgroupworkplanresponseMarshalled {
        return []byte("{}"), nil
    }
    BidgroupworkplanresponseMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        OverrideAgentCount int `json:"overrideAgentCount"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

