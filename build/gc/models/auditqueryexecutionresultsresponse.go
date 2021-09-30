package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditqueryexecutionresultsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditqueryexecutionresultsresponseDud struct { 
    


    


    


    

}

// Auditqueryexecutionresultsresponse
type Auditqueryexecutionresultsresponse struct { 
    // Id - Id of the audit query execution request.
    Id string `json:"id"`


    // PageSize - Number of results in a page.
    PageSize int `json:"pageSize"`


    // Cursor - Optional cursor to indicate where to resume the results.
    Cursor string `json:"cursor"`


    // Entities - List of audit messages.
    Entities []Auditlogmessage `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Auditqueryexecutionresultsresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Entities = []Auditlogmessage{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditqueryexecutionresultsresponse) MarshalJSON() ([]byte, error) {
    type Alias Auditqueryexecutionresultsresponse

    if AuditqueryexecutionresultsresponseMarshalled {
        return []byte("{}"), nil
    }
    AuditqueryexecutionresultsresponseMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        PageSize int `json:"pageSize"`
        
        Cursor string `json:"cursor"`
        
        Entities []Auditlogmessage `json:"entities"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Entities: []Auditlogmessage{{}},
        

        
        Alias: (*Alias)(u),
    })
}

