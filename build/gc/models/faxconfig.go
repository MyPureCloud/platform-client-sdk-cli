package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FaxconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FaxconfigDud struct { 
    


    

}

// Faxconfig
type Faxconfig struct { 
    // SendEmailNotifications - Whether to enable email notifications for this organization
    SendEmailNotifications bool `json:"sendEmailNotifications"`


    // DisableEmailPii - Whether to disable PII for email notifications
    DisableEmailPii bool `json:"disableEmailPii"`

}

// String returns a JSON representation of the model
func (o *Faxconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Faxconfig) MarshalJSON() ([]byte, error) {
    type Alias Faxconfig

    if FaxconfigMarshalled {
        return []byte("{}"), nil
    }
    FaxconfigMarshalled = true

    return json.Marshal(&struct {
        
        SendEmailNotifications bool `json:"sendEmailNotifications"`
        
        DisableEmailPii bool `json:"disableEmailPii"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

