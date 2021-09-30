package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CursororganizationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CursororganizationlistingDud struct { 
    


    


    


    

}

// Cursororganizationlisting
type Cursororganizationlisting struct { 
    // Entities
    Entities []Externalorganization `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Cursororganizationlisting) String() string {
    
    
     o.Entities = []Externalorganization{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cursororganizationlisting) MarshalJSON() ([]byte, error) {
    type Alias Cursororganizationlisting

    if CursororganizationlistingMarshalled {
        return []byte("{}"), nil
    }
    CursororganizationlistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Externalorganization `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        *Alias
    }{
        

        
        Entities: []Externalorganization{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

