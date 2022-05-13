package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresponseresultvoidentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresponseresultvoidentityDud struct { 
    


    


    


    

}

// Bulkresponseresultvoidentity
type Bulkresponseresultvoidentity struct { 
    // Id
    Id string `json:"id"`


    // Success
    Success bool `json:"success"`


    // Entity
    Entity interface{} `json:"entity"`


    // VarError
    VarError Bulkerrorentity `json:"error"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultvoidentity) String() string {
    
    
     o.Entity = Interface{} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresponseresultvoidentity) MarshalJSON() ([]byte, error) {
    type Alias Bulkresponseresultvoidentity

    if BulkresponseresultvoidentityMarshalled {
        return []byte("{}"), nil
    }
    BulkresponseresultvoidentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Success bool `json:"success"`
        
        Entity interface{} `json:"entity"`
        
        VarError Bulkerrorentity `json:"error"`
        *Alias
    }{

        


        


        
        Entity: Interface{},
        


        

        Alias: (*Alias)(u),
    })
}

