package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetprofilesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetprofilesresponseDud struct { 
    


    


    

}

// Getprofilesresponse
type Getprofilesresponse struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Performanceprofile `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Getprofilesresponse) String() string {
    
     o.Entities = []Performanceprofile{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getprofilesresponse) MarshalJSON() ([]byte, error) {
    type Alias Getprofilesresponse

    if GetprofilesresponseMarshalled {
        return []byte("{}"), nil
    }
    GetprofilesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Performanceprofile `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Performanceprofile{{}},
        


        

        Alias: (*Alias)(u),
    })
}

