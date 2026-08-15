package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CsvexporterrordetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CsvexporterrordetailsDud struct { 
    


    

}

// Csvexporterrordetails
type Csvexporterrordetails struct { 
    // ErrorCode - The error code of a failed export
    ErrorCode string `json:"errorCode"`


    // ErrorMessage - The error message of a failed export
    ErrorMessage string `json:"errorMessage"`

}

// String returns a JSON representation of the model
func (o *Csvexporterrordetails) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Csvexporterrordetails) MarshalJSON() ([]byte, error) {
    type Alias Csvexporterrordetails

    if CsvexporterrordetailsMarshalled {
        return []byte("{}"), nil
    }
    CsvexporterrordetailsMarshalled = true

    return json.Marshal(&struct {
        
        ErrorCode string `json:"errorCode"`
        
        ErrorMessage string `json:"errorMessage"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

