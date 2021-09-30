package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsphonenumberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsphonenumberDud struct { 
    Id string `json:"id"`


    


    


    PhoneNumberType string `json:"phoneNumberType"`


    


    


    Capabilities []string `json:"capabilities"`


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Smsphonenumber
type Smsphonenumber struct { 
    


    // Name
    Name string `json:"name"`


    // PhoneNumber - A phone number provisioned for SMS communications in E.164 format. E.g. +13175555555 or +34234234234
    PhoneNumber string `json:"phoneNumber"`


    


    // ProvisionedThroughPureCloud - Is set to false, if the phone number is provisioned through a SMS provider, outside of PureCloud
    ProvisionedThroughPureCloud bool `json:"provisionedThroughPureCloud"`


    // PhoneNumberStatus - Status of the provisioned phone number.
    PhoneNumberStatus string `json:"phoneNumberStatus"`


    


    // CountryCode - The ISO 3166-1 alpha-2 country code of the country this phone number is associated with.
    CountryCode string `json:"countryCode"`


    // DateCreated - Date this phone number was provisioned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date this phone number was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // CreatedBy - User that provisioned this phone number
    CreatedBy User `json:"createdBy"`


    // ModifiedBy - User that last modified this phone number
    ModifiedBy User `json:"modifiedBy"`


    // Version - Version number required for updates.
    Version int `json:"version"`


    // PurchaseDate - Date this phone number was purchased, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    PurchaseDate time.Time `json:"purchaseDate"`


    // CancellationDate - Contract end date of this phone number, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CancellationDate time.Time `json:"cancellationDate"`


    // RenewalDate - Contract renewal date of this phone number, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RenewalDate time.Time `json:"renewalDate"`


    // AutoRenewable - Renewal time period of this phone number, if the phoneNumberType is shortcode.
    AutoRenewable string `json:"autoRenewable"`


    // AddressId - The id of an address attached to this phone number.
    AddressId Smsaddress `json:"addressId"`


    // ShortCodeBillingType - BillingType of this phone number, if the phoneNumberType is shortcode.
    ShortCodeBillingType string `json:"shortCodeBillingType"`


    

}

// String returns a JSON representation of the model
func (o *Smsphonenumber) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsphonenumber) MarshalJSON() ([]byte, error) {
    type Alias Smsphonenumber

    if SmsphonenumberMarshalled {
        return []byte("{}"), nil
    }
    SmsphonenumberMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        
        
        ProvisionedThroughPureCloud bool `json:"provisionedThroughPureCloud"`
        
        PhoneNumberStatus string `json:"phoneNumberStatus"`
        
        
        
        CountryCode string `json:"countryCode"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        CreatedBy User `json:"createdBy"`
        
        ModifiedBy User `json:"modifiedBy"`
        
        Version int `json:"version"`
        
        PurchaseDate time.Time `json:"purchaseDate"`
        
        CancellationDate time.Time `json:"cancellationDate"`
        
        RenewalDate time.Time `json:"renewalDate"`
        
        AutoRenewable string `json:"autoRenewable"`
        
        AddressId Smsaddress `json:"addressId"`
        
        ShortCodeBillingType string `json:"shortCodeBillingType"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

