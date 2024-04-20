package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgramtranscriptionengineMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgramtranscriptionengineDud struct { 
    


    


    

}

// Programtranscriptionengine
type Programtranscriptionengine struct { 
    // Engine
    Engine string `json:"engine"`


    // Dialects
    Dialects []string `json:"dialects"`


    // EngineIntegration
    EngineIntegration Engineintegration `json:"engineIntegration"`

}

// String returns a JSON representation of the model
func (o *Programtranscriptionengine) String() string {
    
     o.Dialects = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programtranscriptionengine) MarshalJSON() ([]byte, error) {
    type Alias Programtranscriptionengine

    if ProgramtranscriptionengineMarshalled {
        return []byte("{}"), nil
    }
    ProgramtranscriptionengineMarshalled = true

    return json.Marshal(&struct {
        
        Engine string `json:"engine"`
        
        Dialects []string `json:"dialects"`
        
        EngineIntegration Engineintegration `json:"engineIntegration"`
        *Alias
    }{

        


        
        Dialects: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

