package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BillingusagereportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BillingusagereportDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Billingusagereport
type Billingusagereport struct { 
    


    // Name
    Name string `json:"name"`


    // StartDate - The period start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // EndDate - The period end date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndDate time.Time `json:"endDate"`


    // Status - Generation status of report
    Status string `json:"status"`


    // Usages - The usages for the given period.
    Usages []Billingusage `json:"usages"`


    

}

// String returns a JSON representation of the model
func (o *Billingusagereport) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Usages = []Billingusage{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Billingusagereport) MarshalJSON() ([]byte, error) {
    type Alias Billingusagereport

    if BillingusagereportMarshalled {
        return []byte("{}"), nil
    }
    BillingusagereportMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        Status string `json:"status"`
        
        Usages []Billingusage `json:"usages"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        Usages: []Billingusage{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

