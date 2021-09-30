package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainorganizationproductMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainorganizationproductDud struct { 
    

}

// Domainorganizationproduct
type Domainorganizationproduct struct { 
    // Id
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Domainorganizationproduct) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainorganizationproduct) MarshalJSON() ([]byte, error) {
    type Alias Domainorganizationproduct

    if DomainorganizationproductMarshalled {
        return []byte("{}"), nil
    }
    DomainorganizationproductMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

