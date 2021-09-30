package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BaseprogramentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BaseprogramentityDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Baseprogramentity
type Baseprogramentity struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Baseprogramentity) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Baseprogramentity) MarshalJSON() ([]byte, error) {
    type Alias Baseprogramentity

    if BaseprogramentityMarshalled {
        return []byte("{}"), nil
    }
    BaseprogramentityMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

