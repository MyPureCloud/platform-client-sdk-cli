package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappidDud struct { 
    


    

}

// Whatsappid - User information for a WhatsApp account
type Whatsappid struct { 
    // PhoneNumber - The phone number associated with this WhatsApp account. Requires 'E164 without a leading plus' phone number.
    PhoneNumber Phonenumber `json:"phoneNumber"`


    // DisplayName - The displayName of this person's account in WhatsApp. Max: 100 characters.
    DisplayName string `json:"displayName"`

}

// String returns a JSON representation of the model
func (o *Whatsappid) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappid) MarshalJSON() ([]byte, error) {
    type Alias Whatsappid

    if WhatsappidMarshalled {
        return []byte("{}"), nil
    }
    WhatsappidMarshalled = true

    return json.Marshal(&struct {
        
        PhoneNumber Phonenumber `json:"phoneNumber"`
        
        DisplayName string `json:"displayName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

