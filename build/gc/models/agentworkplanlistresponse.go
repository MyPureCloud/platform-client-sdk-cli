package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentworkplanlistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentworkplanlistresponseDud struct { 
    


    

}

// Agentworkplanlistresponse
type Agentworkplanlistresponse struct { 
    // Entities
    Entities []Agentworkplan `json:"entities"`


    // ManagementUnit - The management unit of the work plans
    ManagementUnit Managementunitreference `json:"managementUnit"`

}

// String returns a JSON representation of the model
func (o *Agentworkplanlistresponse) String() string {
     o.Entities = []Agentworkplan{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentworkplanlistresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentworkplanlistresponse

    if AgentworkplanlistresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentworkplanlistresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Agentworkplan `json:"entities"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        *Alias
    }{

        
        Entities: []Agentworkplan{{}},
        


        

        Alias: (*Alias)(u),
    })
}

