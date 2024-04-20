package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EngineintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EngineintegrationDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Engineintegration
type Engineintegration struct { 
    


    // Name - Name of the transcription engine
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Engineintegration) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Engineintegration) MarshalJSON() ([]byte, error) {
    type Alias Engineintegration

    if EngineintegrationMarshalled {
        return []byte("{}"), nil
    }
    EngineintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

