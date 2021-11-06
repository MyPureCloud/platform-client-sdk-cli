package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgecontextresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgecontextresponseDud struct { 
    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgecontextresponse
type Knowledgecontextresponse struct { 
    // Id - Context ID.
    Id string `json:"id"`


    // Name - Context name.
    Name string `json:"name"`


    // Description - Context description.
    Description string `json:"description"`


    // DateCreated - The date when the context was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date when the context was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // Values - Knowledge context values.
    Values []Knowledgecontextvalueresponse `json:"values"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgecontextresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Values = []Knowledgecontextvalueresponse{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgecontextresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgecontextresponse

    if KnowledgecontextresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgecontextresponseMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        Values []Knowledgecontextvalueresponse `json:"values"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        Values: []Knowledgecontextvalueresponse{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

