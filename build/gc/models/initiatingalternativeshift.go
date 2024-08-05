package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InitiatingalternativeshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InitiatingalternativeshiftDud struct { 
    


    

}

// Initiatingalternativeshift
type Initiatingalternativeshift struct { 
    // Id - The ID of the shift from a user's schedule
    Id string `json:"id"`


    // StartDate - The start date for the shift in the user's schedule in ISO-8601 format. For example: YYYY-MM-DDThh:mm:ss.SSSZ
    StartDate time.Time `json:"startDate"`

}

// String returns a JSON representation of the model
func (o *Initiatingalternativeshift) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Initiatingalternativeshift) MarshalJSON() ([]byte, error) {
    type Alias Initiatingalternativeshift

    if InitiatingalternativeshiftMarshalled {
        return []byte("{}"), nil
    }
    InitiatingalternativeshiftMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        StartDate time.Time `json:"startDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

