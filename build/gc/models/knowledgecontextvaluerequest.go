package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgecontextvaluerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgecontextvaluerequestDud struct { 
    


    

}

// Knowledgecontextvaluerequest
type Knowledgecontextvaluerequest struct { 
    // Name - Context value name.
    Name string `json:"name"`


    // Description - Context value description.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Knowledgecontextvaluerequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgecontextvaluerequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgecontextvaluerequest

    if KnowledgecontextvaluerequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgecontextvaluerequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

