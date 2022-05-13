package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonenumberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonenumberDud struct { 
    


    


    


    


    


    

}

// Phonenumber
type Phonenumber struct { 
    // Display
    Display string `json:"display"`


    // Extension
    Extension int `json:"extension"`


    // AcceptsSMS
    AcceptsSMS bool `json:"acceptsSMS"`


    // UserInput
    UserInput string `json:"userInput"`


    // E164
    E164 string `json:"e164"`


    // CountryCode
    CountryCode string `json:"countryCode"`

}

// String returns a JSON representation of the model
func (o *Phonenumber) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonenumber) MarshalJSON() ([]byte, error) {
    type Alias Phonenumber

    if PhonenumberMarshalled {
        return []byte("{}"), nil
    }
    PhonenumberMarshalled = true

    return json.Marshal(&struct {
        
        Display string `json:"display"`
        
        Extension int `json:"extension"`
        
        AcceptsSMS bool `json:"acceptsSMS"`
        
        UserInput string `json:"userInput"`
        
        E164 string `json:"e164"`
        
        CountryCode string `json:"countryCode"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

