package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotconnectorintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotconnectorintentDud struct { 
    


    

}

// Botconnectorintent - A bot intention
type Botconnectorintent struct { 
    // Name - The name of this intent.
    Name string `json:"name"`


    // Entities - A list of entity values that can be associated with this intent
    Entities []Botentity `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Botconnectorintent) String() string {
    
     o.Entities = []Botentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botconnectorintent) MarshalJSON() ([]byte, error) {
    type Alias Botconnectorintent

    if BotconnectorintentMarshalled {
        return []byte("{}"), nil
    }
    BotconnectorintentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Entities []Botentity `json:"entities"`
        *Alias
    }{

        


        
        Entities: []Botentity{{}},
        

        Alias: (*Alias)(u),
    })
}

