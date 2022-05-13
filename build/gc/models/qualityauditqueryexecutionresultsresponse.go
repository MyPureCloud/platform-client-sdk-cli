package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QualityauditqueryexecutionresultsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QualityauditqueryexecutionresultsresponseDud struct { 
    


    


    


    

}

// Qualityauditqueryexecutionresultsresponse
type Qualityauditqueryexecutionresultsresponse struct { 
    // Id - Id of the audit query execution request.
    Id string `json:"id"`


    // PageSize - Number of results in a page.
    PageSize int `json:"pageSize"`


    // Cursor - Optional cursor to indicate where to resume the results.
    Cursor string `json:"cursor"`


    // Entities - List of audit messages.
    Entities []Qualityauditlogmessage `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Qualityauditqueryexecutionresultsresponse) String() string {
    
    
    
     o.Entities = []Qualityauditlogmessage{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Qualityauditqueryexecutionresultsresponse) MarshalJSON() ([]byte, error) {
    type Alias Qualityauditqueryexecutionresultsresponse

    if QualityauditqueryexecutionresultsresponseMarshalled {
        return []byte("{}"), nil
    }
    QualityauditqueryexecutionresultsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        PageSize int `json:"pageSize"`
        
        Cursor string `json:"cursor"`
        
        Entities []Qualityauditlogmessage `json:"entities"`
        *Alias
    }{

        


        


        


        
        Entities: []Qualityauditlogmessage{{}},
        

        Alias: (*Alias)(u),
    })
}

