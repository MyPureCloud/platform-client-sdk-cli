package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DirectoryuserdeviceslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DirectoryuserdeviceslistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Directoryuserdeviceslisting
type Directoryuserdeviceslisting struct { 
    // Entities
    Entities []Userdevice `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Directoryuserdeviceslisting) String() string {
     o.Entities = []Userdevice{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Directoryuserdeviceslisting) MarshalJSON() ([]byte, error) {
    type Alias Directoryuserdeviceslisting

    if DirectoryuserdeviceslistingMarshalled {
        return []byte("{}"), nil
    }
    DirectoryuserdeviceslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Userdevice `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Userdevice{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

