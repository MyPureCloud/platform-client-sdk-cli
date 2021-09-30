package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LogcaptureuserconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LogcaptureuserconfigurationDud struct { 
    Id string `json:"id"`


    DateExpired time.Time `json:"dateExpired"`


    SelfUri string `json:"selfUri"`

}

// Logcaptureuserconfiguration
type Logcaptureuserconfiguration struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Logcaptureuserconfiguration) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Logcaptureuserconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Logcaptureuserconfiguration

    if LogcaptureuserconfigurationMarshalled {
        return []byte("{}"), nil
    }
    LogcaptureuserconfigurationMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

