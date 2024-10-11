package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemruleactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemruleactionDud struct { 
    

}

// Workitemruleaction
type Workitemruleaction struct { 
    // VarType
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Workitemruleaction) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemruleaction) MarshalJSON() ([]byte, error) {
    type Alias Workitemruleaction

    if WorkitemruleactionMarshalled {
        return []byte("{}"), nil
    }
    WorkitemruleactionMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

