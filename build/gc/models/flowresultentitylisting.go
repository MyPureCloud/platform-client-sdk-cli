package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowresultentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowresultentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Flowresultentitylisting
type Flowresultentitylisting struct { 
    // Entities
    Entities []Flowexecutiondataqueryresult `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Flowresultentitylisting) String() string {
     o.Entities = []Flowexecutiondataqueryresult{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowresultentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Flowresultentitylisting

    if FlowresultentitylistingMarshalled {
        return []byte("{}"), nil
    }
    FlowresultentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Flowexecutiondataqueryresult `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Flowexecutiondataqueryresult{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

