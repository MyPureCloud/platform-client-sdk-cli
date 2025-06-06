package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkentityerrorexternalorganizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkentityerrorexternalorganizationDud struct { 
    


    


    


    


    


    

}

// Bulkentityerrorexternalorganization
type Bulkentityerrorexternalorganization struct { 
    // Code - An error code for the specific error condition.
    Code string `json:"code"`


    // Message - A short error message.
    Message string `json:"message"`


    // Status - The HTTP Status Code for the error.
    Status int `json:"status"`


    // Retryable - Whether this particular error should be retried.
    Retryable bool `json:"retryable"`


    // Details - Additional error details for specific fields.
    Details []Bulkerrordetail `json:"details"`


    // Entity - The entity body specified in the Bulk request operation that caused this error.
    Entity Externalorganization `json:"entity"`

}

// String returns a JSON representation of the model
func (o *Bulkentityerrorexternalorganization) String() string {
    
    
    
    
     o.Details = []Bulkerrordetail{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkentityerrorexternalorganization) MarshalJSON() ([]byte, error) {
    type Alias Bulkentityerrorexternalorganization

    if BulkentityerrorexternalorganizationMarshalled {
        return []byte("{}"), nil
    }
    BulkentityerrorexternalorganizationMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        
        Status int `json:"status"`
        
        Retryable bool `json:"retryable"`
        
        Details []Bulkerrordetail `json:"details"`
        
        Entity Externalorganization `json:"entity"`
        *Alias
    }{

        


        


        


        


        
        Details: []Bulkerrordetail{{}},
        


        

        Alias: (*Alias)(u),
    })
}

