package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MaxlengthMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MaxlengthDud struct { 
    


    

}

// Maxlength
type Maxlength struct { 
    // Min - A non-negative integer for a text-based schema field denoting the minimum largest length string the field can contain for a schema instance.
    Min int `json:"min"`


    // Max - A non-negative integer for a text-based schema field denoting the maximum largest string the field can contain for a schema instance.
    Max int `json:"max"`

}

// String returns a JSON representation of the model
func (o *Maxlength) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Maxlength) MarshalJSON() ([]byte, error) {
    type Alias Maxlength

    if MaxlengthMarshalled {
        return []byte("{}"), nil
    }
    MaxlengthMarshalled = true

    return json.Marshal(&struct { 
        Min int `json:"min"`
        
        Max int `json:"max"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

