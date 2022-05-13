package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainroleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainroleDud struct { 
    


    

}

// Domainrole
type Domainrole struct { 
    // Id - The ID of the role
    Id string `json:"id"`


    // Name - The name of the role
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Domainrole) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainrole) MarshalJSON() ([]byte, error) {
    type Alias Domainrole

    if DomainroleMarshalled {
        return []byte("{}"), nil
    }
    DomainroleMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

