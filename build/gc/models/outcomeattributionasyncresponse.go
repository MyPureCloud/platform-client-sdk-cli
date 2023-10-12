package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeattributionasyncresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeattributionasyncresponseDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Outcomeattributionasyncresponse
type Outcomeattributionasyncresponse struct { 
    


    // PercentFailedThreshold - Optional percent failed threshold for validation errors; if reached will halt the job. Default is 5 percent, allowed values 0 to 100.
    PercentFailedThreshold int `json:"percentFailedThreshold"`


    

}

// String returns a JSON representation of the model
func (o *Outcomeattributionasyncresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeattributionasyncresponse) MarshalJSON() ([]byte, error) {
    type Alias Outcomeattributionasyncresponse

    if OutcomeattributionasyncresponseMarshalled {
        return []byte("{}"), nil
    }
    OutcomeattributionasyncresponseMarshalled = true

    return json.Marshal(&struct {
        
        PercentFailedThreshold int `json:"percentFailedThreshold"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

