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


    // GetprepaidQuantity - The quantity of usage that is prepaid.
    GetprepaidQuantity int `json:"getprepaidQuantity"`


    // GetfairuseQuantity - The quantity of usage allowed under fair use policies.
    GetfairuseQuantity int `json:"getfairuseQuantity"`


    // GetactualQuantity - The actual quantity of usage.
    GetactualQuantity int `json:"getactualQuantity"`


    // GetoverageQuantity - The quantity of usage that exceeds prepaid or fair use limits.
    GetoverageQuantity int `json:"getoverageQuantity"`


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
        
        GetprepaidQuantity int `json:"getprepaidQuantity"`
        
        GetfairuseQuantity int `json:"getfairuseQuantity"`
        
        GetactualQuantity int `json:"getactualQuantity"`
        
        GetoverageQuantity int `json:"getoverageQuantity"`
        
        OverageRate float32 `json:"overageRate"`
        
        OverageCharge float32 `json:"overageCharge"`
        
        OverageCurrency string `json:"overageCurrency"`
        *Alias
    }{

        


        


        
        Organizations: []Namedentity{{}},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

