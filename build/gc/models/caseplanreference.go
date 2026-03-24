package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseplanreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseplanreferenceDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Caseplanreference
type Caseplanreference struct { 
    


    // Name
    Name string `json:"name"`


    // Version - The version of the Caseplan.
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Caseplanreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseplanreference) MarshalJSON() ([]byte, error) {
    type Alias Caseplanreference

    if CaseplanreferenceMarshalled {
        return []byte("{}"), nil
    }
    CaseplanreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

