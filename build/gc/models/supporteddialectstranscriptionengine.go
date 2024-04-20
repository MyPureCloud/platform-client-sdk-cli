package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupporteddialectstranscriptionengineMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupporteddialectstranscriptionengineDud struct { 
    


    


    

}

// Supporteddialectstranscriptionengine
type Supporteddialectstranscriptionengine struct { 
    // Engine
    Engine string `json:"engine"`


    // Dialects
    Dialects []string `json:"dialects"`


    // EngineIntegration
    EngineIntegration Engineintegration `json:"engineIntegration"`

}

// String returns a JSON representation of the model
func (o *Supporteddialectstranscriptionengine) String() string {
    
     o.Dialects = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supporteddialectstranscriptionengine) MarshalJSON() ([]byte, error) {
    type Alias Supporteddialectstranscriptionengine

    if SupporteddialectstranscriptionengineMarshalled {
        return []byte("{}"), nil
    }
    SupporteddialectstranscriptionengineMarshalled = true

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

