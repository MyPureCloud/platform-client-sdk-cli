package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableexportjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableexportjobDud struct { 
    Id string `json:"id"`


    TableVersion int `json:"tableVersion"`


    Status string `json:"status"`


    CreatedBy Addressableentityref `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    FileName string `json:"fileName"`


    Download Addressableentityref `json:"download"`


    DateDownloadExpires time.Time `json:"dateDownloadExpires"`


    ExportType string `json:"exportType"`


    TotalRows int `json:"totalRows"`


    RowsExported int `json:"rowsExported"`


    Format string `json:"format"`


    VarError Decisiontableexportjoberror `json:"error"`


    SelfUri string `json:"selfUri"`

}

// Decisiontableexportjob
type Decisiontableexportjob struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Decisiontableexportjob) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableexportjob) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableexportjob

    if DecisiontableexportjobMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableexportjobMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

