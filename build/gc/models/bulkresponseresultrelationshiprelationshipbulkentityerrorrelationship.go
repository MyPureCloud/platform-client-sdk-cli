package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkresponseresultrelationshiprelationshipbulkentityerrorrelationshipMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkresponseresultrelationshiprelationshipbulkentityerrorrelationshipDud struct { 
    


    


    


    


    

}

// Bulkresponseresultrelationshiprelationshipbulkentityerrorrelationship
type Bulkresponseresultrelationshiprelationshipbulkentityerrorrelationship struct { 
    // Id - The id associated with this operation. For Bulk Enrich, this id is specified in the request; for all other Bulk endpoints, this id is the id of the affected entity.
    Id string `json:"id"`


    // Success - Whether the requested operation completed successfully.
    Success bool `json:"success"`


    // Entity - The entity which was affected by this Bulk operation. Only returned on success.
    Entity Relationship `json:"entity"`


    // VarError - An error describing why this Bulk operation failed. Only returned on failure.
    VarError Bulkentityerrorrelationship `json:"error"`


    // Status - Status Code for the requested operation.
    Status int `json:"status"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultrelationshiprelationshipbulkentityerrorrelationship) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkresponseresultrelationshiprelationshipbulkentityerrorrelationship) MarshalJSON() ([]byte, error) {
    type Alias Bulkresponseresultrelationshiprelationshipbulkentityerrorrelationship

    if BulkresponseresultrelationshiprelationshipbulkentityerrorrelationshipMarshalled {
        return []byte("{}"), nil
    }
    BulkresponseresultrelationshiprelationshipbulkentityerrorrelationshipMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Success bool `json:"success"`
        
        Entity Relationship `json:"entity"`
        
        VarError Bulkentityerrorrelationship `json:"error"`
        
        Status int `json:"status"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

