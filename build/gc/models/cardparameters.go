package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CardparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CardparametersDud struct { 
    


    


    

}

// Cardparameters - Template parameters for a single carousel card
type Cardparameters struct { 
    // Index - Index of the card in the carousel template
    Index int `json:"index"`


    // BodyParameters - A list of Response Management carousel card body parameter substitutions for the response's messaging template
    BodyParameters []Templateparameter `json:"bodyParameters"`


    // ButtonUrlParameters - A list of Response Management carousel card button URL parameter substitutions for the response's messaging template
    ButtonUrlParameters []Templateparameter `json:"buttonUrlParameters"`

}

// String returns a JSON representation of the model
func (o *Cardparameters) String() string {
    
     o.BodyParameters = []Templateparameter{{}} 
     o.ButtonUrlParameters = []Templateparameter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cardparameters) MarshalJSON() ([]byte, error) {
    type Alias Cardparameters

    if CardparametersMarshalled {
        return []byte("{}"), nil
    }
    CardparametersMarshalled = true

    return json.Marshal(&struct {
        
        Index int `json:"index"`
        
        BodyParameters []Templateparameter `json:"bodyParameters"`
        
        ButtonUrlParameters []Templateparameter `json:"buttonUrlParameters"`
        *Alias
    }{

        


        
        BodyParameters: []Templateparameter{{}},
        


        
        ButtonUrlParameters: []Templateparameter{{}},
        

        Alias: (*Alias)(u),
    })
}

