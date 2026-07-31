package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentversionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentversionsettingsDud struct { 
    

}

// Agenticvirtualagentversionsettings - Runtime behavior settings for a virtual agent.
type Agenticvirtualagentversionsettings struct { 
    // ComfortStatement - Comfort statement settings for tool calls.
    ComfortStatement Agenticvirtualagentcomfortstatementsettings `json:"comfortStatement"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentversionsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentversionsettings) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentversionsettings

    if AgenticvirtualagentversionsettingsMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentversionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ComfortStatement Agenticvirtualagentcomfortstatementsettings `json:"comfortStatement"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

