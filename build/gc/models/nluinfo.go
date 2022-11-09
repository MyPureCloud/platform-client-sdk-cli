package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NluinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NluinfoDud struct { 
    Domain Addressableentityref `json:"domain"`


    Version *Nludomainversion `json:"version"`


    


    


    

}

// Nluinfo
type Nluinfo struct { 
    


    


    // Intents
    Intents []Intent `json:"intents"`


    // EngineVersion
    EngineVersion string `json:"engineVersion"`


    // NluData
    NluData Nludomainversion `json:"nluData"`

}

// String returns a JSON representation of the model
func (o *Nluinfo) String() string {
     o.Intents = []Intent{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nluinfo) MarshalJSON() ([]byte, error) {
    type Alias Nluinfo

    if NluinfoMarshalled {
        return []byte("{}"), nil
    }
    NluinfoMarshalled = true

    return json.Marshal(&struct {
        
        Intents []Intent `json:"intents"`
        
        EngineVersion string `json:"engineVersion"`
        
        NluData Nludomainversion `json:"nluData"`
        *Alias
    }{

        


        


        
        Intents: []Intent{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

