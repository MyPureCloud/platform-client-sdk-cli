package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LinkconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LinkconfigurationDud struct { 
    

}

// Linkconfiguration
type Linkconfiguration struct { 
    // UriTemplate - Should be a valid URL (including the http/https protocol, port, and path [if any]). Leading and trailing whitespace stripped. Max 400 characters.
    UriTemplate string `json:"uriTemplate"`

}

// String returns a JSON representation of the model
func (o *Linkconfiguration) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Linkconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Linkconfiguration

    if LinkconfigurationMarshalled {
        return []byte("{}"), nil
    }
    LinkconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        UriTemplate string `json:"uriTemplate"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

