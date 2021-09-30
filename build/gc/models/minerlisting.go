package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinerlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinerlistingDud struct { 
    


    


    


    

}

// Minerlisting
type Minerlisting struct { 
    // Entities
    Entities []Miner `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Minerlisting) String() string {
    
    
     o.Entities = []Miner{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Minerlisting) MarshalJSON() ([]byte, error) {
    type Alias Minerlisting

    if MinerlistingMarshalled {
        return []byte("{}"), nil
    }
    MinerlistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Miner `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        *Alias
    }{
        

        
        Entities: []Miner{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

