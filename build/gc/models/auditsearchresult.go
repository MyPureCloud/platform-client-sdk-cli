package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditsearchresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditsearchresultDud struct { 
    


    


    


    


    


    

}

// Auditsearchresult
type Auditsearchresult struct { 
    // PageNumber - Which page was returned.
    PageNumber int `json:"pageNumber"`


    // PageSize - The number of results in a page.
    PageSize int `json:"pageSize"`


    // Total - The total number of results.
    Total int `json:"total"`


    // PageCount - The number of pages of results.
    PageCount int `json:"pageCount"`


    // FacetInfo
    FacetInfo []Facetinfo `json:"facetInfo"`


    // AuditMessages
    AuditMessages []Auditmessage `json:"auditMessages"`

}

// String returns a JSON representation of the model
func (o *Auditsearchresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.FacetInfo = []Facetinfo{{}} 
    
    
    
     o.AuditMessages = []Auditmessage{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditsearchresult) MarshalJSON() ([]byte, error) {
    type Alias Auditsearchresult

    if AuditsearchresultMarshalled {
        return []byte("{}"), nil
    }
    AuditsearchresultMarshalled = true

    return json.Marshal(&struct { 
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        FacetInfo []Facetinfo `json:"facetInfo"`
        
        AuditMessages []Auditmessage `json:"auditMessages"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        FacetInfo: []Facetinfo{{}},
        

        

        
        AuditMessages: []Auditmessage{{}},
        

        
        Alias: (*Alias)(u),
    })
}

