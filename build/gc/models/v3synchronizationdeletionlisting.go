package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationdeletionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationdeletionlistingDud struct { 
    


    


    


    

}

// V3synchronizationdeletionlisting
type V3synchronizationdeletionlisting struct { 
    // Entities
    Entities []V3synchronizationdeletion `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *V3synchronizationdeletionlisting) String() string {
     o.Entities = []V3synchronizationdeletion{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationdeletionlisting) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationdeletionlisting

    if V3synchronizationdeletionlistingMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationdeletionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []V3synchronizationdeletion `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []V3synchronizationdeletion{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

