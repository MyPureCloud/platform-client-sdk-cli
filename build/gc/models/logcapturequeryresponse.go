package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LogcapturequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LogcapturequeryresponseDud struct { 
    


    


    


    

}

// Logcapturequeryresponse
type Logcapturequeryresponse struct { 
    // Entities
    Entities []Logentry `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Logcapturequeryresponse) String() string {
     o.Entities = []Logentry{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Logcapturequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Logcapturequeryresponse

    if LogcapturequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    LogcapturequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Logentry `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Logentry{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

