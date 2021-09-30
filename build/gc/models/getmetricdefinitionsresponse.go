package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetmetricdefinitionsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetmetricdefinitionsresponseDud struct { 
    


    


    

}

// Getmetricdefinitionsresponse
type Getmetricdefinitionsresponse struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Metricdefinition `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Getmetricdefinitionsresponse) String() string {
    
    
    
    
    
    
     o.Entities = []Metricdefinition{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getmetricdefinitionsresponse) MarshalJSON() ([]byte, error) {
    type Alias Getmetricdefinitionsresponse

    if GetmetricdefinitionsresponseMarshalled {
        return []byte("{}"), nil
    }
    GetmetricdefinitionsresponseMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        Entities []Metricdefinition `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        
        Entities: []Metricdefinition{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

