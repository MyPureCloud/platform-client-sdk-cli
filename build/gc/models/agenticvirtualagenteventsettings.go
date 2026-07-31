package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagenteventsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagenteventsettingsDud struct { 
    


    

}

// Agenticvirtualagenteventsettings
type Agenticvirtualagenteventsettings struct { 
    // VarType - Event type discriminator.
    VarType string `json:"type"`


    // Message - Message the virtual agent should return when this event is handled.
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagenteventsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagenteventsettings) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagenteventsettings

    if AgenticvirtualagenteventsettingsMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagenteventsettingsMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

