package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DigitlengthMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DigitlengthDud struct { 
    


    

}

// Digitlength
type Digitlength struct { 
    // Start
    Start string `json:"start"`


    // End
    End string `json:"end"`

}

// String returns a JSON representation of the model
func (o *Digitlength) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Digitlength) MarshalJSON() ([]byte, error) {
    type Alias Digitlength

    if DigitlengthMarshalled {
        return []byte("{}"), nil
    }
    DigitlengthMarshalled = true

    return json.Marshal(&struct { 
        Start string `json:"start"`
        
        End string `json:"end"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

