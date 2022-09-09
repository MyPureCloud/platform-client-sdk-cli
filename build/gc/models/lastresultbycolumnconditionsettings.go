package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LastresultbycolumnconditionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LastresultbycolumnconditionsettingsDud struct { 
    


    


    


    

}

// Lastresultbycolumnconditionsettings
type Lastresultbycolumnconditionsettings struct { 
    // EmailColumnName - The name of the contact column to evaluate for Email.
    EmailColumnName string `json:"emailColumnName"`


    // EmailWrapupCodes - A list of wrapup code identifiers to match for Email.
    EmailWrapupCodes []string `json:"emailWrapupCodes"`


    // SmsColumnName - The name of the contact column to evaluate for SMS.
    SmsColumnName string `json:"smsColumnName"`


    // SmsWrapupCodes - A list of wrapup code identifiers to match for SMS.
    SmsWrapupCodes []string `json:"smsWrapupCodes"`

}

// String returns a JSON representation of the model
func (o *Lastresultbycolumnconditionsettings) String() string {
    
     o.EmailWrapupCodes = []string{""} 
    
     o.SmsWrapupCodes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lastresultbycolumnconditionsettings) MarshalJSON() ([]byte, error) {
    type Alias Lastresultbycolumnconditionsettings

    if LastresultbycolumnconditionsettingsMarshalled {
        return []byte("{}"), nil
    }
    LastresultbycolumnconditionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        EmailColumnName string `json:"emailColumnName"`
        
        EmailWrapupCodes []string `json:"emailWrapupCodes"`
        
        SmsColumnName string `json:"smsColumnName"`
        
        SmsWrapupCodes []string `json:"smsWrapupCodes"`
        *Alias
    }{

        


        
        EmailWrapupCodes: []string{""},
        


        


        
        SmsWrapupCodes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

