package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Agenticvirtualagentexternala2aservertoolMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Agenticvirtualagentexternala2aservertoolDud struct { 
    

}

// Agenticvirtualagentexternala2aservertool
type Agenticvirtualagentexternala2aservertool struct { 
    // Skills - Agent card skills available for this external A2A server tool.
    Skills []Agenticvirtualagentagentcardskill `json:"skills"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentexternala2aservertool) String() string {
     o.Skills = []Agenticvirtualagentagentcardskill{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentexternala2aservertool) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentexternala2aservertool

    if Agenticvirtualagentexternala2aservertoolMarshalled {
        return []byte("{}"), nil
    }
    Agenticvirtualagentexternala2aservertoolMarshalled = true

    return json.Marshal(&struct {
        
        Skills []Agenticvirtualagentagentcardskill `json:"skills"`
        *Alias
    }{

        
        Skills: []Agenticvirtualagentagentcardskill{{}},
        

        Alias: (*Alias)(u),
    })
}

