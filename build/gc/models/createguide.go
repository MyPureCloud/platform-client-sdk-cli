package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateguideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateguideDud struct { 
    


    

}

// Createguide
type Createguide struct { 
    // Name - The name of the guide.
    Name string `json:"name"`


    // Source - Indicates how the guide content was generated.
    Source string `json:"source"`

}

// String returns a JSON representation of the model
func (o *Createguide) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createguide) MarshalJSON() ([]byte, error) {
    type Alias Createguide

    if CreateguideMarshalled {
        return []byte("{}"), nil
    }
    CreateguideMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Source string `json:"source"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

