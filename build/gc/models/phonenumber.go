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
    // Display - The displayed form of the phone number string. Users should input the phone number in this field, but it will be altered by the API on write. If the phone number can be read as E164, the value will be replaced with international formatted-version of the number. If the number cannot be read as E164, the value will be preserved as-is. In both cases, the provided input string will be copied to the userInput field. Max: 512 characters.
    Display string `json:"display"`


    // Extension - An optional extension for the provided phone number.
    Extension int `json:"extension"`


    // AcceptsSMS - Whether this phone number can accept SMS messages.
    AcceptsSMS bool `json:"acceptsSMS"`


    // NormalizationCountryCode - The country code that will be used for E164 conversion of a provided phone number. If the country code is omitted from the provided phone number, the country code provided in this field will be used during the E164 conversion attempt. If this field is left empty, the default country code for any provided phone number that does not explicitly include a country code is assumed to be +1 (North America).
    NormalizationCountryCode string `json:"normalizationCountryCode"`


    // UserInput - The user-inputted phone number string that was provided to the display field on write. This field is not user-writeable and will always be set by the system.
    UserInput string `json:"userInput"`


    // E164 - The E164-formatted form of the provided phone number. This field is not user-writeable and will only be set when the provided phone number could be read as E164.
    E164 string `json:"e164"`


    // CountryCode - The detected country code from the provided phone number. This field is not user-writeable and will only be set when the provided phone number could be read as E164. Max: 4 characters.
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
        
        NormalizationCountryCode string `json:"normalizationCountryCode"`
        
        UserInput string `json:"userInput"`
        
        E164 string `json:"e164"`
        
        CountryCode string `json:"countryCode"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

