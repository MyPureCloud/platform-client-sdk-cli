package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DncpatchemailsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DncpatchemailsrequestDud struct { 
    


    


    

}

// Dncpatchemailsrequest
type Dncpatchemailsrequest struct { 
    // Action - The action to perform
    Action string `json:"action"`


    // EmailAddresses - The list of email addresses to Add to / Remove from the DNC list 
    EmailAddresses []string `json:"emailAddresses"`


    // ExpirationDateTime - Expiration date for DNC email addresses in yyyy-MM-ddTHH:mmZ format
    ExpirationDateTime string `json:"expirationDateTime"`

}

// String returns a JSON representation of the model
func (o *Dncpatchemailsrequest) String() string {
    
     o.EmailAddresses = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dncpatchemailsrequest) MarshalJSON() ([]byte, error) {
    type Alias Dncpatchemailsrequest

    if DncpatchemailsrequestMarshalled {
        return []byte("{}"), nil
    }
    DncpatchemailsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        
        EmailAddresses []string `json:"emailAddresses"`
        
        ExpirationDateTime string `json:"expirationDateTime"`
        *Alias
    }{

        


        
        EmailAddresses: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

