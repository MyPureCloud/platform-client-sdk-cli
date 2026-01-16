package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuidesessionturnrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuidesessionturnrequestDud struct { 
    


    


    


    


    

}

// Guidesessionturnrequest - Request for a guide session turn
type Guidesessionturnrequest struct { 
    // InputEvent - The input event for this turn.
    InputEvent Guidesessioninputevent `json:"inputEvent"`


    // LanguageCode - The language code for this turn.
    LanguageCode string `json:"languageCode"`


    // Version - The version for this turn.
    Version string `json:"version"`


    // InputVariables - The input variables for this turn.
    InputVariables []Guidesessionvariable `json:"inputVariables"`


    // KnowledgeSettings - The knowledge settings for this turn.
    KnowledgeSettings Knowledgesettings `json:"knowledgeSettings"`

}

// String returns a JSON representation of the model
func (o *Guidesessionturnrequest) String() string {
    
    
    
     o.InputVariables = []Guidesessionvariable{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guidesessionturnrequest) MarshalJSON() ([]byte, error) {
    type Alias Guidesessionturnrequest

    if GuidesessionturnrequestMarshalled {
        return []byte("{}"), nil
    }
    GuidesessionturnrequestMarshalled = true

    return json.Marshal(&struct {
        
        InputEvent Guidesessioninputevent `json:"inputEvent"`
        
        LanguageCode string `json:"languageCode"`
        
        Version string `json:"version"`
        
        InputVariables []Guidesessionvariable `json:"inputVariables"`
        
        KnowledgeSettings Knowledgesettings `json:"knowledgeSettings"`
        *Alias
    }{

        


        


        


        
        InputVariables: []Guidesessionvariable{{}},
        


        

        Alias: (*Alias)(u),
    })
}

