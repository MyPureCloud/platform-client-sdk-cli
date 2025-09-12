package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainsDud struct { 
    

}

// Domains - The domain list settings.
type Domains struct { 
    // AuthorizedDomains - The authorized domains settings for email processing.
    AuthorizedDomains Authorizeddomains `json:"authorizedDomains"`

}

// String returns a JSON representation of the model
func (o *Domains) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domains) MarshalJSON() ([]byte, error) {
    type Alias Domains

    if DomainsMarshalled {
        return []byte("{}"), nil
    }
    DomainsMarshalled = true

    return json.Marshal(&struct {
        
        AuthorizedDomains Authorizeddomains `json:"authorizedDomains"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

