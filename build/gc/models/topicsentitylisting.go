package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopicsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopicsentitylistingDud struct { 
    


    


    


    


    


    

}

// Topicsentitylisting
type Topicsentitylisting struct { 
    // Entities
    Entities []Listedtopic `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // Total
    Total int `json:"total"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Topicsentitylisting) String() string {
    
    
     o.Entities = []Listedtopic{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topicsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Topicsentitylisting

    if TopicsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TopicsentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Listedtopic `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        Total int `json:"total"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        
        *Alias
    }{
        

        
        Entities: []Listedtopic{{}},
        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

