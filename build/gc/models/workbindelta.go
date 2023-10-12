package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkbindeltaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkbindeltaDud struct { 
    


    


    


    

}

// Workbindelta
type Workbindelta struct { 
    // Name
    Name Workitemsattributechangestring `json:"name"`


    // Description
    Description Workitemsattributechangestring `json:"description"`


    // DateModified
    DateModified Workitemsattributechangeinstant `json:"dateModified"`


    // ModifiedBy
    ModifiedBy Workitemsattributechangestring `json:"modifiedBy"`

}

// String returns a JSON representation of the model
func (o *Workbindelta) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workbindelta) MarshalJSON() ([]byte, error) {
    type Alias Workbindelta

    if WorkbindeltaMarshalled {
        return []byte("{}"), nil
    }
    WorkbindeltaMarshalled = true

    return json.Marshal(&struct {
        
        Name Workitemsattributechangestring `json:"name"`
        
        Description Workitemsattributechangestring `json:"description"`
        
        DateModified Workitemsattributechangeinstant `json:"dateModified"`
        
        ModifiedBy Workitemsattributechangestring `json:"modifiedBy"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

