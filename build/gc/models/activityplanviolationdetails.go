package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanviolationdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanviolationdetailsDud struct { 
    


    

}

// Activityplanviolationdetails
type Activityplanviolationdetails struct { 
    // Expected - The constraint value
    Expected string `json:"expected"`


    // Actual - The actual value
    Actual string `json:"actual"`

}

// String returns a JSON representation of the model
func (o *Activityplanviolationdetails) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanviolationdetails) MarshalJSON() ([]byte, error) {
    type Alias Activityplanviolationdetails

    if ActivityplanviolationdetailsMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanviolationdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Expected string `json:"expected"`
        
        Actual string `json:"actual"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

