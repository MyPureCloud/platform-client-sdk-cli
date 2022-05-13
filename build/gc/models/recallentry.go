package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecallentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecallentryDud struct { 
    


    

}

// Recallentry
type Recallentry struct { 
    // NbrAttempts
    NbrAttempts int `json:"nbrAttempts"`


    // MinutesBetweenAttempts
    MinutesBetweenAttempts int `json:"minutesBetweenAttempts"`

}

// String returns a JSON representation of the model
func (o *Recallentry) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recallentry) MarshalJSON() ([]byte, error) {
    type Alias Recallentry

    if RecallentryMarshalled {
        return []byte("{}"), nil
    }
    RecallentryMarshalled = true

    return json.Marshal(&struct {
        
        NbrAttempts int `json:"nbrAttempts"`
        
        MinutesBetweenAttempts int `json:"minutesBetweenAttempts"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

