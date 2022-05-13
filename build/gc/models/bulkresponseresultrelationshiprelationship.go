package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresponseresultrelationshiprelationshipMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresponseresultrelationshiprelationshipDud struct { 
    


    


    


    

}

// Bulkresponseresultrelationshiprelationship
type Bulkresponseresultrelationshiprelationship struct { 
    // Id
    Id string `json:"id"`


    // Success
    Success bool `json:"success"`


    // Entity
    Entity Relationship `json:"entity"`


    // VarError
    VarError Bulkerrorrelationship `json:"error"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultrelationshiprelationship) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresponseresultrelationshiprelationship) MarshalJSON() ([]byte, error) {
    type Alias Bulkresponseresultrelationshiprelationship

    if BulkresponseresultrelationshiprelationshipMarshalled {
        return []byte("{}"), nil
    }
    BulkresponseresultrelationshiprelationshipMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Success bool `json:"success"`
        
        Entity Relationship `json:"entity"`
        
        VarError Bulkerrorrelationship `json:"error"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

