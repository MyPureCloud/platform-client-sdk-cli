package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableexportjoblistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableexportjoblistingDud struct { 
    


    


    


    

}

// Decisiontableexportjoblisting
type Decisiontableexportjoblisting struct { 
    // Entities
    Entities []Decisiontableexportjob `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Decisiontableexportjoblisting) String() string {
     o.Entities = []Decisiontableexportjob{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableexportjoblisting) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableexportjoblisting

    if DecisiontableexportjoblistingMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableexportjoblistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Decisiontableexportjob `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Decisiontableexportjob{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

