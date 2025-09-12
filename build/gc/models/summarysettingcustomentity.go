package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummarysettingcustomentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummarysettingcustomentityDud struct { 
    


    

}

// Summarysettingcustomentity
type Summarysettingcustomentity struct { 
    // Label - Label how the entity should be called.
    Label string `json:"label"`


    // Description - Describe the information the entity captures.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Summarysettingcustomentity) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summarysettingcustomentity) MarshalJSON() ([]byte, error) {
    type Alias Summarysettingcustomentity

    if SummarysettingcustomentityMarshalled {
        return []byte("{}"), nil
    }
    SummarysettingcustomentityMarshalled = true

    return json.Marshal(&struct {
        
        Label string `json:"label"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

