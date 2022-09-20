package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TesttargetoperationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TesttargetoperationDud struct { 
    


    


    


    

}

// Testtargetoperation - Information about the Trigger test mode target validation step
type Testtargetoperation struct { 
    // Name - The name of the processing step
    Name string `json:"name"`


    // Step - The number of the processing step
    Step int `json:"step"`


    // Matches - Whether or not the operation matches expectations
    Matches bool `json:"matches"`


    // Details - Details about why the operation did or did not succeed
    Details []string `json:"details"`

}

// String returns a JSON representation of the model
func (o *Testtargetoperation) String() string {
    
    
    
     o.Details = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testtargetoperation) MarshalJSON() ([]byte, error) {
    type Alias Testtargetoperation

    if TesttargetoperationMarshalled {
        return []byte("{}"), nil
    }
    TesttargetoperationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Step int `json:"step"`
        
        Matches bool `json:"matches"`
        
        Details []string `json:"details"`
        *Alias
    }{

        


        


        


        
        Details: []string{""},
        

        Alias: (*Alias)(u),
    })
}

