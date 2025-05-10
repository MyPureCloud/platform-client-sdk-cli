package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InvoiceaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InvoiceaddressDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Invoiceaddress
type Invoiceaddress struct { 
    // GetdateEffective - The date when the Address became effective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    GetdateEffective time.Time `json:"getdateEffective"`


    // AddressType - The type of address.
    AddressType string `json:"addressType"`


    // CurrencyIsoCode - Contains the ISO code for any currency allowed by the organization.
    CurrencyIsoCode string `json:"currencyIsoCode"`


    // Line1 - The first line of the Address.
    Line1 string `json:"line1"`


    // Line2 - The second line of the Address.
    Line2 string `json:"line2"`


    // Line3 - The third line of the Address.
    Line3 string `json:"line3"`


    // CityName - The city name.
    CityName string `json:"cityName"`


    // PostalCode - The Postal or Zip Code.
    PostalCode string `json:"postalCode"`


    // StateCode - The code that reflects the geographic state for the Address.
    StateCode string `json:"stateCode"`


    // CountryCode - The code representing the country for the Address (ISO_3166).
    CountryCode string `json:"countryCode"`


    // GetcitySubdivision1 - The primary subdivision within a city (e.g., district, neighborhood).
    GetcitySubdivision1 string `json:"getcitySubdivision1"`


    // RegionSubdivision1 - The primary administrative division within a region (e.g., state, province).
    RegionSubdivision1 string `json:"regionSubdivision1"`


    // RegionSubdivision2 - A secondary subdivision within the primary region subdivision (e.g., county, district).
    RegionSubdivision2 string `json:"regionSubdivision2"`


    // Country - Stores the name of the country in which the address is situated.
    Country string `json:"country"`

}

// String returns a JSON representation of the model
func (o *Invoiceaddress) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Invoiceaddress) MarshalJSON() ([]byte, error) {
    type Alias Invoiceaddress

    if InvoiceaddressMarshalled {
        return []byte("{}"), nil
    }
    InvoiceaddressMarshalled = true

    return json.Marshal(&struct {
        
        GetdateEffective time.Time `json:"getdateEffective"`
        
        AddressType string `json:"addressType"`
        
        CurrencyIsoCode string `json:"currencyIsoCode"`
        
        Line1 string `json:"line1"`
        
        Line2 string `json:"line2"`
        
        Line3 string `json:"line3"`
        
        CityName string `json:"cityName"`
        
        PostalCode string `json:"postalCode"`
        
        StateCode string `json:"stateCode"`
        
        CountryCode string `json:"countryCode"`
        
        GetcitySubdivision1 string `json:"getcitySubdivision1"`
        
        RegionSubdivision1 string `json:"regionSubdivision1"`
        
        RegionSubdivision2 string `json:"regionSubdivision2"`
        
        Country string `json:"country"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

