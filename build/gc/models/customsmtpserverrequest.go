package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomsmtpserverrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomsmtpserverrequestDud struct { 
    

}

// Customsmtpserverrequest
type Customsmtpserverrequest struct { 
    // Id - The ID of the integration that contains the SMTP configuration. 
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Customsmtpserverrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customsmtpserverrequest) MarshalJSON() ([]byte, error) {
    type Alias Customsmtpserverrequest

    if CustomsmtpserverrequestMarshalled {
        return []byte("{}"), nil
    }
    CustomsmtpserverrequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

