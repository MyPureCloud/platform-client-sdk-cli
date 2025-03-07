package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExportdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExportdetailsDud struct { 
    


    


    

}

// Exportdetails
type Exportdetails struct { 
    // Flow - The flow to export. If you do not provide the flow ID, you must provide both the name and type.
    Flow Architectflowreference `json:"flow"`


    // FileName - Name to assign to the file after download. The extension will be automatically appended based on desired export type. Must contain only alphanumeric characters, underscores, or hyphens.
    FileName string `json:"fileName"`


    // ExportType - The export type for the flow. Default: 'Yaml'.
    ExportType string `json:"exportType"`

}

// String returns a JSON representation of the model
func (o *Exportdetails) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Exportdetails) MarshalJSON() ([]byte, error) {
    type Alias Exportdetails

    if ExportdetailsMarshalled {
        return []byte("{}"), nil
    }
    ExportdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Flow Architectflowreference `json:"flow"`
        
        FileName string `json:"fileName"`
        
        ExportType string `json:"exportType"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

