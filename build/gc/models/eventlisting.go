package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventlistingDud struct { 
    


    


    


    

}

// Eventlisting
type Eventlisting struct { 
    // Entities
    Entities []Event `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Eventlisting) String() string {
     o.Entities = []Event{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventlisting) MarshalJSON() ([]byte, error) {
    type Alias Eventlisting

    if EventlistingMarshalled {
        return []byte("{}"), nil
    }
    EventlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Event `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Event{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

