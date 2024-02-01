package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OperationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OperationresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Operationresponse
type Operationresponse struct { 
    


    // Status - Status of the operation.
    Status string `json:"status"`


    // VarType - Type of the operation.
    VarType string `json:"type"`


    // CreatedBy - The user who created the operation.
    CreatedBy Userreference `json:"createdBy"`


    // DateCreated - Operation creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Operation last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Operationresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Operationresponse) MarshalJSON() ([]byte, error) {
    type Alias Operationresponse

    if OperationresponseMarshalled {
        return []byte("{}"), nil
    }
    OperationresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        VarType string `json:"type"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

