package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresponseresultrelationshipentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresponseresultrelationshipentityDud struct { 
    


    


    


    

}

// Bulkresponseresultrelationshipentity
type Bulkresponseresultrelationshipentity struct { 
    // Id
    Id string `json:"id"`


    // Success
    Success bool `json:"success"`


    // Entity
    Entity Relationship `json:"entity"`


    // VarError
    VarError Bulkerrorentity `json:"error"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultrelationshipentity) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresponseresultrelationshipentity) MarshalJSON() ([]byte, error) {
    type Alias Bulkresponseresultrelationshipentity

    if BulkresponseresultrelationshipentityMarshalled {
        return []byte("{}"), nil
    }
    BulkresponseresultrelationshipentityMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Success bool `json:"success"`
        
        Entity Relationship `json:"entity"`
        
        VarError Bulkerrorentity `json:"error"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

