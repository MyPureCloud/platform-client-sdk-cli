package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainuserdivisionsresultlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainuserdivisionsresultlistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Domainuserdivisionsresultlisting
type Domainuserdivisionsresultlisting struct { 
    // Entities
    Entities []Domainuserdivisionsresult `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Domainuserdivisionsresultlisting) String() string {
     o.Entities = []Domainuserdivisionsresult{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainuserdivisionsresultlisting) MarshalJSON() ([]byte, error) {
    type Alias Domainuserdivisionsresultlisting

    if DomainuserdivisionsresultlistingMarshalled {
        return []byte("{}"), nil
    }
    DomainuserdivisionsresultlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Domainuserdivisionsresult `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Domainuserdivisionsresult{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

