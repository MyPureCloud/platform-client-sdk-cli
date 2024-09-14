package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AudioconnectorintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AudioconnectorintegrationDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Audioconnectorintegration
type Audioconnectorintegration struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Audioconnectorintegration) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Audioconnectorintegration) MarshalJSON() ([]byte, error) {
    type Alias Audioconnectorintegration

    if AudioconnectorintegrationMarshalled {
        return []byte("{}"), nil
    }
    AudioconnectorintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

