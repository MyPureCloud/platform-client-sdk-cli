package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentversionpublishMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentversionpublishDud struct { 
    

}

// Agenticvirtualagentversionpublish
type Agenticvirtualagentversionpublish struct { 
    // Status - The status of the virtual agent version to update as part of this publish job.
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentversionpublish) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentversionpublish) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentversionpublish

    if AgenticvirtualagentversionpublishMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentversionpublishMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

