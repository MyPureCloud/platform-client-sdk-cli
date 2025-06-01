package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LogcaptureuserconfigurationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LogcaptureuserconfigurationresponseDud struct { 
    Id string `json:"id"`


    DateStarted time.Time `json:"dateStarted"`


    


    SelfUri string `json:"selfUri"`

}

// Logcaptureuserconfigurationresponse
type Logcaptureuserconfigurationresponse struct { 
    


    


    // DateExpired - Indicates when log capture will be turned off for the user. (Must be within 24 hours). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpired time.Time `json:"dateExpired"`


    

}

// String returns a JSON representation of the model
func (o *Logcaptureuserconfigurationresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Logcaptureuserconfigurationresponse) MarshalJSON() ([]byte, error) {
    type Alias Logcaptureuserconfigurationresponse

    if LogcaptureuserconfigurationresponseMarshalled {
        return []byte("{}"), nil
    }
    LogcaptureuserconfigurationresponseMarshalled = true

    return json.Marshal(&struct {
        
        DateExpired time.Time `json:"dateExpired"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

