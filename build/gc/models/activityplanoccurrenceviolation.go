package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanoccurrenceviolationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanoccurrenceviolationDud struct { 
    


    

}

// Activityplanoccurrenceviolation
type Activityplanoccurrenceviolation struct { 
    // VarType - The type of violation for this occurrence
    VarType string `json:"type"`


    // Details - The detail of the violation
    Details Activityplanviolationdetails `json:"details"`

}

// String returns a JSON representation of the model
func (o *Activityplanoccurrenceviolation) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanoccurrenceviolation) MarshalJSON() ([]byte, error) {
    type Alias Activityplanoccurrenceviolation

    if ActivityplanoccurrenceviolationMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanoccurrenceviolationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Details Activityplanviolationdetails `json:"details"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

