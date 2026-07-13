package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationuploadlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationuploadlistingDud struct { 
    


    


    


    

}

// V3synchronizationuploadlisting
type V3synchronizationuploadlisting struct { 
    // Entities
    Entities []V3synchronizationupload `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *V3synchronizationuploadlisting) String() string {
     o.Entities = []V3synchronizationupload{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationuploadlisting) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationuploadlisting

    if V3synchronizationuploadlistingMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationuploadlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []V3synchronizationupload `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []V3synchronizationupload{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

