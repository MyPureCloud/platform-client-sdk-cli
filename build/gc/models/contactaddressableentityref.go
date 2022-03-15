package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactaddressableentityrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactaddressableentityrefDud struct { 
    


    

}

// Contactaddressableentityref
type Contactaddressableentityref struct { 
    // Id
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Contactaddressableentityref) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactaddressableentityref) MarshalJSON() ([]byte, error) {
    type Alias Contactaddressableentityref

    if ContactaddressableentityrefMarshalled {
        return []byte("{}"), nil
    }
    ContactaddressableentityrefMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

