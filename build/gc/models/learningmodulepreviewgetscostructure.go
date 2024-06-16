package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulepreviewgetscostructureMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulepreviewgetscostructureDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    SuccessStatus string `json:"successStatus"`


    CompletionStatus string `json:"completionStatus"`


    PercentageScore float32 `json:"percentageScore"`


    


    Children []Learningmodulepreviewgetscostructure `json:"children"`

}

// Learningmodulepreviewgetscostructure - Learning module preview get SCO structure
type Learningmodulepreviewgetscostructure struct { 
    


    


    


    


    


    // ShareableContentObject - The SCO (Shareable Content Object) data
    ShareableContentObject Learningshareablecontentobject `json:"shareableContentObject"`


    

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewgetscostructure) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulepreviewgetscostructure) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulepreviewgetscostructure

    if LearningmodulepreviewgetscostructureMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulepreviewgetscostructureMarshalled = true

    return json.Marshal(&struct {
        
        ShareableContentObject Learningshareablecontentobject `json:"shareableContentObject"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

