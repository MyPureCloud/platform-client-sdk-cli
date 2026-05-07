package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotcontextentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotcontextentityDud struct { 
    


    

}

// Copilotcontextentity
type Copilotcontextentity struct { 
    // Assistant - The assistant associated with this context.
    Assistant Addressableentityref `json:"assistant"`


    // ContextValues - List of copilot context values for this assistant.
    ContextValues []Copilotcontextvalueitem `json:"contextValues"`

}

// String returns a JSON representation of the model
func (o *Copilotcontextentity) String() string {
    
     o.ContextValues = []Copilotcontextvalueitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotcontextentity) MarshalJSON() ([]byte, error) {
    type Alias Copilotcontextentity

    if CopilotcontextentityMarshalled {
        return []byte("{}"), nil
    }
    CopilotcontextentityMarshalled = true

    return json.Marshal(&struct {
        
        Assistant Addressableentityref `json:"assistant"`
        
        ContextValues []Copilotcontextvalueitem `json:"contextValues"`
        *Alias
    }{

        


        
        ContextValues: []Copilotcontextvalueitem{{}},
        

        Alias: (*Alias)(u),
    })
}

