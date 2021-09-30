package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidationservicerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidationservicerequestDud struct { 
    


    

}

// Validationservicerequest
type Validationservicerequest struct { 
    // DateImportEnded - The last day of the data you are importing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateImportEnded time.Time `json:"dateImportEnded"`


    // FileUrl - Path to the file in the storage including the file name
    FileUrl string `json:"fileUrl"`

}

// String returns a JSON representation of the model
func (o *Validationservicerequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validationservicerequest) MarshalJSON() ([]byte, error) {
    type Alias Validationservicerequest

    if ValidationservicerequestMarshalled {
        return []byte("{}"), nil
    }
    ValidationservicerequestMarshalled = true

    return json.Marshal(&struct { 
        DateImportEnded time.Time `json:"dateImportEnded"`
        
        FileUrl string `json:"fileUrl"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

