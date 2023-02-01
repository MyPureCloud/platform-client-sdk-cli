package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsphonenumberprovisionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsphonenumberprovisionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Smsphonenumberprovision
type Smsphonenumberprovision struct { 
    


    // PhoneNumber - A phone number to be used for SMS communications. E.g. +13175555555 or +34234234234
    PhoneNumber string `json:"phoneNumber"`


    // PhoneNumberType - Type of the phone number provisioned.
    PhoneNumberType string `json:"phoneNumberType"`


    // CountryCode - The ISO 3166-1 alpha-2 country code of the country this phone number is associated with.
    CountryCode string `json:"countryCode"`


    // Name
    Name string `json:"name"`


    // AddressId - The id of an address added on your account. Due to regulatory requirements in some countries, an address may be required when provisioning a sms number. In those cases you should provide the provisioned sms address id here
    AddressId string `json:"addressId"`


    

}

// String returns a JSON representation of the model
func (o *Smsphonenumberprovision) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsphonenumberprovision) MarshalJSON() ([]byte, error) {
    type Alias Smsphonenumberprovision

    if SmsphonenumberprovisionMarshalled {
        return []byte("{}"), nil
    }
    SmsphonenumberprovisionMarshalled = true

    return json.Marshal(&struct {
        
        PhoneNumber string `json:"phoneNumber"`
        
        PhoneNumberType string `json:"phoneNumberType"`
        
        CountryCode string `json:"countryCode"`
        
        Name string `json:"name"`
        
        AddressId string `json:"addressId"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

