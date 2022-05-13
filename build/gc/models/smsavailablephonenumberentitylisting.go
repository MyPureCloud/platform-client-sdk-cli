package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsavailablephonenumberentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsavailablephonenumberentitylistingDud struct { 
    

}

// Smsavailablephonenumberentitylisting
type Smsavailablephonenumberentitylisting struct { 
    // Entities
    Entities []Smsavailablephonenumber `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Smsavailablephonenumberentitylisting) String() string {
     o.Entities = []Smsavailablephonenumber{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsavailablephonenumberentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Smsavailablephonenumberentitylisting

    if SmsavailablephonenumberentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SmsavailablephonenumberentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Smsavailablephonenumber `json:"entities"`
        *Alias
    }{

        
        Entities: []Smsavailablephonenumber{{}},
        

        Alias: (*Alias)(u),
    })
}

