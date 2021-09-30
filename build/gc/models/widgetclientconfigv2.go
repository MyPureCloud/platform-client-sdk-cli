package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Widgetclientconfigv2Marshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Widgetclientconfigv2Dud struct { }

// Widgetclientconfigv2
type Widgetclientconfigv2 struct { }

// String returns a JSON representation of the model
func (o *Widgetclientconfigv2) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Widgetclientconfigv2) MarshalJSON() ([]byte, error) {
    type Alias Widgetclientconfigv2

    if Widgetclientconfigv2Marshalled {
        return []byte("{}"), nil
    }
    Widgetclientconfigv2Marshalled = true

    return json.Marshal(&struct { 
        *Alias
    }{
        
        Alias: (*Alias)(u),
    })
}

