package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgramrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgramrequestDud struct { 
    


    


    


    

}

// Programrequest
type Programrequest struct { 
    // Name - The program name
    Name string `json:"name"`


    // Description - The program description
    Description string `json:"description"`


    // TopicIds - The ids of topics associated to the program
    TopicIds []string `json:"topicIds"`


    // Tags - The program tags
    Tags []string `json:"tags"`

}

// String returns a JSON representation of the model
func (o *Programrequest) String() string {
    
    
    
    
    
    
    
    
    
    
     o.TopicIds = []string{""} 
    
    
    
     o.Tags = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programrequest) MarshalJSON() ([]byte, error) {
    type Alias Programrequest

    if ProgramrequestMarshalled {
        return []byte("{}"), nil
    }
    ProgramrequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        TopicIds []string `json:"topicIds"`
        
        Tags []string `json:"tags"`
        
        *Alias
    }{
        

        

        

        

        

        
        TopicIds: []string{""},
        

        

        
        Tags: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

