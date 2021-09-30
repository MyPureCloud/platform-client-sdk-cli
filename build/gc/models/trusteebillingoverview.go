package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrusteebillingoverviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrusteebillingoverviewDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Trusteebillingoverview
type Trusteebillingoverview struct { 
    


    // Name
    Name string `json:"name"`


    // Organization - Organization
    Organization Namedentity `json:"organization"`


    // Currency - The currency type.
    Currency string `json:"currency"`


    // EnabledProducts - The charge short names for products enabled during the specified period.
    EnabledProducts []string `json:"enabledProducts"`


    // SubscriptionType - The subscription type.
    SubscriptionType string `json:"subscriptionType"`


    // RampPeriodStartDate - Date-time the ramp period starts. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RampPeriodStartDate time.Time `json:"rampPeriodStartDate"`


    // RampPeriodEndDate - Date-time the ramp period ends. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    RampPeriodEndDate time.Time `json:"rampPeriodEndDate"`


    // BillingPeriodStartDate - Date-time the billing period started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    BillingPeriodStartDate time.Time `json:"billingPeriodStartDate"`


    // BillingPeriodEndDate - Date-time the billing period ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    BillingPeriodEndDate time.Time `json:"billingPeriodEndDate"`


    // Usages - Usages for the specified period.
    Usages []Subscriptionoverviewusage `json:"usages"`


    // ContractAmendmentDate - Date-time the contract was last amended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ContractAmendmentDate time.Time `json:"contractAmendmentDate"`


    // ContractEffectiveDate - Date-time the contract became effective. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ContractEffectiveDate time.Time `json:"contractEffectiveDate"`


    // ContractEndDate - Date-time the contract ends. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ContractEndDate time.Time `json:"contractEndDate"`


    // MinimumMonthlyAmount - Minimum amount that will be charged for the month
    MinimumMonthlyAmount string `json:"minimumMonthlyAmount"`


    // InRampPeriod
    InRampPeriod bool `json:"inRampPeriod"`


    

}

// String returns a JSON representation of the model
func (o *Trusteebillingoverview) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.EnabledProducts = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Usages = []Subscriptionoverviewusage{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trusteebillingoverview) MarshalJSON() ([]byte, error) {
    type Alias Trusteebillingoverview

    if TrusteebillingoverviewMarshalled {
        return []byte("{}"), nil
    }
    TrusteebillingoverviewMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Organization Namedentity `json:"organization"`
        
        Currency string `json:"currency"`
        
        EnabledProducts []string `json:"enabledProducts"`
        
        SubscriptionType string `json:"subscriptionType"`
        
        RampPeriodStartDate time.Time `json:"rampPeriodStartDate"`
        
        RampPeriodEndDate time.Time `json:"rampPeriodEndDate"`
        
        BillingPeriodStartDate time.Time `json:"billingPeriodStartDate"`
        
        BillingPeriodEndDate time.Time `json:"billingPeriodEndDate"`
        
        Usages []Subscriptionoverviewusage `json:"usages"`
        
        ContractAmendmentDate time.Time `json:"contractAmendmentDate"`
        
        ContractEffectiveDate time.Time `json:"contractEffectiveDate"`
        
        ContractEndDate time.Time `json:"contractEndDate"`
        
        MinimumMonthlyAmount string `json:"minimumMonthlyAmount"`
        
        InRampPeriod bool `json:"inRampPeriod"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        EnabledProducts: []string{""},
        

        

        

        

        

        

        

        

        

        

        

        

        
        Usages: []Subscriptionoverviewusage{{}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

