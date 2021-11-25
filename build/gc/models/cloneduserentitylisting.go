package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CloneduserentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CloneduserentitylistingDud struct { 
    


    


    

}

// Cloneduserentitylisting
type Cloneduserentitylisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Cloneduser `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Cloneduserentitylisting) String() string {
    
    
    
    
    
    
     o.Entities = []Cloneduser{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cloneduserentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Cloneduserentitylisting

    if CloneduserentitylistingMarshalled {
        return []byte("{}"), nil
    }
    CloneduserentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        Entities []Cloneduser `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        
        Entities: []Cloneduser{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

