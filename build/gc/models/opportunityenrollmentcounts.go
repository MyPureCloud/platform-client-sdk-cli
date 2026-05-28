package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpportunityenrollmentcountsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpportunityenrollmentcountsDud struct { 
    


    


    


    

}

// Opportunityenrollmentcounts
type Opportunityenrollmentcounts struct { 
    // Pending - The number of pending enrollments for this opportunity
    Pending int `json:"pending"`


    // Approved - The number of approved enrollments for this opportunity
    Approved int `json:"approved"`


    // Denied - The number of denied enrollments for this opportunity
    Denied int `json:"denied"`


    // Withdrawn - The number of withdrawn enrollments for this opportunity
    Withdrawn int `json:"withdrawn"`

}

// String returns a JSON representation of the model
func (o *Opportunityenrollmentcounts) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opportunityenrollmentcounts) MarshalJSON() ([]byte, error) {
    type Alias Opportunityenrollmentcounts

    if OpportunityenrollmentcountsMarshalled {
        return []byte("{}"), nil
    }
    OpportunityenrollmentcountsMarshalled = true

    return json.Marshal(&struct {
        
        Pending int `json:"pending"`
        
        Approved int `json:"approved"`
        
        Denied int `json:"denied"`
        
        Withdrawn int `json:"withdrawn"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

