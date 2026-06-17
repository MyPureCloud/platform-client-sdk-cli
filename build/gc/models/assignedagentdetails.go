package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssignedagentdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssignedagentdetailsDud struct { 
    


    

}

// Assignedagentdetails
type Assignedagentdetails struct { 
    // Id - The ID of the agent
    Id string `json:"id"`


    // EndDate - The end date of this schedule set for the agent, relative to the business unit time zone in yyyy-MM-dd format. Null denotes an ongoing schedule set. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    EndDate time.Time `json:"endDate"`

}

// String returns a JSON representation of the model
func (o *Assignedagentdetails) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assignedagentdetails) MarshalJSON() ([]byte, error) {
    type Alias Assignedagentdetails

    if AssignedagentdetailsMarshalled {
        return []byte("{}"), nil
    }
    AssignedagentdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        EndDate time.Time `json:"endDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

