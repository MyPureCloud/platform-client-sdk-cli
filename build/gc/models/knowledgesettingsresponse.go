package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesettingsresponseDud struct { 
    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgesettingsresponse
type Knowledgesettingsresponse struct { 
    // Id - Knowledge Setting Id.
    Id string `json:"id"`


    // Name - Knowledge Setting Name.
    Name string `json:"name"`


    // Description - Knowledge setting description.
    Description string `json:"description"`


    // Sources - Knowledge source information searched upon.
    Sources []V3sourceref `json:"sources"`


    // GenerationSetting - Settings for answer generation.
    GenerationSetting Knowledgegenerationsetting `json:"generationSetting"`


    // Stateful - Indicates if stateful search and generation is enabled for the knowledge setting.
    Stateful bool `json:"stateful"`


    // DateCreated - Knowledge setting created date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Knowledge setting last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The user who modified the knowledge setting.
    ModifiedBy Userreference `json:"modifiedBy"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgesettingsresponse) String() string {
    
    
    
     o.Sources = []V3sourceref{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesettingsresponse

    if KnowledgesettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Sources []V3sourceref `json:"sources"`
        
        GenerationSetting Knowledgegenerationsetting `json:"generationSetting"`
        
        Stateful bool `json:"stateful"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        *Alias
    }{

        


        


        


        
        Sources: []V3sourceref{{}},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

