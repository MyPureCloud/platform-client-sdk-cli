package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuideversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuideversionDud struct { 
    


    Guide Addressableentityref `json:"guide"`


    Version string `json:"version"`


    Instruction string `json:"instruction"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    Variables []Variable `json:"variables"`


    


    

}

// Guideversion
type Guideversion struct { 
    // SelfUri
    SelfUri string `json:"selfUri"`


    


    


    


    // State - The current state of the guide version.
    State string `json:"state"`


    


    


    


    // Resources - The resources associated with this version of the guide.
    Resources Guideversionresources `json:"resources"`


    // KnowledgeSettings - The knowledge settings associated with this version of the guide.
    KnowledgeSettings Authoringknowledgesettings `json:"knowledgeSettings"`

}

// String returns a JSON representation of the model
func (o *Guideversion) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guideversion) MarshalJSON() ([]byte, error) {
    type Alias Guideversion

    if GuideversionMarshalled {
        return []byte("{}"), nil
    }
    GuideversionMarshalled = true

    return json.Marshal(&struct {
        
        SelfUri string `json:"selfUri"`
        
        State string `json:"state"`
        
        Resources Guideversionresources `json:"resources"`
        
        KnowledgeSettings Authoringknowledgesettings `json:"knowledgeSettings"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

