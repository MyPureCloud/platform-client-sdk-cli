package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntakeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntakeDud struct { 
    

}

// Intake
type Intake struct { 
    // Data - The intake data containing key-value pairs.
    Data map[string]interface{} `json:"data"`

}

// String returns a JSON representation of the model
func (o *Intake) String() string {
     o.Data = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intake) MarshalJSON() ([]byte, error) {
    type Alias Intake

    if IntakeMarshalled {
        return []byte("{}"), nil
    }
    IntakeMarshalled = true

    return json.Marshal(&struct {
        
        Data map[string]interface{} `json:"data"`
        *Alias
    }{

        
        Data: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

