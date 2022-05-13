package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresponseresultexternalcontactexternalcontactMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresponseresultexternalcontactexternalcontactDud struct { 
    


    


    


    

}

// Bulkresponseresultexternalcontactexternalcontact
type Bulkresponseresultexternalcontactexternalcontact struct { 
    // Id
    Id string `json:"id"`


    // Success
    Success bool `json:"success"`


    // Entity
    Entity Externalcontact `json:"entity"`


    // VarError
    VarError Bulkerrorexternalcontact `json:"error"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalcontactexternalcontact) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresponseresultexternalcontactexternalcontact) MarshalJSON() ([]byte, error) {
    type Alias Bulkresponseresultexternalcontactexternalcontact

    if BulkresponseresultexternalcontactexternalcontactMarshalled {
        return []byte("{}"), nil
    }
    BulkresponseresultexternalcontactexternalcontactMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Success bool `json:"success"`
        
        Entity Externalcontact `json:"entity"`
        
        VarError Bulkerrorexternalcontact `json:"error"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

