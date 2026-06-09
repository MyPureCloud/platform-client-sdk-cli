package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableexportjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableexportjobrequestDud struct { 
    


    


    

}

// Decisiontableexportjobrequest
type Decisiontableexportjobrequest struct { 
    // TableVersion
    TableVersion int `json:"tableVersion"`


    // ExportType - The type of export to perform.
    ExportType string `json:"exportType"`


    // Format - The format of the exported file.
    Format string `json:"format"`

}

// String returns a JSON representation of the model
func (o *Decisiontableexportjobrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableexportjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableexportjobrequest

    if DecisiontableexportjobrequestMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableexportjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        TableVersion int `json:"tableVersion"`
        
        ExportType string `json:"exportType"`
        
        Format string `json:"format"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

