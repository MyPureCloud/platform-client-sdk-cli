package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportrunentryentitydomainlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportrunentryentitydomainlistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Reportrunentryentitydomainlisting
type Reportrunentryentitydomainlisting struct { 
    // Entities
    Entities []Reportrunentry `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Reportrunentryentitydomainlisting) String() string {
     o.Entities = []Reportrunentry{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportrunentryentitydomainlisting) MarshalJSON() ([]byte, error) {
    type Alias Reportrunentryentitydomainlisting

    if ReportrunentryentitydomainlistingMarshalled {
        return []byte("{}"), nil
    }
    ReportrunentryentitydomainlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Reportrunentry `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        PreviousUri string `json:"previousUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Reportrunentry{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

