package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailsettingreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailsettingreferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Emailsettingreference - Email Setting reference for email routing settings
type Emailsettingreference struct { 
    // Id - The email setting unique identifier
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Emailsettingreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailsettingreference) MarshalJSON() ([]byte, error) {
    type Alias Emailsettingreference

    if EmailsettingreferenceMarshalled {
        return []byte("{}"), nil
    }
    EmailsettingreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

