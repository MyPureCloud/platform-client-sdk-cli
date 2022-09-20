package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TriggerentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TriggerentitylistingDud struct { 
    


    


    


    

}

// Triggerentitylisting
type Triggerentitylisting struct { 
    // Entities
    Entities []Trigger `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Triggerentitylisting) String() string {
     o.Entities = []Trigger{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Triggerentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Triggerentitylisting

    if TriggerentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TriggerentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Trigger `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Trigger{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

