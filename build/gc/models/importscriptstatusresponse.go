package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ImportscriptstatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ImportscriptstatusresponseDud struct { 
    


    


    

}

// Importscriptstatusresponse
type Importscriptstatusresponse struct { 
    // Url
    Url string `json:"url"`


    // Succeeded
    Succeeded bool `json:"succeeded"`


    // Message
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Importscriptstatusresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Importscriptstatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Importscriptstatusresponse

    if ImportscriptstatusresponseMarshalled {
        return []byte("{}"), nil
    }
    ImportscriptstatusresponseMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        Succeeded bool `json:"succeeded"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

