package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DeleteretentionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DeleteretentionDud struct { 
    

}

// Deleteretention
type Deleteretention struct { 
    // Days
    Days int `json:"days"`

}

// String returns a JSON representation of the model
func (o *Deleteretention) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Deleteretention) MarshalJSON() ([]byte, error) {
    type Alias Deleteretention

    if DeleteretentionMarshalled {
        return []byte("{}"), nil
    }
    DeleteretentionMarshalled = true

    return json.Marshal(&struct { 
        Days int `json:"days"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

