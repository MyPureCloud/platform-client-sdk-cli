package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionsresponseDud struct { 
    


    


    


    

}

// Sessionsresponse
type Sessionsresponse struct { 
    // Entities
    Entities []Botflowsession `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Sessionsresponse) String() string {
     o.Entities = []Botflowsession{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessionsresponse) MarshalJSON() ([]byte, error) {
    type Alias Sessionsresponse

    if SessionsresponseMarshalled {
        return []byte("{}"), nil
    }
    SessionsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Botflowsession `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Botflowsession{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

