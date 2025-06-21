package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PaymentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PaymentrequestDud struct { 
    


    


    


    


    


    

}

// Paymentrequest
type Paymentrequest struct { 
    // PaymentPlatform - The payment platform being used (e.g. Apple Pay)
    PaymentPlatform string `json:"paymentPlatform"`


    // CountryCode - The merchant's two-letter ISO 3166 country code.
    CountryCode string `json:"countryCode"`


    // CurrencyCode - The three-letter ISO 4217 currency code for the payment.
    CurrencyCode string `json:"currencyCode"`


    // OrderTotal - The total price of the order.
    OrderTotal float64 `json:"orderTotal"`


    // LineItems - The items that make up the order.
    LineItems []Paymentlineitem `json:"lineItems"`


    // ShippingOptions - The available shipping options.
    ShippingOptions []Paymentlineitem `json:"shippingOptions"`

}

// String returns a JSON representation of the model
func (o *Paymentrequest) String() string {
    
    
    
    
     o.LineItems = []Paymentlineitem{{}} 
     o.ShippingOptions = []Paymentlineitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Paymentrequest) MarshalJSON() ([]byte, error) {
    type Alias Paymentrequest

    if PaymentrequestMarshalled {
        return []byte("{}"), nil
    }
    PaymentrequestMarshalled = true

    return json.Marshal(&struct {
        
        PaymentPlatform string `json:"paymentPlatform"`
        
        CountryCode string `json:"countryCode"`
        
        CurrencyCode string `json:"currencyCode"`
        
        OrderTotal float64 `json:"orderTotal"`
        
        LineItems []Paymentlineitem `json:"lineItems"`
        
        ShippingOptions []Paymentlineitem `json:"shippingOptions"`
        *Alias
    }{

        


        


        


        


        
        LineItems: []Paymentlineitem{{}},
        


        
        ShippingOptions: []Paymentlineitem{{}},
        

        Alias: (*Alias)(u),
    })
}

