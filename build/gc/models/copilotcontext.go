package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotcontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotcontextDud struct { 
    

}

// Copilotcontext
type Copilotcontext struct { 
    // Values - Copilot context values.
    Values []Copilotcontextvalue `json:"values"`

}

// String returns a JSON representation of the model
func (o *Copilotcontext) String() string {
     o.Values = []Copilotcontextvalue{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotcontext) MarshalJSON() ([]byte, error) {
    type Alias Copilotcontext

    if CopilotcontextMarshalled {
        return []byte("{}"), nil
    }
    CopilotcontextMarshalled = true

    return json.Marshal(&struct {
        
        Values []Copilotcontextvalue `json:"values"`
        *Alias
    }{

        
        Values: []Copilotcontextvalue{{}},
        

        Alias: (*Alias)(u),
    })
}

