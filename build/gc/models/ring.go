package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RingDud struct { 
    


    

}

// Ring
type Ring struct { 
    // ExpansionCriteria - The conditions that will trigger conversations to move to the next bullseye ring.
    ExpansionCriteria []Expansioncriterium `json:"expansionCriteria"`


    // Actions - The actions that will be performed just before moving conversations to the next bullseye ring.
    Actions Actions `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Ring) String() string {
     o.ExpansionCriteria = []Expansioncriterium{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ring) MarshalJSON() ([]byte, error) {
    type Alias Ring

    if RingMarshalled {
        return []byte("{}"), nil
    }
    RingMarshalled = true

    return json.Marshal(&struct {
        
        ExpansionCriteria []Expansioncriterium `json:"expansionCriteria"`
        
        Actions Actions `json:"actions"`
        *Alias
    }{

        
        ExpansionCriteria: []Expansioncriterium{{}},
        


        

        Alias: (*Alias)(u),
    })
}

