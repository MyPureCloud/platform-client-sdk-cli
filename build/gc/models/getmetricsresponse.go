package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetmetricsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetmetricsresponseDud struct { 
    


    


    

}

// Getmetricsresponse
type Getmetricsresponse struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Metrics `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Getmetricsresponse) String() string {
    
     o.Entities = []Metrics{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getmetricsresponse) MarshalJSON() ([]byte, error) {
    type Alias Getmetricsresponse

    if GetmetricsresponseMarshalled {
        return []byte("{}"), nil
    }
    GetmetricsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Metrics `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Metrics{{}},
        


        

        Alias: (*Alias)(u),
    })
}

