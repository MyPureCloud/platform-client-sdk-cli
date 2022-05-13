package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkrecordingenabledcountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkrecordingenabledcountDud struct { 
    


    

}

// Trunkrecordingenabledcount
type Trunkrecordingenabledcount struct { 
    // EnabledCount - The amount of trunks that have recording enabled
    EnabledCount int `json:"enabledCount"`


    // DisabledCount - The amount of trunks that do not have recording enabled
    DisabledCount int `json:"disabledCount"`

}

// String returns a JSON representation of the model
func (o *Trunkrecordingenabledcount) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkrecordingenabledcount) MarshalJSON() ([]byte, error) {
    type Alias Trunkrecordingenabledcount

    if TrunkrecordingenabledcountMarshalled {
        return []byte("{}"), nil
    }
    TrunkrecordingenabledcountMarshalled = true

    return json.Marshal(&struct {
        
        EnabledCount int `json:"enabledCount"`
        
        DisabledCount int `json:"disabledCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

