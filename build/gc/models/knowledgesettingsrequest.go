package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesettingsrequestDud struct { 
    


    


    


    


    

}

// Knowledgesettingsrequest
type Knowledgesettingsrequest struct { 
    // Name - Knowledge setting name.
    Name string `json:"name"`


    // Description - Knowledge setting description.
    Description string `json:"description"`


    // Sources - Knowledge source information to search upon.
    Sources []V3sourceref `json:"sources"`


    // GenerationSetting - Setting for answer generation.
    GenerationSetting Knowledgegenerationsetting `json:"generationSetting"`


    // Stateful - Indicates if stateful search and generation is enabled for the knowledge setting.
    Stateful bool `json:"stateful"`

}

// String returns a JSON representation of the model
func (o *Knowledgesettingsrequest) String() string {
    
    
     o.Sources = []V3sourceref{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesettingsrequest

    if KnowledgesettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Sources []V3sourceref `json:"sources"`
        
        GenerationSetting Knowledgegenerationsetting `json:"generationSetting"`
        
        Stateful bool `json:"stateful"`
        *Alias
    }{

        


        


        
        Sources: []V3sourceref{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

