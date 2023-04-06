package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportedcontentprofileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportedcontentprofileDud struct { 
    

}

// Supportedcontentprofile
type Supportedcontentprofile struct { 
    // Id - The supported content profile ID
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Supportedcontentprofile) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportedcontentprofile) MarshalJSON() ([]byte, error) {
    type Alias Supportedcontentprofile

    if SupportedcontentprofileMarshalled {
        return []byte("{}"), nil
    }
    SupportedcontentprofileMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

