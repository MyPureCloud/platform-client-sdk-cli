package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingcontractperioddetailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingcontractperioddetailDud struct { 
    Id string `json:"id"`


    


    


    


    


    

}

// Billingcontractperioddetail
type Billingcontractperioddetail struct { 
    


    // DateStart - The date when the Billing Period starts. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The date when the Billing Period starts. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEnd time.Time `json:"dateEnd"`


    // Status - The type of address.
    Status string `json:"status"`


    // Charges - Represents usage metrics including prepaid, actual, and overage quantities along with associated charges.
    Charges []Billingcharge `json:"charges"`


    // Wallets - Represents usage metrics including prepaid, actual, and overage quantities along with associated charges.
    Wallets []Billingwallet `json:"wallets"`

}

// String returns a JSON representation of the model
func (o *Billingcontractperioddetail) String() string {
    
    
    
     o.Charges = []Billingcharge{{}} 
     o.Wallets = []Billingwallet{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingcontractperioddetail) MarshalJSON() ([]byte, error) {
    type Alias Billingcontractperioddetail

    if BillingcontractperioddetailMarshalled {
        return []byte("{}"), nil
    }
    BillingcontractperioddetailMarshalled = true

    return json.Marshal(&struct {
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        Status string `json:"status"`
        
        Charges []Billingcharge `json:"charges"`
        
        Wallets []Billingwallet `json:"wallets"`
        *Alias
    }{

        


        


        


        


        
        Charges: []Billingcharge{{}},
        


        
        Wallets: []Billingwallet{{}},
        

        Alias: (*Alias)(u),
    })
}

