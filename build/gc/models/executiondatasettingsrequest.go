package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExecutiondatasettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExecutiondatasettingsrequestDud struct { 
    

}

// Executiondatasettingsrequest
type Executiondatasettingsrequest struct { 
    // Enabled - whether or not the setting is enabled.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Executiondatasettingsrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Executiondatasettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Executiondatasettingsrequest

    if ExecutiondatasettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    ExecutiondatasettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

