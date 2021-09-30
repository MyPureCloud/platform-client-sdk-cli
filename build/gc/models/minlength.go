package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinlengthMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinlengthDud struct { 
    


    

}

// Minlength
type Minlength struct { 
    // Min - A non-negative integer for a text-based schema field denoting the minimum smallest length a string field can contain for a schema instance.
    Min int `json:"min"`


    // Max - A non-negative integer for a text-based schema field denoting the maximum smallest length string the field can contain for a schema instance.
    Max int `json:"max"`

}

// String returns a JSON representation of the model
func (o *Minlength) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Minlength) MarshalJSON() ([]byte, error) {
    type Alias Minlength

    if MinlengthMarshalled {
        return []byte("{}"), nil
    }
    MinlengthMarshalled = true

    return json.Marshal(&struct { 
        Min int `json:"min"`
        
        Max int `json:"max"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

