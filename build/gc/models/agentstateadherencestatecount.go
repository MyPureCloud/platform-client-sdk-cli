package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentstateadherencestatecountMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentstateadherencestatecountDud struct { 
    


    

}

// Agentstateadherencestatecount
type Agentstateadherencestatecount struct { 
    // AdherenceState - Adherence state
    AdherenceState string `json:"adherenceState"`


    // Count - Count of users with this adherence state
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Agentstateadherencestatecount) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentstateadherencestatecount) MarshalJSON() ([]byte, error) {
    type Alias Agentstateadherencestatecount

    if AgentstateadherencestatecountMarshalled {
        return []byte("{}"), nil
    }
    AgentstateadherencestatecountMarshalled = true

    return json.Marshal(&struct {
        
        AdherenceState string `json:"adherenceState"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

