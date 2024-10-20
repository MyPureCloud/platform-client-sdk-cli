package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytablerowblockwithhighlightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytablerowblockwithhighlightDud struct { 
    


    

}

// Documentbodytablerowblockwithhighlight
type Documentbodytablerowblockwithhighlight struct { 
    // Properties - The properties for the table rows.
    Properties Documentbodytablerowblockproperties `json:"properties"`


    // Cells - The list of cells for the table.
    Cells []Documentbodytablecellblockwithhighlight `json:"cells"`

}

// String returns a JSON representation of the model
func (o *Documentbodytablerowblockwithhighlight) String() string {
    
     o.Cells = []Documentbodytablecellblockwithhighlight{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytablerowblockwithhighlight) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytablerowblockwithhighlight

    if DocumentbodytablerowblockwithhighlightMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytablerowblockwithhighlightMarshalled = true

    return json.Marshal(&struct {
        
        Properties Documentbodytablerowblockproperties `json:"properties"`
        
        Cells []Documentbodytablecellblockwithhighlight `json:"cells"`
        *Alias
    }{

        


        
        Cells: []Documentbodytablecellblockwithhighlight{{}},
        

        Alias: (*Alias)(u),
    })
}

