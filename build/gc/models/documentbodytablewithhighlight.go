package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentbodytablewithhighlightMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentbodytablewithhighlightDud struct { 
    


    

}

// Documentbodytablewithhighlight
type Documentbodytablewithhighlight struct { 
    // Properties - The properties for the table.
    Properties Documentbodytableproperties `json:"properties"`


    // Rows - The list of rows for the table.
    Rows []Documentbodytablerowblockwithhighlight `json:"rows"`

}

// String returns a JSON representation of the model
func (o *Documentbodytablewithhighlight) String() string {
    
     o.Rows = []Documentbodytablerowblockwithhighlight{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentbodytablewithhighlight) MarshalJSON() ([]byte, error) {
    type Alias Documentbodytablewithhighlight

    if DocumentbodytablewithhighlightMarshalled {
        return []byte("{}"), nil
    }
    DocumentbodytablewithhighlightMarshalled = true

    return json.Marshal(&struct {
        
        Properties Documentbodytableproperties `json:"properties"`
        
        Rows []Documentbodytablerowblockwithhighlight `json:"rows"`
        *Alias
    }{

        


        
        Rows: []Documentbodytablerowblockwithhighlight{{}},
        

        Alias: (*Alias)(u),
    })
}

