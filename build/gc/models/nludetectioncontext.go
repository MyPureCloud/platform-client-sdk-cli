package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludetectioncontextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludetectioncontextDud struct { 
    


    

}

// Nludetectioncontext
type Nludetectioncontext struct { 
    // Intent - Restrict detection to this intent.
    Intent Contextintent `json:"intent"`


    // Entity - Use this entity to restrict detection.
    Entity Contextentity `json:"entity"`

}

// String returns a JSON representation of the model
func (o *Nludetectioncontext) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludetectioncontext) MarshalJSON() ([]byte, error) {
    type Alias Nludetectioncontext

    if NludetectioncontextMarshalled {
        return []byte("{}"), nil
    }
    NludetectioncontextMarshalled = true

    return json.Marshal(&struct {
        
        Intent Contextintent `json:"intent"`
        
        Entity Contextentity `json:"entity"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

