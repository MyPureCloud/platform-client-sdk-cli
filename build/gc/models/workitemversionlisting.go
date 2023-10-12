package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemversionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemversionlistingDud struct { 
    


    


    


    


    

}

// Workitemversionlisting
type Workitemversionlisting struct { 
    // Entities
    Entities []Workitemversion `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // After
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Workitemversionlisting) String() string {
     o.Entities = []Workitemversion{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemversionlisting) MarshalJSON() ([]byte, error) {
    type Alias Workitemversionlisting

    if WorkitemversionlistingMarshalled {
        return []byte("{}"), nil
    }
    WorkitemversionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitemversion `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Workitemversion{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

