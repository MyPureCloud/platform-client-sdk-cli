package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ImportreportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ImportreportDud struct { 
    


    


    


    

}

// Importreport
type Importreport struct { 
    // Errors
    Errors []Importerror `json:"errors"`


    // Validated
    Validated Resultcounters `json:"validated"`


    // Imported
    Imported Resultcounters `json:"imported"`


    // TotalDocuments
    TotalDocuments int `json:"totalDocuments"`

}

// String returns a JSON representation of the model
func (o *Importreport) String() string {
    
    
     o.Errors = []Importerror{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Importreport) MarshalJSON() ([]byte, error) {
    type Alias Importreport

    if ImportreportMarshalled {
        return []byte("{}"), nil
    }
    ImportreportMarshalled = true

    return json.Marshal(&struct { 
        Errors []Importerror `json:"errors"`
        
        Validated Resultcounters `json:"validated"`
        
        Imported Resultcounters `json:"imported"`
        
        TotalDocuments int `json:"totalDocuments"`
        
        *Alias
    }{
        

        
        Errors: []Importerror{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

