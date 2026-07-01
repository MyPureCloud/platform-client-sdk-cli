package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenmonitoringdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenmonitoringdetailsDud struct { 
    

}

// Screenmonitoringdetails
type Screenmonitoringdetails struct { 
    // Count - The total count of screen monitoring sessions in the organization
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Screenmonitoringdetails) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenmonitoringdetails) MarshalJSON() ([]byte, error) {
    type Alias Screenmonitoringdetails

    if ScreenmonitoringdetailsMarshalled {
        return []byte("{}"), nil
    }
    ScreenmonitoringdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Count int `json:"count"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

