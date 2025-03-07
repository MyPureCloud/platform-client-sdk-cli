package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresponseresultvoidentitybulkentityerrorentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresponseresultvoidentitybulkentityerrorentityDud struct { 
    


    


    


    

}

// Bulkresponseresultvoidentitybulkentityerrorentity
type Bulkresponseresultvoidentitybulkentityerrorentity struct { 
    // Id - The id associated with this operation. For Bulk Enrich, this id is specified in the request; for all other Bulk endpoints, this id is the id of the affected entity.
    Id string `json:"id"`


    // Success - Whether the requested operation completed successfully.
    Success bool `json:"success"`


    // Entity - The entity which was affected by this Bulk operation. Only returned on success.
    Entity interface{} `json:"entity"`


    // VarError - An error describing why this Bulk operation failed. Only returned on failure.
    VarError Bulkentityerrorentity `json:"error"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultvoidentitybulkentityerrorentity) String() string {
    
    
     o.Entity = Interface{} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresponseresultvoidentitybulkentityerrorentity) MarshalJSON() ([]byte, error) {
    type Alias Bulkresponseresultvoidentitybulkentityerrorentity

    if BulkresponseresultvoidentitybulkentityerrorentityMarshalled {
        return []byte("{}"), nil
    }
    BulkresponseresultvoidentitybulkentityerrorentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Success bool `json:"success"`
        
        Entity interface{} `json:"entity"`
        
        VarError Bulkentityerrorentity `json:"error"`
        *Alias
    }{

        


        


        
        Entity: Interface{},
        


        

        Alias: (*Alias)(u),
    })
}

