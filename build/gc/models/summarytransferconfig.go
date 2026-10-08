package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummarytransferconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummarytransferconfigDud struct { 
    

}

// Summarytransferconfig
type Summarytransferconfig struct { 
    // Enabled - Whether summary transfer data is visible to the receiving agent. Defaults to true when not set.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Summarytransferconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summarytransferconfig) MarshalJSON() ([]byte, error) {
    type Alias Summarytransferconfig

    if SummarytransferconfigMarshalled {
        return []byte("{}"), nil
    }
    SummarytransferconfigMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

