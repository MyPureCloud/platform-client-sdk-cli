package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateguideversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateguideversionDud struct { 
    


    


    


    

}

// Updateguideversion - Request body for updating a guide version
type Updateguideversion struct { 
    // Instruction - The instruction given to this version of the guide, for how it should behave when interacting with a User.
    Instruction string `json:"instruction"`


    // Variables - The variables associated with this version of the guide. Includes input variables (provided) and output variables (captured during execution).
    Variables []Variable `json:"variables"`


    // Resources - The resources associated with this version of the guide.
    Resources Guideversionresources `json:"resources"`


    // KnowledgeSettings - The knowledge settings associated with this version of the guide.
    KnowledgeSettings Authoringknowledgesettings `json:"knowledgeSettings"`

}

// String returns a JSON representation of the model
func (o *Updateguideversion) String() string {
    
     o.Variables = []Variable{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateguideversion) MarshalJSON() ([]byte, error) {
    type Alias Updateguideversion

    if UpdateguideversionMarshalled {
        return []byte("{}"), nil
    }
    UpdateguideversionMarshalled = true

    return json.Marshal(&struct {
        
        Instruction string `json:"instruction"`
        
        Variables []Variable `json:"variables"`
        
        Resources Guideversionresources `json:"resources"`
        
        KnowledgeSettings Authoringknowledgesettings `json:"knowledgeSettings"`
        *Alias
    }{

        


        
        Variables: []Variable{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

