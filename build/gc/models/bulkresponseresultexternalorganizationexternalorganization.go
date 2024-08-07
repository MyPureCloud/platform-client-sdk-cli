package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresponseresultexternalorganizationexternalorganizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresponseresultexternalorganizationexternalorganizationDud struct { 
    


    


    


    

}

// Bulkresponseresultexternalorganizationexternalorganization
type Bulkresponseresultexternalorganizationexternalorganization struct { 
    // Id
    Id string `json:"id"`


    // Success
    Success bool `json:"success"`


    // Entity
    Entity Externalorganization `json:"entity"`


    // VarError
    VarError Bulkerrorexternalorganization `json:"error"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalorganizationexternalorganization) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresponseresultexternalorganizationexternalorganization) MarshalJSON() ([]byte, error) {
    type Alias Bulkresponseresultexternalorganizationexternalorganization

    if BulkresponseresultexternalorganizationexternalorganizationMarshalled {
        return []byte("{}"), nil
    }
    BulkresponseresultexternalorganizationexternalorganizationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Success bool `json:"success"`
        
        Entity Externalorganization `json:"entity"`
        
        VarError Bulkerrorexternalorganization `json:"error"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

