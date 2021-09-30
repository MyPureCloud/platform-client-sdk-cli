package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallrecordMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallrecordDud struct { 
    LastAttempt time.Time `json:"lastAttempt"`


    LastResult string `json:"lastResult"`

}

// Callrecord
type Callrecord struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Callrecord) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callrecord) MarshalJSON() ([]byte, error) {
    type Alias Callrecord

    if CallrecordMarshalled {
        return []byte("{}"), nil
    }
    CallrecordMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

