package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanninggroupreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanninggroupreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Planninggroupreference - Planning Group
type Planninggroupreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Planninggroupreference) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planninggroupreference) MarshalJSON() ([]byte, error) {
    type Alias Planninggroupreference

    if PlanninggroupreferenceMarshalled {
        return []byte("{}"), nil
    }
    PlanninggroupreferenceMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

