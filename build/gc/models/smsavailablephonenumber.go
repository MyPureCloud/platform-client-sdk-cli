package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsavailablephonenumberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsavailablephonenumberDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Smsavailablephonenumber
type Smsavailablephonenumber struct { 
    


    // Name
    Name string `json:"name"`


    // PhoneNumber - A phone number available for provisioning in E.164 format. E.g. +13175555555 or +34234234234
    PhoneNumber string `json:"phoneNumber"`


    // CountryCode - The ISO 3166-1 alpha-2 country code of the country this phone number is associated with.
    CountryCode string `json:"countryCode"`


    // Region - The region/province/state the phone number is associated with.
    Region string `json:"region"`


    // City - The city the phone number is associated with.
    City string `json:"city"`


    // Capabilities - The capabilities of the phone number available for provisioning.
    Capabilities []string `json:"capabilities"`


    // PhoneNumberType - The type of phone number available for provisioning.
    PhoneNumberType string `json:"phoneNumberType"`


    // AddressRequirement - The address requirement needed for provisioning this number. If there is a requirement, the address must be the residence or place of business of the individual or entity using the phone number.
    AddressRequirement string `json:"addressRequirement"`


    

}

// String returns a JSON representation of the model
func (o *Smsavailablephonenumber) String() string {
    
    
    
    
    
     o.Capabilities = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsavailablephonenumber) MarshalJSON() ([]byte, error) {
    type Alias Smsavailablephonenumber

    if SmsavailablephonenumberMarshalled {
        return []byte("{}"), nil
    }
    SmsavailablephonenumberMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        CountryCode string `json:"countryCode"`
        
        Region string `json:"region"`
        
        City string `json:"city"`
        
        Capabilities []string `json:"capabilities"`
        
        PhoneNumberType string `json:"phoneNumberType"`
        
        AddressRequirement string `json:"addressRequirement"`
        *Alias
    }{

        


        


        


        


        


        


        
        Capabilities: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

