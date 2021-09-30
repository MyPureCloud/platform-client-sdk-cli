package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DeletableuserreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DeletableuserreferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Deletableuserreference - User reference with delete flag to remove the user from an associated entity
type Deletableuserreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Delete - If marked true, the user will be removed an associated entity
    Delete bool `json:"delete"`


    

}

// String returns a JSON representation of the model
func (o *Deletableuserreference) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Deletableuserreference) MarshalJSON() ([]byte, error) {
    type Alias Deletableuserreference

    if DeletableuserreferenceMarshalled {
        return []byte("{}"), nil
    }
    DeletableuserreferenceMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Delete bool `json:"delete"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

