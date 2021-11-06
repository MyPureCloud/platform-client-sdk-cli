package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgecontextrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgecontextrequestDud struct { 
    


    

}

// Knowledgecontextrequest
type Knowledgecontextrequest struct { 
    // Name - Context name.
    Name string `json:"name"`


    // Description - Context description.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Knowledgecontextrequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgecontextrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgecontextrequest

    if KnowledgecontextrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgecontextrequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

