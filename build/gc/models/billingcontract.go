package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingcontractMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingcontractDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Billingcontract
type Billingcontract struct { 
    


    // ExternalNumber - Unique external number.
    ExternalNumber string `json:"externalNumber"`


    // Status - The status of the contract.
    Status string `json:"status"`


    // CommercialModel - The type of commercial model.
    CommercialModel string `json:"commercialModel"`


    // GetpurchaseOrderNumbers - List of po numbers periods for this contract.
    GetpurchaseOrderNumbers []string `json:"getpurchaseOrderNumbers"`


    // BillToCustomer - The bill to customer.
    BillToCustomer Customer `json:"billToCustomer"`


    // SoldToCustomer - The sold to customer.
    SoldToCustomer Customer `json:"soldToCustomer"`


    // EndCustomer - The end customer.
    EndCustomer Customer `json:"endCustomer"`


    // DateStart - The start date of the contract. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The end date of the contract. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEnd time.Time `json:"dateEnd"`


    // DateRampStart - the date when the first revenue or quantity in a ramped deal begins. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateRampStart time.Time `json:"dateRampStart"`


    // DateRampEnd - the date when the first revenue or quantity in a ramped deal ends. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateRampEnd time.Time `json:"dateRampEnd"`


    // BillingPeriods - List of billing periods for the contract.
    BillingPeriods []Billingcontractperiod `json:"billingPeriods"`


    // Plans - The collection of prices for the related organizations.
    Plans []Billingplan `json:"plans"`

}

// String returns a JSON representation of the model
func (o *Billingcontract) String() string {
    
    
    
     o.GetpurchaseOrderNumbers = []string{""} 
    
    
    
    
    
    
    
     o.BillingPeriods = []Billingcontractperiod{{}} 
     o.Plans = []Billingplan{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingcontract) MarshalJSON() ([]byte, error) {
    type Alias Billingcontract

    if BillingcontractMarshalled {
        return []byte("{}"), nil
    }
    BillingcontractMarshalled = true

    return json.Marshal(&struct {
        
        ExternalNumber string `json:"externalNumber"`
        
        Status string `json:"status"`
        
        CommercialModel string `json:"commercialModel"`
        
        GetpurchaseOrderNumbers []string `json:"getpurchaseOrderNumbers"`
        
        BillToCustomer Customer `json:"billToCustomer"`
        
        SoldToCustomer Customer `json:"soldToCustomer"`
        
        EndCustomer Customer `json:"endCustomer"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        DateRampStart time.Time `json:"dateRampStart"`
        
        DateRampEnd time.Time `json:"dateRampEnd"`
        
        BillingPeriods []Billingcontractperiod `json:"billingPeriods"`
        
        Plans []Billingplan `json:"plans"`
        *Alias
    }{

        


        


        


        


        
        GetpurchaseOrderNumbers: []string{""},
        


        


        


        


        


        


        


        


        
        BillingPeriods: []Billingcontractperiod{{}},
        


        
        Plans: []Billingplan{{}},
        

        Alias: (*Alias)(u),
    })
}

