package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponseassetentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponseassetentitylistingDud struct { 
    


    

}

// Responseassetentitylisting
type Responseassetentitylisting struct { 
    // Entities - List of response assets
    Entities []Responseasset `json:"entities"`


    // NotFound - Asset IDs not found
    NotFound []string `json:"notFound"`

}

// String returns a JSON representation of the model
func (o *Responseassetentitylisting) String() string {
     o.Entities = []Responseasset{{}} 
     o.NotFound = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responseassetentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Responseassetentitylisting

    if ResponseassetentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ResponseassetentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Responseasset `json:"entities"`
        
        NotFound []string `json:"notFound"`
        *Alias
    }{

        
        Entities: []Responseasset{{}},
        


        
        NotFound: []string{""},
        

        Alias: (*Alias)(u),
    })
}

