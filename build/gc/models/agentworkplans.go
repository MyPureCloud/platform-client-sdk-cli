package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentworkplansMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentworkplansDud struct { 
    


    

}

// Agentworkplans
type Agentworkplans struct { 
    // User - The user (agent) for whom the work plans were requested
    User Userreference `json:"user"`


    // WorkPlanLookupKeysPerWeek - The list of weekly work plan lookup keys. Keys to be searched under workPlanLookup property at top level of result
    WorkPlanLookupKeysPerWeek []int `json:"workPlanLookupKeysPerWeek"`

}

// String returns a JSON representation of the model
func (o *Agentworkplans) String() string {
    
     o.WorkPlanLookupKeysPerWeek = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentworkplans) MarshalJSON() ([]byte, error) {
    type Alias Agentworkplans

    if AgentworkplansMarshalled {
        return []byte("{}"), nil
    }
    AgentworkplansMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        WorkPlanLookupKeysPerWeek []int `json:"workPlanLookupKeysPerWeek"`
        *Alias
    }{

        


        
        WorkPlanLookupKeysPerWeek: []int{0},
        

        Alias: (*Alias)(u),
    })
}

