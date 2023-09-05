package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgebaseupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgebaseupdaterequestDud struct { 
    


    

}

// Knowledgebaseupdaterequest
type Knowledgebaseupdaterequest struct { 
    // Name - Knowledge base name
    Name string `json:"name"`


    // Description - Knowledge base description
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Knowledgebaseupdaterequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgebaseupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgebaseupdaterequest

    if KnowledgebaseupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgebaseupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

