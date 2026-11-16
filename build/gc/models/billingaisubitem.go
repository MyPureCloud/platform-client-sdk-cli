package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingaisubitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingaisubitemDud struct { 
    


    


    


    


    

}

// Billingaisubitem
type Billingaisubitem struct { 
    // ChargeName - Charge-level name of the member charge. Omitted when the source charge has no name.
    ChargeName string `json:"chargeName"`


    // ActualQuantity - Usage in the child product's original unit.
    ActualQuantity float32 `json:"actualQuantity"`


    // UnitOfMeasure - Original unit of measure for the child product usage.
    UnitOfMeasure string `json:"unitOfMeasure"`


    // ConvertedTokenCount - Child usage converted to parent-bundle tokens.
    ConvertedTokenCount float32 `json:"convertedTokenCount"`


    // TokenConversionRate - Bundle tokens per one original unit.
    TokenConversionRate float32 `json:"tokenConversionRate"`

}

// String returns a JSON representation of the model
func (o *Billingaisubitem) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingaisubitem) MarshalJSON() ([]byte, error) {
    type Alias Billingaisubitem

    if BillingaisubitemMarshalled {
        return []byte("{}"), nil
    }
    BillingaisubitemMarshalled = true

    return json.Marshal(&struct {
        
        ChargeName string `json:"chargeName"`
        
        ActualQuantity float32 `json:"actualQuantity"`
        
        UnitOfMeasure string `json:"unitOfMeasure"`
        
        ConvertedTokenCount float32 `json:"convertedTokenCount"`
        
        TokenConversionRate float32 `json:"tokenConversionRate"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

