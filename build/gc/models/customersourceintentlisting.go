package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomersourceintentlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomersourceintentlistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Customersourceintentlisting
type Customersourceintentlisting struct { 
    // Entities
    Entities []Customersourceintent `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Customersourceintentlisting) String() string {
     o.Entities = []Customersourceintent{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customersourceintentlisting) MarshalJSON() ([]byte, error) {
    type Alias Customersourceintentlisting

    if CustomersourceintentlistingMarshalled {
        return []byte("{}"), nil
    }
    CustomersourceintentlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Customersourceintent `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Customersourceintent{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

