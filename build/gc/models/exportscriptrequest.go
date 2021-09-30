package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExportscriptrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExportscriptrequestDud struct { 
    


    

}

// Exportscriptrequest - Creating an exported script via Download Service
type Exportscriptrequest struct { 
    // FileName - The final file name (no extension) of the script download: <fileName>.script
    FileName string `json:"fileName"`


    // VersionId - The UUID version of the script to be exported.  Defaults to the current editable version.
    VersionId string `json:"versionId"`

}

// String returns a JSON representation of the model
func (o *Exportscriptrequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Exportscriptrequest) MarshalJSON() ([]byte, error) {
    type Alias Exportscriptrequest

    if ExportscriptrequestMarshalled {
        return []byte("{}"), nil
    }
    ExportscriptrequestMarshalled = true

    return json.Marshal(&struct { 
        FileName string `json:"fileName"`
        
        VersionId string `json:"versionId"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

