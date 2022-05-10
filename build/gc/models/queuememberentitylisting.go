package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueuememberentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueuememberentitylistingDud struct { 
    


    


    


    


    


    


    

}

// Queuememberentitylisting
type Queuememberentitylisting struct { 
    // Entities
    Entities []Queuemember `json:"entities"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // PageSize
    PageSize int `json:"pageSize"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`

}

// String returns a JSON representation of the model
func (o *Queuememberentitylisting) String() string {
    
    
     o.Entities = []Queuemember{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queuememberentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Queuememberentitylisting

    if QueuememberentitylistingMarshalled {
        return []byte("{}"), nil
    }
    QueuememberentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Queuemember `json:"entities"`
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        *Alias
    }{
        

        
        Entities: []Queuemember{{}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

