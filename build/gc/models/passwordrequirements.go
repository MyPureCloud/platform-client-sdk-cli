package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PasswordrequirementsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PasswordrequirementsDud struct { 
    


    


    


    


    


    


    


    

}

// Passwordrequirements
type Passwordrequirements struct { 
    // MinimumLength
    MinimumLength int `json:"minimumLength"`


    // MinimumDigits
    MinimumDigits int `json:"minimumDigits"`


    // MinimumLetters
    MinimumLetters int `json:"minimumLetters"`


    // MinimumUpper
    MinimumUpper int `json:"minimumUpper"`


    // MinimumLower
    MinimumLower int `json:"minimumLower"`


    // MinimumSpecials
    MinimumSpecials int `json:"minimumSpecials"`


    // MinimumAgeSeconds
    MinimumAgeSeconds int `json:"minimumAgeSeconds"`


    // ExpirationDays
    ExpirationDays int `json:"expirationDays"`

}

// String returns a JSON representation of the model
func (o *Passwordrequirements) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Passwordrequirements) MarshalJSON() ([]byte, error) {
    type Alias Passwordrequirements

    if PasswordrequirementsMarshalled {
        return []byte("{}"), nil
    }
    PasswordrequirementsMarshalled = true

    return json.Marshal(&struct {
        
        MinimumLength int `json:"minimumLength"`
        
        MinimumDigits int `json:"minimumDigits"`
        
        MinimumLetters int `json:"minimumLetters"`
        
        MinimumUpper int `json:"minimumUpper"`
        
        MinimumLower int `json:"minimumLower"`
        
        MinimumSpecials int `json:"minimumSpecials"`
        
        MinimumAgeSeconds int `json:"minimumAgeSeconds"`
        
        ExpirationDays int `json:"expirationDays"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

