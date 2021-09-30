package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetmetricresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetmetricresponseDud struct { 
    


    


    

}

// Getmetricresponse
type Getmetricresponse struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Metric `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Getmetricresponse) String() string {
    
    
    
    
    
    
     o.Entities = []Metric{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getmetricresponse) MarshalJSON() ([]byte, error) {
    type Alias Getmetricresponse

    if GetmetricresponseMarshalled {
        return []byte("{}"), nil
    }
    GetmetricresponseMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        Entities []Metric `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        
        Entities: []Metric{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

