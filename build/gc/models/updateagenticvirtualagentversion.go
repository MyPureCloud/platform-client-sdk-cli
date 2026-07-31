package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateagenticvirtualagentversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateagenticvirtualagentversionDud struct { 
    

}

// Updateagenticvirtualagentversion
type Updateagenticvirtualagentversion struct { 
    // Definition - The definition of the virtual agent.
    Definition Agenticvirtualagentversiondefinition `json:"definition"`

}

// String returns a JSON representation of the model
func (o *Updateagenticvirtualagentversion) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateagenticvirtualagentversion) MarshalJSON() ([]byte, error) {
    type Alias Updateagenticvirtualagentversion

    if UpdateagenticvirtualagentversionMarshalled {
        return []byte("{}"), nil
    }
    UpdateagenticvirtualagentversionMarshalled = true

    return json.Marshal(&struct {
        
        Definition Agenticvirtualagentversiondefinition `json:"definition"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

