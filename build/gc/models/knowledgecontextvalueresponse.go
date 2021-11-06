package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgecontextvalueresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgecontextvalueresponseDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgecontextvalueresponse
type Knowledgecontextvalueresponse struct { 
    // Id - Context value ID.
    Id string `json:"id"`


    // Name - Context value name.
    Name string `json:"name"`


    // Description - Context value description.
    Description string `json:"description"`


    // DateCreated - The date when the context value was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date when the context value was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgecontextvalueresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgecontextvalueresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgecontextvalueresponse

    if KnowledgecontextvalueresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgecontextvalueresponseMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

