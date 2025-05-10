package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CsvjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CsvjobrequestDud struct { 
    


    


    

}

// Csvjobrequest
type Csvjobrequest struct { 
    // UploadId - Upload for the csv job
    UploadId string `json:"uploadId"`


    // SettingsId - Settings for the csv job
    SettingsId string `json:"settingsId"`


    // Division - Division for the csv job
    Division Writablestarrabledivision `json:"division"`

}

// String returns a JSON representation of the model
func (o *Csvjobrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Csvjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Csvjobrequest

    if CsvjobrequestMarshalled {
        return []byte("{}"), nil
    }
    CsvjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        UploadId string `json:"uploadId"`
        
        SettingsId string `json:"settingsId"`
        
        Division Writablestarrabledivision `json:"division"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

