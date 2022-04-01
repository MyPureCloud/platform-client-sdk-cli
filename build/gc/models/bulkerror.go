package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkerrorDud struct { 
    


    

}

// Bulkerror
type Bulkerror struct { 
    // Message - Error message of the bulk operation result.
    Message string `json:"message"`


    // Code - Error code of the bulk operation result.
    Code string `json:"code"`

}

// String returns a JSON representation of the model
func (o *Bulkerror) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkerror) MarshalJSON() ([]byte, error) {
    type Alias Bulkerror

    if BulkerrorMarshalled {
        return []byte("{}"), nil
    }
    BulkerrorMarshalled = true

    return json.Marshal(&struct { 
        Message string `json:"message"`
        
        Code string `json:"code"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

