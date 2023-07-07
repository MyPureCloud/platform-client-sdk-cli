package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytableMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytableDud struct { 
    


    

}

// Documentbodytable
type Documentbodytable struct { 
    // Rows - The list of rows for the table.
    Rows []Documentbodytablerowblock `json:"rows"`


    // Properties - The properties for the table.
    Properties Documentbodytableproperties `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Documentbodytable) String() string {
     o.Rows = []Documentbodytablerowblock{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytable) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytable

    if DocumentbodytableMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytableMarshalled = true

    return json.Marshal(&struct {
        
        Rows []Documentbodytablerowblock `json:"rows"`
        
        Properties Documentbodytableproperties `json:"properties"`
        *Alias
    }{

        
        Rows: []Documentbodytablerowblock{{}},
        


        

        Alias: (*Alias)(u),
    })
}

