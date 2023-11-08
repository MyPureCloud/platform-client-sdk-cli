package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DisablesiteconnectionsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DisablesiteconnectionsrequestDud struct { 
    

}

// Disablesiteconnectionsrequest
type Disablesiteconnectionsrequest struct { 
    // Enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Disablesiteconnectionsrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Disablesiteconnectionsrequest) MarshalJSON() ([]byte, error) {
    type Alias Disablesiteconnectionsrequest

    if DisablesiteconnectionsrequestMarshalled {
        return []byte("{}"), nil
    }
    DisablesiteconnectionsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

