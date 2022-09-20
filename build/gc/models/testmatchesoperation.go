package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TestmatchesoperationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TestmatchesoperationDud struct { 
    


    


    


    

}

// Testmatchesoperation - Information about the Trigger test mode processing step
type Testmatchesoperation struct { 
    // Name - The name of the processing step
    Name string `json:"name"`


    // Step - The number of the processing step
    Step int `json:"step"`


    // Matches - Whether or not the operation matches expectations
    Matches bool `json:"matches"`


    // Details - Details about why the operation did or did not succeed
    Details []Matchcriteriatestresult `json:"details"`

}

// String returns a JSON representation of the model
func (o *Testmatchesoperation) String() string {
    
    
    
     o.Details = []Matchcriteriatestresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testmatchesoperation) MarshalJSON() ([]byte, error) {
    type Alias Testmatchesoperation

    if TestmatchesoperationMarshalled {
        return []byte("{}"), nil
    }
    TestmatchesoperationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Step int `json:"step"`
        
        Matches bool `json:"matches"`
        
        Details []Matchcriteriatestresult `json:"details"`
        *Alias
    }{

        


        


        


        
        Details: []Matchcriteriatestresult{{}},
        

        Alias: (*Alias)(u),
    })
}

