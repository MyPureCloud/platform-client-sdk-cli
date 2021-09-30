package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookappcredentialsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookappcredentialsDud struct { 
    Id string `json:"id"`

}

// Facebookappcredentials
type Facebookappcredentials struct { 
    

}

// String returns a JSON representation of the model
func (o *Facebookappcredentials) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookappcredentials) MarshalJSON() ([]byte, error) {
    type Alias Facebookappcredentials

    if FacebookappcredentialsMarshalled {
        return []byte("{}"), nil
    }
    FacebookappcredentialsMarshalled = true

    return json.Marshal(&struct { 
        
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

