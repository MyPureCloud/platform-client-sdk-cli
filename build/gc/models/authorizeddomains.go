package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthorizeddomainsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthorizeddomainsDud struct { 
    

}

// Authorizeddomains - Domains authorized for email processing.
type Authorizeddomains struct { 
    // Outbound - List of authorized domains for outbound send.
    Outbound []string `json:"outbound"`

}

// String returns a JSON representation of the model
func (o *Authorizeddomains) String() string {
     o.Outbound = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authorizeddomains) MarshalJSON() ([]byte, error) {
    type Alias Authorizeddomains

    if AuthorizeddomainsMarshalled {
        return []byte("{}"), nil
    }
    AuthorizeddomainsMarshalled = true

    return json.Marshal(&struct {
        
        Outbound []string `json:"outbound"`
        *Alias
    }{

        
        Outbound: []string{""},
        

        Alias: (*Alias)(u),
    })
}

