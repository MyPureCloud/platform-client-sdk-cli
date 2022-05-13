package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresponseresultexternalcontactentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresponseresultexternalcontactentityDud struct { 
    


    


    


    

}

// Bulkresponseresultexternalcontactentity
type Bulkresponseresultexternalcontactentity struct { 
    // Id
    Id string `json:"id"`


    // Success
    Success bool `json:"success"`


    // Entity
    Entity Externalcontact `json:"entity"`


    // VarError
    VarError Bulkerrorentity `json:"error"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalcontactentity) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresponseresultexternalcontactentity) MarshalJSON() ([]byte, error) {
    type Alias Bulkresponseresultexternalcontactentity

    if BulkresponseresultexternalcontactentityMarshalled {
        return []byte("{}"), nil
    }
    BulkresponseresultexternalcontactentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Success bool `json:"success"`
        
        Entity Externalcontact `json:"entity"`
        
        VarError Bulkerrorentity `json:"error"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

