package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingcontractperiodMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingcontractperiodDud struct { 
    Id string `json:"id"`


    


    


    

}

// Billingcontractperiod
type Billingcontractperiod struct { 
    


    // DateStart - The date when the Billing Period starts. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The date when the Billing Period starts. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEnd time.Time `json:"dateEnd"`


    // Status - The type of address.
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Billingcontractperiod) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingcontractperiod) MarshalJSON() ([]byte, error) {
    type Alias Billingcontractperiod

    if BillingcontractperiodMarshalled {
        return []byte("{}"), nil
    }
    BillingcontractperiodMarshalled = true

    return json.Marshal(&struct {
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

