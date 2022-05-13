package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GettemplatesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GettemplatesresponseDud struct { 
    


    


    

}

// Gettemplatesresponse
type Gettemplatesresponse struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Objectivetemplate `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Gettemplatesresponse) String() string {
    
     o.Entities = []Objectivetemplate{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Gettemplatesresponse) MarshalJSON() ([]byte, error) {
    type Alias Gettemplatesresponse

    if GettemplatesresponseMarshalled {
        return []byte("{}"), nil
    }
    GettemplatesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Objectivetemplate `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Objectivetemplate{{}},
        


        

        Alias: (*Alias)(u),
    })
}

