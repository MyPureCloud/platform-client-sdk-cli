package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotcontextvaluesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotcontextvaluesDud struct { 
    

}

// Copilotcontextvalues
type Copilotcontextvalues struct { 
    // Entities - List of copilot context entities grouped by assistant.
    Entities []Copilotcontextentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Copilotcontextvalues) String() string {
     o.Entities = []Copilotcontextentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotcontextvalues) MarshalJSON() ([]byte, error) {
    type Alias Copilotcontextvalues

    if CopilotcontextvaluesMarshalled {
        return []byte("{}"), nil
    }
    CopilotcontextvaluesMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Copilotcontextentity `json:"entities"`
        *Alias
    }{

        
        Entities: []Copilotcontextentity{{}},
        

        Alias: (*Alias)(u),
    })
}

