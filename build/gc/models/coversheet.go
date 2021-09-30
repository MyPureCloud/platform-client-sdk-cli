package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoversheetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoversheetDud struct { 
    


    

}

// Coversheet
type Coversheet struct { 
    // Notes - Text to be added to the coversheet
    Notes string `json:"notes"`


    // Locale - Locale, e.g. = en-US
    Locale string `json:"locale"`

}

// String returns a JSON representation of the model
func (o *Coversheet) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coversheet) MarshalJSON() ([]byte, error) {
    type Alias Coversheet

    if CoversheetMarshalled {
        return []byte("{}"), nil
    }
    CoversheetMarshalled = true

    return json.Marshal(&struct { 
        Notes string `json:"notes"`
        
        Locale string `json:"locale"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

