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
    Id string `json:"id"`


    Name string `json:"name"`


    SuccessStatus string `json:"successStatus"`


    CompletionStatus string `json:"completionStatus"`


    PercentageScore float32 `json:"percentageScore"`


    Children []Learningmodulepreviewupdatescostructure `json:"children"`

}

// Learningmodulepreviewupdatescostructure - Learning module preview update SCO structure
type Learningmodulepreviewupdatescostructure struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdatescostructure) String() string {

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
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

