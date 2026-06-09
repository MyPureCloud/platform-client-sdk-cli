package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatedecisiontableimportjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatedecisiontableimportjobrequestDud struct { 
    


    

}

// Createdecisiontableimportjobrequest
type Createdecisiontableimportjobrequest struct { 
    // ImportMode - How imported rows are merged with existing rows
    ImportMode string `json:"importMode"`


    // FileName - Name of the file to import. Must include the file extension.
    FileName string `json:"fileName"`

}

// String returns a JSON representation of the model
func (o *Createdecisiontableimportjobrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createdecisiontableimportjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Createdecisiontableimportjobrequest

    if CreatedecisiontableimportjobrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatedecisiontableimportjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        ImportMode string `json:"importMode"`
        
        FileName string `json:"fileName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

