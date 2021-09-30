package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditrealtimequeryresultsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditrealtimequeryresultsresponseDud struct { 
    


    


    


    


    

}

// Auditrealtimequeryresultsresponse
type Auditrealtimequeryresultsresponse struct { 
    // Entities
    Entities []Auditlogmessage `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Auditrealtimequeryresultsresponse) String() string {
    
    
     o.Entities = []Auditlogmessage{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditrealtimequeryresultsresponse) MarshalJSON() ([]byte, error) {
    type Alias Auditrealtimequeryresultsresponse

    if AuditrealtimequeryresultsresponseMarshalled {
        return []byte("{}"), nil
    }
    AuditrealtimequeryresultsresponseMarshalled = true

    return json.Marshal(&struct { 
        Entities []Auditlogmessage `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        *Alias
    }{
        

        
        Entities: []Auditlogmessage{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

