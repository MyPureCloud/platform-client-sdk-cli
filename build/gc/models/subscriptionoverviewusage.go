package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SubscriptionoverviewusageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SubscriptionoverviewusageDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Subscriptionoverviewusage
type Subscriptionoverviewusage struct { 
    // Name - Product charge name
    Name string `json:"name"`


    // PartNumber - Product part number
    PartNumber string `json:"partNumber"`


    // Grouping - UI grouping key
    Grouping string `json:"grouping"`


    // UnitOfMeasureType - UI unit of measure
    UnitOfMeasureType string `json:"unitOfMeasureType"`


    // UsageQuantity - Usage count for specified period
    UsageQuantity string `json:"usageQuantity"`


    // OveragePrice - Price for usage / overage charge
    OveragePrice string `json:"overagePrice"`


    // PrepayQuantity - Items prepaid for specified period
    PrepayQuantity string `json:"prepayQuantity"`


    // PrepayPrice - Price for prepay charge
    PrepayPrice string `json:"prepayPrice"`


    // UsageNotes - Notes about the usage/charge item
    UsageNotes string `json:"usageNotes"`


    // IsCancellable - Indicates whether the item is cancellable
    IsCancellable bool `json:"isCancellable"`


    // BundleQuantity - Quantity multiplier for this charge
    BundleQuantity string `json:"bundleQuantity"`


    // IsThirdParty - A charge from a third party entity
    IsThirdParty bool `json:"isThirdParty"`

}

// String returns a JSON representation of the model
func (o *Subscriptionoverviewusage) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Subscriptionoverviewusage) MarshalJSON() ([]byte, error) {
    type Alias Subscriptionoverviewusage

    if SubscriptionoverviewusageMarshalled {
        return []byte("{}"), nil
    }
    SubscriptionoverviewusageMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        PartNumber string `json:"partNumber"`
        
        Grouping string `json:"grouping"`
        
        UnitOfMeasureType string `json:"unitOfMeasureType"`
        
        UsageQuantity string `json:"usageQuantity"`
        
        OveragePrice string `json:"overagePrice"`
        
        PrepayQuantity string `json:"prepayQuantity"`
        
        PrepayPrice string `json:"prepayPrice"`
        
        UsageNotes string `json:"usageNotes"`
        
        IsCancellable bool `json:"isCancellable"`
        
        BundleQuantity string `json:"bundleQuantity"`
        
        IsThirdParty bool `json:"isThirdParty"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

