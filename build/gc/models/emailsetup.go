package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailsetupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailsetupDud struct { 
    

}

// Emailsetup
type Emailsetup struct { 
    // RootDomain - The root PureCloud domain that all sub-domains are created from.
    RootDomain string `json:"rootDomain"`

}

// String returns a JSON representation of the model
func (o *Emailsetup) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailsetup) MarshalJSON() ([]byte, error) {
    type Alias Emailsetup

    if EmailsetupMarshalled {
        return []byte("{}"), nil
    }
    EmailsetupMarshalled = true

    return json.Marshal(&struct { 
        RootDomain string `json:"rootDomain"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

