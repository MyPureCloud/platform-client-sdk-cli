package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CloneduserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CloneduserDud struct { 
    Id string `json:"id"`


    


    Trustor Domainentityref `json:"trustor"`


    SelfUri string `json:"selfUri"`

}

// Cloneduser - Represents a cloned user in a trustor organization.
type Cloneduser struct { 
    


    // Name
    Name string `json:"name"`


    


    

}

// String returns a JSON representation of the model
func (o *Cloneduser) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cloneduser) MarshalJSON() ([]byte, error) {
    type Alias Cloneduser

    if CloneduserMarshalled {
        return []byte("{}"), nil
    }
    CloneduserMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

