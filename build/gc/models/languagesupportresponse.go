package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LanguagesupportresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LanguagesupportresponseDud struct { 
    Languages []Languagesupportinforecord `json:"languages"`

}

// Languagesupportresponse
type Languagesupportresponse struct { 
    

}

// String returns a JSON representation of the model
func (o *Languagesupportresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Languagesupportresponse) MarshalJSON() ([]byte, error) {
    type Alias Languagesupportresponse

    if LanguagesupportresponseMarshalled {
        return []byte("{}"), nil
    }
    LanguagesupportresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

