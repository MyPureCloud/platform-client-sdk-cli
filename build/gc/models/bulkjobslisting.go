package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkjobslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkjobslistingDud struct { 
    


    


    


    


    

}

// Bulkjobslisting
type Bulkjobslisting struct { 
    // Entities
    Entities []Bulkjob `json:"entities"`


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
func (o *Bulkjobslisting) String() string {
     o.Entities = []Bulkjob{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkjobslisting) MarshalJSON() ([]byte, error) {
    type Alias Bulkjobslisting

    if BulkjobslistingMarshalled {
        return []byte("{}"), nil
    }
    BulkjobslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Bulkjob `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Bulkjob{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

