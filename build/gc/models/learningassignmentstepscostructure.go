package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentstepscostructureMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentstepscostructureDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    SuccessStatus string `json:"successStatus"`


    CompletionStatus string `json:"completionStatus"`


    Children []Learningassignmentstepscostructure `json:"children"`

}

// Learningassignmentstepscostructure
type Learningassignmentstepscostructure struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Learningassignmentstepscostructure) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentstepscostructure) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentstepscostructure

    if LearningassignmentstepscostructureMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentstepscostructureMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

