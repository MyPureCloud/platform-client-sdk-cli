package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingcontractsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingcontractsummaryDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Billingcontractsummary
type Billingcontractsummary struct { 
    


    // ExternalNumber - Unique external number.
    ExternalNumber string `json:"externalNumber"`


    // Status - The status of the contract.
    Status string `json:"status"`


    // CommercialModel - The type of commercial model.
    CommercialModel string `json:"commercialModel"`


    // BillingFrequency - How the contract is billed. Values include Monthly Pre-pay, Annual Pre-pay and Monthly in Arrears.
    BillingFrequency string `json:"billingFrequency"`


    // PurchaseOrderNumbers - List of po numbers periods for this contract.
    PurchaseOrderNumbers []string `json:"purchaseOrderNumbers"`


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


    // Organizations - Organizations associated with the contract, deduplicated from plan associations.
    Organizations []Namedentity `json:"organizations"`

}

// String returns a JSON representation of the model
func (o *Billingcontractsummary) String() string {
    
    
    
    
     o.PurchaseOrderNumbers = []string{""} 
    
    
    
    
    
    
    
     o.BillingPeriods = []Billingcontractperiod{{}} 
     o.Organizations = []Namedentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingcontractsummary) MarshalJSON() ([]byte, error) {
    type Alias Billingcontractsummary

    if BillingcontractsummaryMarshalled {
        return []byte("{}"), nil
    }
    BillingcontractsummaryMarshalled = true

    return json.Marshal(&struct {
        
        ExternalNumber string `json:"externalNumber"`
        
        Status string `json:"status"`
        
        CommercialModel string `json:"commercialModel"`
        
        BillingFrequency string `json:"billingFrequency"`
        
        PurchaseOrderNumbers []string `json:"purchaseOrderNumbers"`
        
        BillToCustomer Customer `json:"billToCustomer"`
        
        SoldToCustomer Customer `json:"soldToCustomer"`
        
        EndCustomer Customer `json:"endCustomer"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        DateRampStart time.Time `json:"dateRampStart"`
        
        DateRampEnd time.Time `json:"dateRampEnd"`
        
        BillingPeriods []Billingcontractperiod `json:"billingPeriods"`
        
        Organizations []Namedentity `json:"organizations"`
        *Alias
    }{

        


        


        


        


        


        
        PurchaseOrderNumbers: []string{""},
        


        


        


        


        


        


        


        


        
        BillingPeriods: []Billingcontractperiod{{}},
        


        
        Organizations: []Namedentity{{}},
        

        Alias: (*Alias)(u),
    })
}

