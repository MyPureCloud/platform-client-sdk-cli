package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationlistingDud struct { 
    


    


    


    

}

// V3synchronizationlisting
type V3synchronizationlisting struct { 
    // Entities
    Entities []V3synchronization `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *V3synchronizationlisting) String() string {
     o.Entities = []V3synchronization{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationlisting) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationlisting

    if V3synchronizationlistingMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []V3synchronization `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []V3synchronization{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

