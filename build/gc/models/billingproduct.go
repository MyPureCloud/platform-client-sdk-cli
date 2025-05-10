package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingproductMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingproductDud struct { 
    Id string `json:"id"`


    


    


    

}

// Billingproduct
type Billingproduct struct { 
    


    // Sku - The product associated with the fund.
    Sku string `json:"sku"`


    // Name - The name of the product.
    Name string `json:"name"`


    // UnitOfMeasure - The unit of measure for the product.
    UnitOfMeasure string `json:"unitOfMeasure"`

}

// String returns a JSON representation of the model
func (o *Billingproduct) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingproduct) MarshalJSON() ([]byte, error) {
    type Alias Billingproduct

    if BillingproductMarshalled {
        return []byte("{}"), nil
    }
    BillingproductMarshalled = true

    return json.Marshal(&struct {
        
        Sku string `json:"sku"`
        
        Name string `json:"name"`
        
        UnitOfMeasure string `json:"unitOfMeasure"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

