package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CursorcontactlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CursorcontactlistingDud struct { 
    


    


    


    

}

// Cursorcontactlisting
type Cursorcontactlisting struct { 
    // Entities
    Entities []Externalcontact `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Cursorcontactlisting) String() string {
     o.Entities = []Externalcontact{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cursorcontactlisting) MarshalJSON() ([]byte, error) {
    type Alias Cursorcontactlisting

    if CursorcontactlistingMarshalled {
        return []byte("{}"), nil
    }
    CursorcontactlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Externalcontact `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Externalcontact{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

