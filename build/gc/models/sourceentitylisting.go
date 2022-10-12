package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SourceentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SourceentitylistingDud struct { 
    


    


    

}

// Sourceentitylisting
type Sourceentitylisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Source `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Sourceentitylisting) String() string {
    
     o.Entities = []Source{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sourceentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Sourceentitylisting

    if SourceentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SourceentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Source `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Source{{}},
        


        

        Alias: (*Alias)(u),
    })
}

