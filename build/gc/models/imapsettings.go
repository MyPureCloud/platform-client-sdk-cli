package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ImapsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ImapsettingsDud struct { 
    


    Status string `json:"status"`


    ErrorInfo Emailerrorinfo `json:"errorInfo"`

}

// Imapsettings
type Imapsettings struct { 
    // Integration - The IMAP server integration to use for ingesting emails.
    Integration Domainentityref `json:"integration"`


    


    

}

// String returns a JSON representation of the model
func (o *Imapsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Imapsettings) MarshalJSON() ([]byte, error) {
    type Alias Imapsettings

    if ImapsettingsMarshalled {
        return []byte("{}"), nil
    }
    ImapsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Integration Domainentityref `json:"integration"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

