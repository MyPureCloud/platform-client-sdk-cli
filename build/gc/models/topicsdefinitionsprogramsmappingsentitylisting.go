package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopicsdefinitionsprogramsmappingsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopicsdefinitionsprogramsmappingsentitylistingDud struct { 
    


    


    


    


    

}

// Topicsdefinitionsprogramsmappingsentitylisting
type Topicsdefinitionsprogramsmappingsentitylisting struct { 
    // Entities
    Entities []Topicsdefinitionsprogrammappings `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Topicsdefinitionsprogramsmappingsentitylisting) String() string {
     o.Entities = []Topicsdefinitionsprogrammappings{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topicsdefinitionsprogramsmappingsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Topicsdefinitionsprogramsmappingsentitylisting

    if TopicsdefinitionsprogramsmappingsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TopicsdefinitionsprogramsmappingsentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Topicsdefinitionsprogrammappings `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Topicsdefinitionsprogrammappings{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

