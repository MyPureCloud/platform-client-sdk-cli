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


    


    


    


    


    


    


    


    


    

}

// Billingcharge
type Billingcharge struct { 
    


    // Product - Represents the details of a product.
    Product Billingproduct `json:"product"`


    // Organizations - List of plans within the organization.
    Organizations []Namedentity `json:"organizations"`


    // PrepaidQuantity - The quantity of usage that is prepaid.
    PrepaidQuantity int `json:"prepaidQuantity"`


    // FairuseQuantity - The quantity of usage allowed under fair use policies.
    FairuseQuantity int `json:"fairuseQuantity"`


    // ActualQuantity - The actual quantity of usage.
    ActualQuantity int `json:"actualQuantity"`


    // OverageQuantity - The quantity of usage that exceeds prepaid or fair use limits.
    OverageQuantity int `json:"overageQuantity"`


    // OverageRate - The rate charged per unit of overage.
    OverageRate float32 `json:"overageRate"`


    // OverageCharge - The total charge for overage usage.
    OverageCharge float32 `json:"overageCharge"`


    // OverageCurrency - The currency in which the overage charge is billed.
    OverageCurrency string `json:"overageCurrency"`

}

// String returns a JSON representation of the model
func (o *Billingcharge) String() string {
    
     o.Organizations = []Namedentity{{}} 
    
    
    
    
    
    
    

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
        
        PrepaidQuantity int `json:"prepaidQuantity"`
        
        FairuseQuantity int `json:"fairuseQuantity"`
        
        ActualQuantity int `json:"actualQuantity"`
        
        OverageQuantity int `json:"overageQuantity"`
        
        OverageRate float32 `json:"overageRate"`
        
        OverageCharge float32 `json:"overageCharge"`
        
        OverageCurrency string `json:"overageCurrency"`
        *Alias
    }{

        


        


        
        Organizations: []Namedentity{{}},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

