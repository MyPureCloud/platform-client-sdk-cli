package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NluconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NluconfigDud struct { 
    


    

}

// Nluconfig
type Nluconfig struct { 
    // Domain - NLU domain.
    Domain Copilotnludomain `json:"domain"`


    // IntentConfidenceThreshold - Minimum confidence value of accepting NLU intents, must be greater than 0 and less than 1.
    IntentConfidenceThreshold float32 `json:"intentConfidenceThreshold"`

}

// String returns a JSON representation of the model
func (o *Nluconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nluconfig) MarshalJSON() ([]byte, error) {
    type Alias Nluconfig

    if NluconfigMarshalled {
        return []byte("{}"), nil
    }
    NluconfigMarshalled = true

    return json.Marshal(&struct {
        
        Domain Copilotnludomain `json:"domain"`
        
        IntentConfidenceThreshold float32 `json:"intentConfidenceThreshold"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

