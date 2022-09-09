package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LastattemptbycolumnconditionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LastattemptbycolumnconditionsettingsDud struct { 
    


    


    


    

}

// Lastattemptbycolumnconditionsettings
type Lastattemptbycolumnconditionsettings struct { 
    // EmailColumnName - The name of the contact column to evaluate for Email.
    EmailColumnName string `json:"emailColumnName"`


    // SmsColumnName - The name of the contact column to evaluate for SMS.
    SmsColumnName string `json:"smsColumnName"`


    // Operator - The operator to use when comparing values.
    Operator string `json:"operator"`


    // Value - The period value to compare against the contact's data.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Lastattemptbycolumnconditionsettings) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lastattemptbycolumnconditionsettings) MarshalJSON() ([]byte, error) {
    type Alias Lastattemptbycolumnconditionsettings

    if LastattemptbycolumnconditionsettingsMarshalled {
        return []byte("{}"), nil
    }
    LastattemptbycolumnconditionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        EmailColumnName string `json:"emailColumnName"`
        
        SmsColumnName string `json:"smsColumnName"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

