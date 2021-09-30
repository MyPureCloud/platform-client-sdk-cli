package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingusageresourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingusageresourceDud struct { 
    


    

}

// Billingusageresource
type Billingusageresource struct { 
    // Name - Identifies the resource (e.g. license user, device).
    Name string `json:"name"`


    // Date - The date that the usage was first observed by the billing subsystem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Date time.Time `json:"date"`

}

// String returns a JSON representation of the model
func (o *Billingusageresource) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingusageresource) MarshalJSON() ([]byte, error) {
    type Alias Billingusageresource

    if BillingusageresourceMarshalled {
        return []byte("{}"), nil
    }
    BillingusageresourceMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Date time.Time `json:"date"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

