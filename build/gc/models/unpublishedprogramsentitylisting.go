package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnpublishedprogramsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnpublishedprogramsentitylistingDud struct { 
    


    


    


    


    

}

// Unpublishedprogramsentitylisting
type Unpublishedprogramsentitylisting struct { 
    // Entities
    Entities []Program `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Unpublishedprogramsentitylisting) String() string {
     o.Entities = []Program{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unpublishedprogramsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Unpublishedprogramsentitylisting

    if UnpublishedprogramsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    UnpublishedprogramsentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Program `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Program{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

