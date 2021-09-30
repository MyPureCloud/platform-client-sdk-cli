package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PublishformMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PublishformDud struct { 
    


    

}

// Publishform
type Publishform struct { 
    // Published - Is this form published
    Published bool `json:"published"`


    // Id - Unique Id for this version of this form
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Publishform) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Publishform) MarshalJSON() ([]byte, error) {
    type Alias Publishform

    if PublishformMarshalled {
        return []byte("{}"), nil
    }
    PublishformMarshalled = true

    return json.Marshal(&struct { 
        Published bool `json:"published"`
        
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

