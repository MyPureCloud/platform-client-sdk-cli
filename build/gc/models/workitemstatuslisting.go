package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemstatuslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemstatuslistingDud struct { 
    


    


    

}

// Workitemstatuslisting
type Workitemstatuslisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Workitemstatus `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Workitemstatuslisting) String() string {
    
     o.Entities = []Workitemstatus{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemstatuslisting) MarshalJSON() ([]byte, error) {
    type Alias Workitemstatuslisting

    if WorkitemstatuslistingMarshalled {
        return []byte("{}"), nil
    }
    WorkitemstatuslistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Workitemstatus `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Workitemstatus{{}},
        


        

        Alias: (*Alias)(u),
    })
}

