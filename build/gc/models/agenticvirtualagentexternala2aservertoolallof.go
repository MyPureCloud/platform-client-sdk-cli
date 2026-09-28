package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Agenticvirtualagentexternala2aservertoolallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Agenticvirtualagentexternala2aservertoolallofDud struct { 
    

}

// Agenticvirtualagentexternala2aservertoolallof - External A2A server tool.
type Agenticvirtualagentexternala2aservertoolallof struct { 
    // Skills - Agent card skills available for this external A2A server tool.
    Skills []Agenticvirtualagentagentcardskill `json:"skills"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentexternala2aservertoolallof) String() string {
     o.Skills = []Agenticvirtualagentagentcardskill{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentexternala2aservertoolallof) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentexternala2aservertoolallof

    if Agenticvirtualagentexternala2aservertoolallofMarshalled {
        return []byte("{}"), nil
    }
    Agenticvirtualagentexternala2aservertoolallofMarshalled = true

    return json.Marshal(&struct {
        
        Skills []Agenticvirtualagentagentcardskill `json:"skills"`
        *Alias
    }{

        
        Skills: []Agenticvirtualagentagentcardskill{{}},
        

        Alias: (*Alias)(u),
    })
}

