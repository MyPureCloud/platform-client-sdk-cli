package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesettingdynamicfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesettingdynamicfilterDud struct { 
    

}

// Knowledgesettingdynamicfilter
type Knowledgesettingdynamicfilter struct { 
    // Context - Filter based on copilot context values.
    Context Copilotcontextfilter `json:"context"`

}

// String returns a JSON representation of the model
func (o *Knowledgesettingdynamicfilter) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesettingdynamicfilter) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesettingdynamicfilter

    if KnowledgesettingdynamicfilterMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesettingdynamicfilterMarshalled = true

    return json.Marshal(&struct {
        
        Context Copilotcontextfilter `json:"context"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

