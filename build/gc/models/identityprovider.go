package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IdentityproviderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IdentityproviderDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Identityprovider
type Identityprovider struct { 
    


    // Name
    Name string `json:"name"`


    // Disabled
    Disabled bool `json:"disabled"`


    

}

// String returns a JSON representation of the model
func (o *Identityprovider) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Identityprovider) MarshalJSON() ([]byte, error) {
    type Alias Identityprovider

    if IdentityproviderMarshalled {
        return []byte("{}"), nil
    }
    IdentityproviderMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Disabled bool `json:"disabled"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

