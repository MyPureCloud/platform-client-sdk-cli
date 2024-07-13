package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictedanswerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictedanswerDud struct { 
    


    


    

}

// Predictedanswer
type Predictedanswer struct { 
    // AnswerId - The unique identifier of the suggested predicted answer.
    AnswerId string `json:"answerId"`


    // Explanation - An explanation providing the reasoning behind the suggested answer.
    Explanation string `json:"explanation"`


    // FailureType - Describes the type of error associated with the predicted answer.
    FailureType string `json:"failureType"`

}

// String returns a JSON representation of the model
func (o *Predictedanswer) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictedanswer) MarshalJSON() ([]byte, error) {
    type Alias Predictedanswer

    if PredictedanswerMarshalled {
        return []byte("{}"), nil
    }
    PredictedanswerMarshalled = true

    return json.Marshal(&struct {
        
        AnswerId string `json:"answerId"`
        
        Explanation string `json:"explanation"`
        
        FailureType string `json:"failureType"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

