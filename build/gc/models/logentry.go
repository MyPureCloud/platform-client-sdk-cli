package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LogentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LogentryDud struct { 
    


    


    

}

// Logentry
type Logentry struct { 
    // Level - Level of log entry
    Level string `json:"level"`


    // Message - Log message
    Message string `json:"message"`


    // Timestamp - Timestamp of log entry
    Timestamp int `json:"timestamp"`

}

// String returns a JSON representation of the model
func (o *Logentry) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Logentry) MarshalJSON() ([]byte, error) {
    type Alias Logentry

    if LogentryMarshalled {
        return []byte("{}"), nil
    }
    LogentryMarshalled = true

    return json.Marshal(&struct {
        
        Level string `json:"level"`
        
        Message string `json:"message"`
        
        Timestamp int `json:"timestamp"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

