package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacetkeyattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacetkeyattributeDud struct { 
    


    


    

}

// Facetkeyattribute
type Facetkeyattribute struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Count
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Facetkeyattribute) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facetkeyattribute) MarshalJSON() ([]byte, error) {
    type Alias Facetkeyattribute

    if FacetkeyattributeMarshalled {
        return []byte("{}"), nil
    }
    FacetkeyattributeMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

