package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingchargeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingchargeDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Billingcharge
type Billingcharge struct { 
    


    


    // Product - Represents the details of a product.
    Product Billingproduct `json:"product"`


    // Organizations - Organizations associated with this charge.
    Organizations []Namedentity `json:"organizations"`


    // PrepaidQuantity - The quantity of usage that is prepaid.
    PrepaidQuantity float32 `json:"prepaidQuantity"`


    // FairuseAllocation - Quantity of usage allowed under fair-use policies.
    FairuseAllocation float32 `json:"fairuseAllocation"`


    // FairuseAmount - Monetary value of the fair-use allocation, in the same currency as overageCurrency.
    FairuseAmount float32 `json:"fairuseAmount"`


    // ActualQuantity - The actual quantity of usage.
    ActualQuantity float32 `json:"actualQuantity"`


    // OverageQuantity - The quantity of usage that exceeds prepaid or fair use limits.
    OverageQuantity float32 `json:"overageQuantity"`


    // OverageRate - The rate charged per unit of overage.
    OverageRate float32 `json:"overageRate"`


    // OverageCharge - The total charge for overage usage.
    OverageCharge float32 `json:"overageCharge"`


    // OverageCurrency - The currency in which the overage charge is billed.
    OverageCurrency string `json:"overageCurrency"`


    // ConvertedTokenCount - AI Token Count after the conversion for AI Products.
    ConvertedTokenCount float32 `json:"convertedTokenCount"`


    // TokenConversionRate - AI Token Conversion Rate — Units per token conversion ratio.
    TokenConversionRate float32 `json:"tokenConversionRate"`


    // AiSubItems - Usage that makes up this AI parent bundle, including the bundle's own token usage and each child product's usage converted to tokens; the converted total is the parent's actual quantity.
    AiSubItems []Billingaisubitem `json:"aiSubItems"`

}

// String returns a JSON representation of the model
func (o *Billingcharge) String() string {
    
     o.Organizations = []Namedentity{{}} 
    
    
    
    
    
    
    
    
    
    
     o.AiSubItems = []Billingaisubitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingcharge) MarshalJSON() ([]byte, error) {
    type Alias Billingcharge

    if BillingchargeMarshalled {
        return []byte("{}"), nil
    }
    BillingchargeMarshalled = true

    return json.Marshal(&struct {
        
        Product Billingproduct `json:"product"`
        
        Organizations []Namedentity `json:"organizations"`
        
        PrepaidQuantity float32 `json:"prepaidQuantity"`
        
        FairuseAllocation float32 `json:"fairuseAllocation"`
        
        FairuseAmount float32 `json:"fairuseAmount"`
        
        ActualQuantity float32 `json:"actualQuantity"`
        
        OverageQuantity float32 `json:"overageQuantity"`
        
        OverageRate float32 `json:"overageRate"`
        
        OverageCharge float32 `json:"overageCharge"`
        
        OverageCurrency string `json:"overageCurrency"`
        
        ConvertedTokenCount float32 `json:"convertedTokenCount"`
        
        TokenConversionRate float32 `json:"tokenConversionRate"`
        
        AiSubItems []Billingaisubitem `json:"aiSubItems"`
        *Alias
    }{

        


        


        


        
        Organizations: []Namedentity{{}},
        


        


        


        


        


        


        


        


        


        


        


        
        AiSubItems: []Billingaisubitem{{}},
        

        Alias: (*Alias)(u),
    })
}

