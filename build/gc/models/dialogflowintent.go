package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialogflowintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialogflowintentDud struct { 
    


    

}

// Dialogflowintent
type Dialogflowintent struct { 
    // Name - The intent name
    Name string `json:"name"`


    // Parameters - An object mapping parameter names to Parameter objects
    Parameters map[string]Dialogflowparameter `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Dialogflowintent) String() string {
    
     o.Parameters = map[string]Dialogflowparameter{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialogflowintent) MarshalJSON() ([]byte, error) {
    type Alias Dialogflowintent

    if DialogflowintentMarshalled {
        return []byte("{}"), nil
    }
    DialogflowintentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Parameters map[string]Dialogflowparameter `json:"parameters"`
        *Alias
    }{

        


        
        Parameters: map[string]Dialogflowparameter{"": {}},
        

        Alias: (*Alias)(u),
    })
}

