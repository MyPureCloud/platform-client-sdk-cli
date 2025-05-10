package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingwalletMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingwalletDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    


    


    


    


    


    


    


    


    


    


    


    

}

// Billingwallet
type Billingwallet struct { 
    


    


    // Organizations - A Genesys Cloud Organization and it's related plans.
    Organizations []Namedentity `json:"organizations"`


    // Product - Represents the details of a product.
    Product Billingproduct `json:"product"`


    // StartingBalance - The initial balance in the wallet.
    StartingBalance float32 `json:"startingBalance"`


    // EndingBalance - The final balance in the wallet after transactions.
    EndingBalance float32 `json:"endingBalance"`


    // BalanceIncrease - Total amount added to the wallet.
    BalanceIncrease float32 `json:"balanceIncrease"`


    // BalanceDecrease - The amount removed from the wallet due to a contract change.
    BalanceDecrease float32 `json:"balanceDecrease"`


    // BalanceConsumption - Total consumption deducted from the wallet.
    BalanceConsumption float32 `json:"balanceConsumption"`


    // BalanceOverage - The amount exceeding a predefined balance threshold.
    BalanceOverage float32 `json:"balanceOverage"`


    // BalanceOverageRate - The rate charged for an overage..
    BalanceOverageRate float32 `json:"balanceOverageRate"`


    // BalanceOverageCharge - The amount to be charged.
    BalanceOverageCharge float32 `json:"balanceOverageCharge"`


    // BalanceOverageCurrency - The currency in which the overage charge is invoiced.
    BalanceOverageCurrency string `json:"balanceOverageCurrency"`


    // UnitOfMeasure - The unit of measure for the wallet.
    UnitOfMeasure string `json:"unitOfMeasure"`

}

// String returns a JSON representation of the model
func (o *Billingwallet) String() string {
     o.Organizations = []Namedentity{{}} 
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingwallet) MarshalJSON() ([]byte, error) {
    type Alias Billingwallet

    if BillingwalletMarshalled {
        return []byte("{}"), nil
    }
    BillingwalletMarshalled = true

    return json.Marshal(&struct {
        
        Organizations []Namedentity `json:"organizations"`
        
        Product Billingproduct `json:"product"`
        
        StartingBalance float32 `json:"startingBalance"`
        
        EndingBalance float32 `json:"endingBalance"`
        
        BalanceIncrease float32 `json:"balanceIncrease"`
        
        BalanceDecrease float32 `json:"balanceDecrease"`
        
        BalanceConsumption float32 `json:"balanceConsumption"`
        
        BalanceOverage float32 `json:"balanceOverage"`
        
        BalanceOverageRate float32 `json:"balanceOverageRate"`
        
        BalanceOverageCharge float32 `json:"balanceOverageCharge"`
        
        BalanceOverageCurrency string `json:"balanceOverageCurrency"`
        
        UnitOfMeasure string `json:"unitOfMeasure"`
        *Alias
    }{

        


        


        
        Organizations: []Namedentity{{}},
        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

