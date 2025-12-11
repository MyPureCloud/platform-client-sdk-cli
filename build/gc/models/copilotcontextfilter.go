package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotcontextfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotcontextfilterDud struct { 
    


    

}

// Copilotcontextfilter
type Copilotcontextfilter struct { 
    // Operator - Operator to apply for multiple context values, default: MatchAll.
    Operator string `json:"operator"`


    // Values - Context names to use for filtering.
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Copilotcontextfilter) String() string {
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotcontextfilter) MarshalJSON() ([]byte, error) {
    type Alias Copilotcontextfilter

    if CopilotcontextfilterMarshalled {
        return []byte("{}"), nil
    }
    CopilotcontextfilterMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

