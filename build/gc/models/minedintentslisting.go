package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinedintentslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinedintentslistingDud struct { 
    


    


    


    

}

// Minedintentslisting
type Minedintentslisting struct { 
    // Entities
    Entities []Minerintent `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Minedintentslisting) String() string {
    
    
     o.Entities = []Minerintent{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Minedintentslisting) MarshalJSON() ([]byte, error) {
    type Alias Minedintentslisting

    if MinedintentslistingMarshalled {
        return []byte("{}"), nil
    }
    MinedintentslistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Minerintent `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        *Alias
    }{
        

        
        Entities: []Minerintent{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

