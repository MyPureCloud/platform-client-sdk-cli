package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsalphanumericprovisionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsalphanumericprovisionDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Smsalphanumericprovision
type Smsalphanumericprovision struct { 
    


    // PhoneNumber - A phone number to be used for SMS communications. E.g. Genesys123
    PhoneNumber string `json:"phoneNumber"`


    

}

// String returns a JSON representation of the model
func (o *Smsalphanumericprovision) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsalphanumericprovision) MarshalJSON() ([]byte, error) {
    type Alias Smsalphanumericprovision

    if SmsalphanumericprovisionMarshalled {
        return []byte("{}"), nil
    }
    SmsalphanumericprovisionMarshalled = true

    return json.Marshal(&struct {
        
        PhoneNumber string `json:"phoneNumber"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

