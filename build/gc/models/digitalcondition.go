package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DigitalconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DigitalconditionDud struct { 
    


    


    


    


    


    


    


    

}

// Digitalcondition
type Digitalcondition struct { 
    // Inverted - If true, inverts the result of evaluating this condition. Default is false.
    Inverted bool `json:"inverted"`


    // ContactColumnConditionSettings - The settings for a 'contact list column' condition.
    ContactColumnConditionSettings Contactcolumnconditionsettings `json:"contactColumnConditionSettings"`


    // ContactAddressConditionSettings - The settings for a 'contact address' condition.
    ContactAddressConditionSettings Contactaddressconditionsettings `json:"contactAddressConditionSettings"`


    // ContactAddressTypeConditionSettings - The settings for a 'contact address type' condition.
    ContactAddressTypeConditionSettings Contactaddresstypeconditionsettings `json:"contactAddressTypeConditionSettings"`


    // LastAttemptByColumnConditionSettings - The settings for a 'last attempt by column' condition.
    LastAttemptByColumnConditionSettings Lastattemptbycolumnconditionsettings `json:"lastAttemptByColumnConditionSettings"`


    // LastAttemptOverallConditionSettings - The settings for a 'last attempt overall' condition.
    LastAttemptOverallConditionSettings Lastattemptoverallconditionsettings `json:"lastAttemptOverallConditionSettings"`


    // LastResultByColumnConditionSettings - The settings for a 'last result by column' condition.
    LastResultByColumnConditionSettings Lastresultbycolumnconditionsettings `json:"lastResultByColumnConditionSettings"`


    // LastResultOverallConditionSettings - The settings for a 'last result overall' condition.
    LastResultOverallConditionSettings Lastresultoverallconditionsettings `json:"lastResultOverallConditionSettings"`

}

// String returns a JSON representation of the model
func (o *Digitalcondition) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Digitalcondition) MarshalJSON() ([]byte, error) {
    type Alias Digitalcondition

    if DigitalconditionMarshalled {
        return []byte("{}"), nil
    }
    DigitalconditionMarshalled = true

    return json.Marshal(&struct {
        
        Inverted bool `json:"inverted"`
        
        ContactColumnConditionSettings Contactcolumnconditionsettings `json:"contactColumnConditionSettings"`
        
        ContactAddressConditionSettings Contactaddressconditionsettings `json:"contactAddressConditionSettings"`
        
        ContactAddressTypeConditionSettings Contactaddresstypeconditionsettings `json:"contactAddressTypeConditionSettings"`
        
        LastAttemptByColumnConditionSettings Lastattemptbycolumnconditionsettings `json:"lastAttemptByColumnConditionSettings"`
        
        LastAttemptOverallConditionSettings Lastattemptoverallconditionsettings `json:"lastAttemptOverallConditionSettings"`
        
        LastResultByColumnConditionSettings Lastresultbycolumnconditionsettings `json:"lastResultByColumnConditionSettings"`
        
        LastResultOverallConditionSettings Lastresultoverallconditionsettings `json:"lastResultOverallConditionSettings"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

