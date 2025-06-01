package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DashboardconfigurationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DashboardconfigurationlistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Dashboardconfigurationlisting
type Dashboardconfigurationlisting struct { 
    // Entities
    Entities []Dashboardconfiguration `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // LastUri
    LastUri string `json:"lastUri"`


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
func (o *Dashboardconfigurationlisting) String() string {
     o.Entities = []Dashboardconfiguration{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dashboardconfigurationlisting) MarshalJSON() ([]byte, error) {
    type Alias Dashboardconfigurationlisting

    if DashboardconfigurationlistingMarshalled {
        return []byte("{}"), nil
    }
    DashboardconfigurationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Dashboardconfiguration `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Dashboardconfiguration{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

