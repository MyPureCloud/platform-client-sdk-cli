package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Queuememberentitylistingv1Marshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Queuememberentitylistingv1Dud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Queuememberentitylistingv1
type Queuememberentitylistingv1 struct { 
    // Entities
    Entities []Queuemember `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Queuememberentitylistingv1) String() string {
     o.Entities = []Queuemember{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queuememberentitylistingv1) MarshalJSON() ([]byte, error) {
    type Alias Queuememberentitylistingv1

    if Queuememberentitylistingv1Marshalled {
        return []byte("{}"), nil
    }
    Queuememberentitylistingv1Marshalled = true

    return json.Marshal(&struct {
        
        Entities []Queuemember `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Queuemember{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

