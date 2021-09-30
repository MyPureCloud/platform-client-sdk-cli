package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainentityrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainentityrefDud struct { 
    


    


    

}

// Domainentityref
type Domainentityref struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Domainentityref) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainentityref) MarshalJSON() ([]byte, error) {
    type Alias Domainentityref

    if DomainentityrefMarshalled {
        return []byte("{}"), nil
    }
    DomainentityrefMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

