package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewupdatescostructureMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewupdatescostructureDud struct { 
    


    


    


    


    


    

}

// Learningmodulepreviewupdatescostructure - Learning module preview update SCO structure
type Learningmodulepreviewupdatescostructure struct { 
    // Id - The id of this SCO in the course manifest
    Id string `json:"id"`


    // Name - The name of this SCO in the course manifest
    Name string `json:"name"`


    // SuccessStatus - The success status of this SCO
    SuccessStatus string `json:"successStatus"`


    // CompletionStatus - The completion status of this SCO
    CompletionStatus string `json:"completionStatus"`


    // PercentageScore - Percentage Score
    PercentageScore float32 `json:"percentageScore"`


    // Children - Child items belonging to this SCO in the course manifest
    Children []Learningmodulepreviewupdatescostructure `json:"children"`

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdatescostructure) String() string {
    
    
    
    
    
     o.Children = []Learningmodulepreviewupdatescostructure{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewupdatescostructure) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewupdatescostructure

    if LearningmodulepreviewupdatescostructureMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewupdatescostructureMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SuccessStatus string `json:"successStatus"`
        
        CompletionStatus string `json:"completionStatus"`
        
        PercentageScore float32 `json:"percentageScore"`
        
        Children []Learningmodulepreviewupdatescostructure `json:"children"`
        *Alias
    }{

        


        


        


        


        


        
        Children: []Learningmodulepreviewupdatescostructure{{}},
        

        Alias: (*Alias)(u),
    })
}

