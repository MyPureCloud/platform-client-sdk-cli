package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScheduledtriggerentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScheduledtriggerentitylistingDud struct { 
    


    


    


    

}

// Scheduledtriggerentitylisting
type Scheduledtriggerentitylisting struct { 
    // Entities
    Entities []Scheduledtrigger `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Scheduledtriggerentitylisting) String() string {
     o.Entities = []Scheduledtrigger{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scheduledtriggerentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Scheduledtriggerentitylisting

    if ScheduledtriggerentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ScheduledtriggerentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Scheduledtrigger `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Scheduledtrigger{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

