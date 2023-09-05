package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmpathyscoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmpathyscoreDud struct { 
    


    

}

// Empathyscore
type Empathyscore struct { 
    // Score - Empathy score of the agent involved in the conversation
    Score float64 `json:"score"`


    // UserId - UserId of the agent involved in the conversation
    UserId string `json:"userId"`

}

// String returns a JSON representation of the model
func (o *Empathyscore) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Empathyscore) MarshalJSON() ([]byte, error) {
    type Alias Empathyscore

    if EmpathyscoreMarshalled {
        return []byte("{}"), nil
    }
    EmpathyscoreMarshalled = true

    return json.Marshal(&struct {
        
        Score float64 `json:"score"`
        
        UserId string `json:"userId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

