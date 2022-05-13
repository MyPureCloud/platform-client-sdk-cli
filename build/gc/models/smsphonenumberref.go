package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsphonenumberrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsphonenumberrefDud struct { 
    


    

}

// Smsphonenumberref
type Smsphonenumberref struct { 
    // PhoneNumber - A phone number provisioned for SMS communications in E.164 format. E.g. +13175555555 or +34234234234
    PhoneNumber string `json:"phoneNumber"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Smsphonenumberref) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsphonenumberref) MarshalJSON() ([]byte, error) {
    type Alias Smsphonenumberref

    if SmsphonenumberrefMarshalled {
        return []byte("{}"), nil
    }
    SmsphonenumberrefMarshalled = true

    return json.Marshal(&struct {
        
        PhoneNumber string `json:"phoneNumber"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

