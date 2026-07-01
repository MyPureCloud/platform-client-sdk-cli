package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenmonitoringuserdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenmonitoringuserdetailsDud struct { 
    


    

}

// Screenmonitoringuserdetails
type Screenmonitoringuserdetails struct { 
    // Count
    Count int `json:"count"`


    // TargetUser - The user being monitored
    TargetUser Addressableentityref `json:"targetUser"`

}

// String returns a JSON representation of the model
func (o *Screenmonitoringuserdetails) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenmonitoringuserdetails) MarshalJSON() ([]byte, error) {
    type Alias Screenmonitoringuserdetails

    if ScreenmonitoringuserdetailsMarshalled {
        return []byte("{}"), nil
    }
    ScreenmonitoringuserdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Count int `json:"count"`
        
        TargetUser Addressableentityref `json:"targetUser"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

