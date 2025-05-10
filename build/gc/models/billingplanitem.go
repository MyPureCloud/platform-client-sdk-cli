package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingplanitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingplanitemDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Billingplanitem
type Billingplanitem struct { 
    // ItemNumber - Item number.
    ItemNumber string `json:"itemNumber"`


    // Name - Name of the item.
    Name string `json:"name"`


    // VarType - Type of the item.
    VarType string `json:"type"`


    // Function - Function of the item.
    Function string `json:"function"`


    // Description - Detailed description of the item.
    Description string `json:"description"`


    // DateChargedThrough - The date through which a customer has been billed for the charge. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateChargedThrough time.Time `json:"dateChargedThrough"`


    // CurrencyIsoCode - Contains the ISO code for any currency allowed by the organization.
    CurrencyIsoCode string `json:"currencyIsoCode"`


    // DiscountAmount - The amount of the discount.
    DiscountAmount float32 `json:"discountAmount"`


    // DateEffectiveStart - The date when the Address became effective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEffectiveStart time.Time `json:"dateEffectiveStart"`


    // DateEffectiveEnd - The date when the Address became effective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEffectiveEnd time.Time `json:"dateEffectiveEnd"`


    // OveragePrice - The price for units over the allowed amount.
    OveragePrice float32 `json:"overagePrice"`


    // Price - The price associated with the item expressed as a decimal.
    Price float32 `json:"price"`


    // Quantity - The quantity of units.
    Quantity int `json:"quantity"`


    // UnitOfMeasure - The unit of measure for the wallet.
    UnitOfMeasure string `json:"unitOfMeasure"`

}

// String returns a JSON representation of the model
func (o *Billingplanitem) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingplanitem) MarshalJSON() ([]byte, error) {
    type Alias Billingplanitem

    if BillingplanitemMarshalled {
        return []byte("{}"), nil
    }
    BillingplanitemMarshalled = true

    return json.Marshal(&struct {
        
        ItemNumber string `json:"itemNumber"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Function string `json:"function"`
        
        Description string `json:"description"`
        
        DateChargedThrough time.Time `json:"dateChargedThrough"`
        
        CurrencyIsoCode string `json:"currencyIsoCode"`
        
        DiscountAmount float32 `json:"discountAmount"`
        
        DateEffectiveStart time.Time `json:"dateEffectiveStart"`
        
        DateEffectiveEnd time.Time `json:"dateEffectiveEnd"`
        
        OveragePrice float32 `json:"overagePrice"`
        
        Price float32 `json:"price"`
        
        Quantity int `json:"quantity"`
        
        UnitOfMeasure string `json:"unitOfMeasure"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

