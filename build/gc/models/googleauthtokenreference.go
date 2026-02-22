package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GoogleauthtokenreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GoogleauthtokenreferenceDud struct { 
    

}

// Googleauthtokenreference
type Googleauthtokenreference struct { 
    // Id - ID of the Google OAuth 2 access token. The token cannot be accessed via Genesys API, only referenced by this property. When the token is not referenced by any integration, it is deleted after 24 hours.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Googleauthtokenreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googleauthtokenreference) MarshalJSON() ([]byte, error) {
    type Alias Googleauthtokenreference

    if GoogleauthtokenreferenceMarshalled {
        return []byte("{}"), nil
    }
    GoogleauthtokenreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

