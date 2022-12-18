package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DncpatchphonenumbersrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DncpatchphonenumbersrequestDud struct { 
    


    


    

}

// Dncpatchphonenumbersrequest
type Dncpatchphonenumbersrequest struct { 
    // Action - The action to perform
    Action string `json:"action"`


    // PhoneNumbers - The list of phone numbers to Add to / Remove from the DNC list 
    PhoneNumbers []string `json:"phoneNumbers"`


    // ExpirationDateTime - Expiration date for DNC phone numbers in yyyy-MM-ddTHH:mmZ format
    ExpirationDateTime string `json:"expirationDateTime"`

}

// String returns a JSON representation of the model
func (o *Dncpatchphonenumbersrequest) String() string {
    
     o.PhoneNumbers = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dncpatchphonenumbersrequest) MarshalJSON() ([]byte, error) {
    type Alias Dncpatchphonenumbersrequest

    if DncpatchphonenumbersrequestMarshalled {
        return []byte("{}"), nil
    }
    DncpatchphonenumbersrequestMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        
        PhoneNumbers []string `json:"phoneNumbers"`
        
        ExpirationDateTime string `json:"expirationDateTime"`
        *Alias
    }{

        


        
        PhoneNumbers: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

