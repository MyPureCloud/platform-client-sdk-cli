package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseplanqueryentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseplanqueryentitylistingDud struct { 
    


    

}

// Caseplanqueryentitylisting
type Caseplanqueryentitylisting struct { 
    // Entities
    Entities []Caseplan `json:"entities"`


    // After
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Caseplanqueryentitylisting) String() string {
     o.Entities = []Caseplan{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseplanqueryentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Caseplanqueryentitylisting

    if CaseplanqueryentitylistingMarshalled {
        return []byte("{}"), nil
    }
    CaseplanqueryentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Caseplan `json:"entities"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Caseplan{{}},
        


        

        Alias: (*Alias)(u),
    })
}

