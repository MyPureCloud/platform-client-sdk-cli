package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentownedmappingpreviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentownedmappingpreviewDud struct { 
    AgentOwnedColumn string `json:"agentOwnedColumn"`


    Email string `json:"email"`


    UserId string `json:"userId"`


    Exists bool `json:"exists"`


    IsQueueMember bool `json:"isQueueMember"`


    RecordCount int `json:"recordCount"`

}

// Agentownedmappingpreview
type Agentownedmappingpreview struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Agentownedmappingpreview) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentownedmappingpreview) MarshalJSON() ([]byte, error) {
    type Alias Agentownedmappingpreview

    if AgentownedmappingpreviewMarshalled {
        return []byte("{}"), nil
    }
    AgentownedmappingpreviewMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

