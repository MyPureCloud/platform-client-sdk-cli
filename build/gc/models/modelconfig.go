package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ModelconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ModelconfigDud struct { 
    

}

// Modelconfig
type Modelconfig struct { 
    // UseLatestModel - Use the latest model for summarization.
    UseLatestModel bool `json:"useLatestModel"`

}

// String returns a JSON representation of the model
func (o *Modelconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Modelconfig) MarshalJSON() ([]byte, error) {
    type Alias Modelconfig

    if ModelconfigMarshalled {
        return []byte("{}"), nil
    }
    ModelconfigMarshalled = true

    return json.Marshal(&struct {
        
        UseLatestModel bool `json:"useLatestModel"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

