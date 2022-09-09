package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LastresultoverallconditionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LastresultoverallconditionsettingsDud struct { 
    


    

}

// Lastresultoverallconditionsettings
type Lastresultoverallconditionsettings struct { 
    // EmailWrapupCodes - A list of wrapup code identifiers to match for Email.
    EmailWrapupCodes []string `json:"emailWrapupCodes"`


    // SmsWrapupCodes - A list of wrapup code identifiers to match for SMS.
    SmsWrapupCodes []string `json:"smsWrapupCodes"`

}

// String returns a JSON representation of the model
func (o *Lastresultoverallconditionsettings) String() string {
     o.EmailWrapupCodes = []string{""} 
     o.SmsWrapupCodes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lastresultoverallconditionsettings) MarshalJSON() ([]byte, error) {
    type Alias Lastresultoverallconditionsettings

    if LastresultoverallconditionsettingsMarshalled {
        return []byte("{}"), nil
    }
    LastresultoverallconditionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        EmailWrapupCodes []string `json:"emailWrapupCodes"`
        
        SmsWrapupCodes []string `json:"smsWrapupCodes"`
        *Alias
    }{

        
        EmailWrapupCodes: []string{""},
        


        
        SmsWrapupCodes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

