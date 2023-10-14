package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediastatisticstrunkinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediastatisticstrunkinfoDud struct { 
    


    


    

}

// Mediastatisticstrunkinfo
type Mediastatisticstrunkinfo struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // VarType
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Mediastatisticstrunkinfo) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediastatisticstrunkinfo) MarshalJSON() ([]byte, error) {
    type Alias Mediastatisticstrunkinfo

    if MediastatisticstrunkinfoMarshalled {
        return []byte("{}"), nil
    }
    MediastatisticstrunkinfoMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

