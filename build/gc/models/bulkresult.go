package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresultDud struct { 
    


    

}

// Bulkresult
type Bulkresult struct { 
    // VarError - Error details if the operation failed.
    VarError Bulkerror `json:"error"`


    // Entity - The result of the operation if it succeeded.
    Entity interface{} `json:"entity"`

}

// String returns a JSON representation of the model
func (o *Bulkresult) String() string {
    
     o.Entity = Interface{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresult) MarshalJSON() ([]byte, error) {
    type Alias Bulkresult

    if BulkresultMarshalled {
        return []byte("{}"), nil
    }
    BulkresultMarshalled = true

    return json.Marshal(&struct {
        
        VarError Bulkerror `json:"error"`
        
        Entity interface{} `json:"entity"`
        *Alias
    }{

        


        
        Entity: Interface{},
        

        Alias: (*Alias)(u),
    })
}

