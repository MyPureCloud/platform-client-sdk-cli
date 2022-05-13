package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictorlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictorlistingDud struct { 
    


    


    


    

}

// Predictorlisting
type Predictorlisting struct { 
    // Entities
    Entities []Predictor `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Predictorlisting) String() string {
     o.Entities = []Predictor{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictorlisting) MarshalJSON() ([]byte, error) {
    type Alias Predictorlisting

    if PredictorlistingMarshalled {
        return []byte("{}"), nil
    }
    PredictorlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Predictor `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Predictor{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

