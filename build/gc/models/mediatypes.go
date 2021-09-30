package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediatypesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediatypesDud struct { 
    

}

// Mediatypes - Media types
type Mediatypes struct { 
    // Allow - Specify allowed media types for inbound and outbound messages. If this field is empty, all inbound and outbound media will be blocked.
    Allow Mediatypeaccess `json:"allow"`

}

// String returns a JSON representation of the model
func (o *Mediatypes) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediatypes) MarshalJSON() ([]byte, error) {
    type Alias Mediatypes

    if MediatypesMarshalled {
        return []byte("{}"), nil
    }
    MediatypesMarshalled = true

    return json.Marshal(&struct { 
        Allow Mediatypeaccess `json:"allow"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

