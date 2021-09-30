package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingusageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingusageDud struct { 
    


    


    

}

// Billingusage
type Billingusage struct { 
    // Name - Identifies the billable usage.
    Name string `json:"name"`


    // TotalUsage - The total amount of usage, expressed as a decimal number in string format.
    TotalUsage string `json:"totalUsage"`


    // Resources - The resources for which usage was observed (e.g. license users, devices).
    Resources []Billingusageresource `json:"resources"`

}

// String returns a JSON representation of the model
func (o *Billingusage) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Resources = []Billingusageresource{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingusage) MarshalJSON() ([]byte, error) {
    type Alias Billingusage

    if BillingusageMarshalled {
        return []byte("{}"), nil
    }
    BillingusageMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        TotalUsage string `json:"totalUsage"`
        
        Resources []Billingusageresource `json:"resources"`
        
        *Alias
    }{
        

        

        

        

        

        
        Resources: []Billingusageresource{{}},
        

        
        Alias: (*Alias)(u),
    })
}

