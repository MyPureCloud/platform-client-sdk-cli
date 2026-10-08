package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappembeddedsignupintegrationactivationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappembeddedsignupintegrationactivationrequestDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Whatsappembeddedsignupintegrationactivationrequest
type Whatsappembeddedsignupintegrationactivationrequest struct { 
    


    // Name - The WhatsApp integration name. Required for Embedded Signup v4 activation; not used in v2.
    Name string `json:"name"`


    // PhoneNumber - E.164 phone number to associate with the WhatsApp integration. Not required for embedded signup v4 or later.
    PhoneNumber string `json:"phoneNumber"`


    // Pin - Specify the two-step verification PIN for that phone number
    Pin string `json:"pin"`


    

}

// String returns a JSON representation of the model
func (o *Whatsappembeddedsignupintegrationactivationrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappembeddedsignupintegrationactivationrequest) MarshalJSON() ([]byte, error) {
    type Alias Whatsappembeddedsignupintegrationactivationrequest

    if WhatsappembeddedsignupintegrationactivationrequestMarshalled {
        return []byte("{}"), nil
    }
    WhatsappembeddedsignupintegrationactivationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        Pin string `json:"pin"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

