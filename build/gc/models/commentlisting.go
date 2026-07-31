package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommentlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommentlistingDud struct { 
    


    


    


    

}

// Commentlisting
type Commentlisting struct { 
    // Entities
    Entities []Comment `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Commentlisting) String() string {
     o.Entities = []Comment{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commentlisting) MarshalJSON() ([]byte, error) {
    type Alias Commentlisting

    if CommentlistingMarshalled {
        return []byte("{}"), nil
    }
    CommentlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Comment `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Comment{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

