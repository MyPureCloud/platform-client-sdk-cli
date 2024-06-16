package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OperandpositionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OperandpositionDud struct { 
    


    


    


    

}

// Operandposition
type Operandposition struct { 
    // StartingPositionValue - Defines starting point of a position range - number of seconds or words from the start or from the end of the interaction
    StartingPositionValue int `json:"startingPositionValue"`


    // StartingPositionDirection - Dictates starting position directionality
    StartingPositionDirection string `json:"startingPositionDirection"`


    // EndingPositionValue - Defines ending point of a position range - number of seconds or words from the start or from the end of the interaction
    EndingPositionValue int `json:"endingPositionValue"`


    // EndingPositionDirection - Dictates ending position directionality
    EndingPositionDirection string `json:"endingPositionDirection"`

}

// String returns a JSON representation of the model
func (o *Operandposition) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Operandposition) MarshalJSON() ([]byte, error) {
    type Alias Operandposition

    if OperandpositionMarshalled {
        return []byte("{}"), nil
    }
    OperandpositionMarshalled = true

    return json.Marshal(&struct {
        
        StartingPositionValue int `json:"startingPositionValue"`
        
        StartingPositionDirection string `json:"startingPositionDirection"`
        
        EndingPositionValue int `json:"endingPositionValue"`
        
        EndingPositionDirection string `json:"endingPositionDirection"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

