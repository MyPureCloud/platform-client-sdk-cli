package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailerrorinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailerrorinfoDud struct { 
    Message string `json:"message"`


    Code string `json:"code"`

}

// Emailerrorinfo
type Emailerrorinfo struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Emailerrorinfo) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailerrorinfo) MarshalJSON() ([]byte, error) {
    type Alias Emailerrorinfo

    if EmailerrorinfoMarshalled {
        return []byte("{}"), nil
    }
    EmailerrorinfoMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

