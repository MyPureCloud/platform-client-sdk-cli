package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PendingandapprovedopportunityenrollmentcountsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PendingandapprovedopportunityenrollmentcountsDud struct { 
    


    

}

// Pendingandapprovedopportunityenrollmentcounts
type Pendingandapprovedopportunityenrollmentcounts struct { 
    // Pending - The number of pending enrollments
    Pending int `json:"pending"`


    // Approved - The number of approved enrollments
    Approved int `json:"approved"`

}

// String returns a JSON representation of the model
func (o *Pendingandapprovedopportunityenrollmentcounts) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pendingandapprovedopportunityenrollmentcounts) MarshalJSON() ([]byte, error) {
    type Alias Pendingandapprovedopportunityenrollmentcounts

    if PendingandapprovedopportunityenrollmentcountsMarshalled {
        return []byte("{}"), nil
    }
    PendingandapprovedopportunityenrollmentcountsMarshalled = true

    return json.Marshal(&struct {
        
        Pending int `json:"pending"`
        
        Approved int `json:"approved"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

