package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkbinversionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkbinversionlistingDud struct { 
    


    


    


    


    

}

// Workbinversionlisting
type Workbinversionlisting struct { 
    // Entities
    Entities []Workbinversion `json:"entities"`


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
func (o *Workbinversionlisting) String() string {
     o.Entities = []Workbinversion{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workbinversionlisting) MarshalJSON() ([]byte, error) {
    type Alias Workbinversionlisting

    if WorkbinversionlistingMarshalled {
        return []byte("{}"), nil
    }
    WorkbinversionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workbinversion `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Workbinversion{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

