package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeversionreportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeversionreportDud struct { 
    


    

}

// Edgeversionreport
type Edgeversionreport struct { 
    // OldestVersion
    OldestVersion Edgeversioninformation `json:"oldestVersion"`


    // NewestVersion
    NewestVersion Edgeversioninformation `json:"newestVersion"`

}

// String returns a JSON representation of the model
func (o *Edgeversionreport) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeversionreport) MarshalJSON() ([]byte, error) {
    type Alias Edgeversionreport

    if EdgeversionreportMarshalled {
        return []byte("{}"), nil
    }
    EdgeversionreportMarshalled = true

    return json.Marshal(&struct { 
        OldestVersion Edgeversioninformation `json:"oldestVersion"`
        
        NewestVersion Edgeversioninformation `json:"newestVersion"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

