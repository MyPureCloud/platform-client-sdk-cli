package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SetcontenttemplateactionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SetcontenttemplateactionsettingsDud struct { 
    


    

}

// Setcontenttemplateactionsettings
type Setcontenttemplateactionsettings struct { 
    // SmsContentTemplateId - A string of sms contentTemplateId.
    SmsContentTemplateId string `json:"smsContentTemplateId"`


    // EmailContentTemplateId - A string of email contentTemplateId.
    EmailContentTemplateId string `json:"emailContentTemplateId"`

}

// String returns a JSON representation of the model
func (o *Setcontenttemplateactionsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Setcontenttemplateactionsettings) MarshalJSON() ([]byte, error) {
    type Alias Setcontenttemplateactionsettings

    if SetcontenttemplateactionsettingsMarshalled {
        return []byte("{}"), nil
    }
    SetcontenttemplateactionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        SmsContentTemplateId string `json:"smsContentTemplateId"`
        
        EmailContentTemplateId string `json:"emailContentTemplateId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

