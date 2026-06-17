package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RegisterarchitectvalidatejobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RegisterarchitectvalidatejobDud struct { 
    

}

// Registerarchitectvalidatejob
type Registerarchitectvalidatejob struct { 
    // Flows - A list of the flows to be validated.
    Flows []Validatedetails `json:"flows"`

}

// String returns a JSON representation of the model
func (o *Registerarchitectvalidatejob) String() string {
     o.Flows = []Validatedetails{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Registerarchitectvalidatejob) MarshalJSON() ([]byte, error) {
    type Alias Registerarchitectvalidatejob

    if RegisterarchitectvalidatejobMarshalled {
        return []byte("{}"), nil
    }
    RegisterarchitectvalidatejobMarshalled = true

    return json.Marshal(&struct {
        
        Flows []Validatedetails `json:"flows"`
        *Alias
    }{

        
        Flows: []Validatedetails{{}},
        

        Alias: (*Alias)(u),
    })
}

