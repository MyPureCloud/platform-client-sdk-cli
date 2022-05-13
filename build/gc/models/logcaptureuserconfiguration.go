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


    


    SelfUri string `json:"selfUri"`

}

// Logcaptureuserconfiguration
type Logcaptureuserconfiguration struct { 
    


    // DateExpired - Indicates when log capture will be turned off for the user. (Must be within 24 hours). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpired time.Time `json:"dateExpired"`


    

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
        
        DateExpired time.Time `json:"dateExpired"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

