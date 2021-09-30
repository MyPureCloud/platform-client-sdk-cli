package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresponseresultexternalorganizationentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresponseresultexternalorganizationentityDud struct { 
    


    


    


    

}

// Bulkresponseresultexternalorganizationentity
type Bulkresponseresultexternalorganizationentity struct { 
    // Id
    Id string `json:"id"`


    // Success
    Success bool `json:"success"`


    // Entity
    Entity Externalorganization `json:"entity"`


    // VarError
    VarError Bulkerrorentity `json:"error"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalorganizationentity) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresponseresultexternalorganizationentity) MarshalJSON() ([]byte, error) {
    type Alias Bulkresponseresultexternalorganizationentity

    if BulkresponseresultexternalorganizationentityMarshalled {
        return []byte("{}"), nil
    }
    BulkresponseresultexternalorganizationentityMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Success bool `json:"success"`
        
        Entity Externalorganization `json:"entity"`
        
        VarError Bulkerrorentity `json:"error"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

