package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DynamiclistvaluesallofMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DynamiclistvaluesallofDud struct { 
    


    


    


    

}

// Dynamiclistvaluesallof
type Dynamiclistvaluesallof struct { 
    // DataActionId - The ID of the data action to invoke at runtime to retrieve list values and synonyms.
    DataActionId string `json:"dataActionId"`


    // Inputs - Array of input mappings for the data action. Maps guide variables to data action input parameters.
    Inputs []Dataactioninput `json:"inputs"`


    // FieldMapping
    FieldMapping Fieldmapping `json:"fieldMapping"`


    // MatchType - Defines how matching should work at runtime. Only 'Exact' matching is supported for dynamic lists.
    MatchType string `json:"matchType"`

}

// String returns a JSON representation of the model
func (o *Dynamiclistvaluesallof) String() string {
    
     o.Inputs = []Dataactioninput{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dynamiclistvaluesallof) MarshalJSON() ([]byte, error) {
    type Alias Dynamiclistvaluesallof

    if DynamiclistvaluesallofMarshalled {
        return []byte("{}"), nil
    }
    DynamiclistvaluesallofMarshalled = true

    return json.Marshal(&struct {
        
        DataActionId string `json:"dataActionId"`
        
        Inputs []Dataactioninput `json:"inputs"`
        
        FieldMapping Fieldmapping `json:"fieldMapping"`
        
        MatchType string `json:"matchType"`
        *Alias
    }{

        


        
        Inputs: []Dataactioninput{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

