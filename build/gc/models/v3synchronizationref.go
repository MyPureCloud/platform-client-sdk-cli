package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationrefDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// V3synchronizationref
type V3synchronizationref struct { 
    // Id - Synchronization Id.
    Id string `json:"id"`


    // Source - Source.
    Source V3sourceref `json:"source"`


    

}

// String returns a JSON representation of the model
func (o *V3synchronizationref) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationref) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationref

    if V3synchronizationrefMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationrefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Source V3sourceref `json:"source"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

