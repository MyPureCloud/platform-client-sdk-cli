package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManagementunitlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManagementunitlistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Managementunitlisting
type Managementunitlisting struct { 
    // Entities
    Entities []Managementunit `json:"entities"`


    // PageSize - Deprecated, paging is not supported
    PageSize int `json:"pageSize"`


    // PageNumber - Deprecated, paging is not supported
    PageNumber int `json:"pageNumber"`


    // Total - Deprecated, paging is not supported
    Total int `json:"total"`


    // PageCount - Deprecated, paging is not supported
    PageCount int `json:"pageCount"`


    // PreviousUri - Deprecated, paging is not supported
    PreviousUri string `json:"previousUri"`


    // NextUri - Deprecated, paging is not supported
    NextUri string `json:"nextUri"`


    // LastUri - Deprecated, paging is not supported
    LastUri string `json:"lastUri"`


    // FirstUri - Deprecated, paging is not supported
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Managementunitlisting) String() string {
     o.Entities = []Managementunit{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Managementunitlisting) MarshalJSON() ([]byte, error) {
    type Alias Managementunitlisting

    if ManagementunitlistingMarshalled {
        return []byte("{}"), nil
    }
    ManagementunitlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Managementunit `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        
        Entities: []Managementunit{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

