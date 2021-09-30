package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExporturiMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExporturiDud struct { 
    


    

}

// Exporturi
type Exporturi struct { 
    // Uri
    Uri string `json:"uri"`


    // ExportTimestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ExportTimestamp time.Time `json:"exportTimestamp"`

}

// String returns a JSON representation of the model
func (o *Exporturi) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Exporturi) MarshalJSON() ([]byte, error) {
    type Alias Exporturi

    if ExporturiMarshalled {
        return []byte("{}"), nil
    }
    ExporturiMarshalled = true

    return json.Marshal(&struct { 
        Uri string `json:"uri"`
        
        ExportTimestamp time.Time `json:"exportTimestamp"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

