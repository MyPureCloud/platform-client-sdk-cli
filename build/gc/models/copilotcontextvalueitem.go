package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotcontextvalueitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotcontextvalueitemDud struct { 
    


    

}

// Copilotcontextvalueitem
type Copilotcontextvalueitem struct { 
    // Name - Context name.
    Name string `json:"name"`


    // Value - Context value.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Copilotcontextvalueitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotcontextvalueitem) MarshalJSON() ([]byte, error) {
    type Alias Copilotcontextvalueitem

    if CopilotcontextvalueitemMarshalled {
        return []byte("{}"), nil
    }
    CopilotcontextvalueitemMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

