package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsphonenumberimportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsphonenumberimportDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Smsphonenumberimport
type Smsphonenumberimport struct { 
    


    // PhoneNumber - A phone number to be used for SMS communications. E.g. +13175555555 or +34234234234
    PhoneNumber string `json:"phoneNumber"`


    // PhoneNumberType - Type of the phone number provisioned.
    PhoneNumberType string `json:"phoneNumberType"`


    // CountryCode - The ISO 3166-1 alpha-2 country code of the country this phone number is associated with.
    CountryCode string `json:"countryCode"`


    // IntegrationId - The id of the Genesys Cloud integration this phone number belongs to.
    IntegrationId string `json:"integrationId"`


    // Compliance - Compliance configuration for short codes, including help, stop and opt in.
    Compliance Compliance `json:"compliance"`


    // SupportedContent - Defines the media SupportedContent profile configured for an MMS capable phone number.
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    

}

// String returns a JSON representation of the model
func (o *Smsphonenumberimport) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsphonenumberimport) MarshalJSON() ([]byte, error) {
    type Alias Smsphonenumberimport

    if SmsphonenumberimportMarshalled {
        return []byte("{}"), nil
    }
    SmsphonenumberimportMarshalled = true

    return json.Marshal(&struct {
        
        PhoneNumber string `json:"phoneNumber"`
        
        PhoneNumberType string `json:"phoneNumberType"`
        
        CountryCode string `json:"countryCode"`
        
        IntegrationId string `json:"integrationId"`
        
        Compliance Compliance `json:"compliance"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

