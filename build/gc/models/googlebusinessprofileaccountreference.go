package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GooglebusinessprofileaccountreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GooglebusinessprofileaccountreferenceDud struct { 
    

}

// Googlebusinessprofileaccountreference
type Googlebusinessprofileaccountreference struct { 
    // Id - ID of the Google Business Profile account
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Googlebusinessprofileaccountreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googlebusinessprofileaccountreference) MarshalJSON() ([]byte, error) {
    type Alias Googlebusinessprofileaccountreference

    if GooglebusinessprofileaccountreferenceMarshalled {
        return []byte("{}"), nil
    }
    GooglebusinessprofileaccountreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

