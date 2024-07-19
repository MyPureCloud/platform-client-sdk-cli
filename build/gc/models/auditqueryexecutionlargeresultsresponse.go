package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditqueryexecutionlargeresultsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditqueryexecutionlargeresultsresponseDud struct { 
    


    


    


    

}

// Auditqueryexecutionlargeresultsresponse
type Auditqueryexecutionlargeresultsresponse struct { 
    // Id - Id of the audit query execution request.
    Id string `json:"id"`


    // PageSize - Number of results in a page.
    PageSize int `json:"pageSize"`


    // Cursor - Optional cursor to indicate where to resume the results.
    Cursor string `json:"cursor"`


    // DownloadUrl - The presigned url which can be used to download the results.
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Auditqueryexecutionlargeresultsresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditqueryexecutionlargeresultsresponse) MarshalJSON() ([]byte, error) {
    type Alias Auditqueryexecutionlargeresultsresponse

    if AuditqueryexecutionlargeresultsresponseMarshalled {
        return []byte("{}"), nil
    }
    AuditqueryexecutionlargeresultsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        PageSize int `json:"pageSize"`
        
        Cursor string `json:"cursor"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

