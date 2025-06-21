package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuideversionpublishMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuideversionpublishDud struct { 
    

}

// Guideversionpublish
type Guideversionpublish struct { 
    // State - The desired state of the guide version being published.
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Guideversionpublish) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guideversionpublish) MarshalJSON() ([]byte, error) {
    type Alias Guideversionpublish

    if GuideversionpublishMarshalled {
        return []byte("{}"), nil
    }
    GuideversionpublishMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

