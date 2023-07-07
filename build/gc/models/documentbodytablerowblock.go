package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytablerowblockMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytablerowblockDud struct { 
    


    

}

// Documentbodytablerowblock
type Documentbodytablerowblock struct { 
    // Cells - The list of cells for the table.
    Cells []Documentbodytablecellblock `json:"cells"`


    // Properties - The properties for the table rows.
    Properties Documentbodytablerowblockproperties `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Documentbodytablerowblock) String() string {
     o.Cells = []Documentbodytablecellblock{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytablerowblock) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytablerowblock

    if DocumentbodytablerowblockMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytablerowblockMarshalled = true

    return json.Marshal(&struct {
        
        Cells []Documentbodytablecellblock `json:"cells"`
        
        Properties Documentbodytablerowblockproperties `json:"properties"`
        *Alias
    }{

        
        Cells: []Documentbodytablecellblock{{}},
        


        

        Alias: (*Alias)(u),
    })
}

