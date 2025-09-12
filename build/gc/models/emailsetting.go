package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailsettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailsettingDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Emailsetting - Email routing settings.
type Emailsetting struct { 
    


    // Name
    Name string `json:"name"`


    // Domains - The domain list settings.
    Domains Domains `json:"domains"`


    

}

// String returns a JSON representation of the model
func (o *Emailsetting) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailsetting) MarshalJSON() ([]byte, error) {
    type Alias Emailsetting

    if EmailsettingMarshalled {
        return []byte("{}"), nil
    }
    EmailsettingMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Domains Domains `json:"domains"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

