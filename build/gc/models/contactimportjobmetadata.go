package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactimportjobmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactimportjobmetadataDud struct { 
    


    


    


    


    


    


    

}

// Contactimportjobmetadata
type Contactimportjobmetadata struct { 
    // FileName
    FileName string `json:"fileName"`


    // DryRunFailedCount
    DryRunFailedCount int `json:"dryRunFailedCount"`


    // DryRunSuccessCount
    DryRunSuccessCount int `json:"dryRunSuccessCount"`


    // DryRunReportDownloadUrl
    DryRunReportDownloadUrl string `json:"dryRunReportDownloadUrl"`


    // ImportFailedCount
    ImportFailedCount int `json:"importFailedCount"`


    // ImportSuccessCount
    ImportSuccessCount int `json:"importSuccessCount"`


    // ImportReportDownloadUrl
    ImportReportDownloadUrl string `json:"importReportDownloadUrl"`

}

// String returns a JSON representation of the model
func (o *Contactimportjobmetadata) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactimportjobmetadata) MarshalJSON() ([]byte, error) {
    type Alias Contactimportjobmetadata

    if ContactimportjobmetadataMarshalled {
        return []byte("{}"), nil
    }
    ContactimportjobmetadataMarshalled = true

    return json.Marshal(&struct {
        
        FileName string `json:"fileName"`
        
        DryRunFailedCount int `json:"dryRunFailedCount"`
        
        DryRunSuccessCount int `json:"dryRunSuccessCount"`
        
        DryRunReportDownloadUrl string `json:"dryRunReportDownloadUrl"`
        
        ImportFailedCount int `json:"importFailedCount"`
        
        ImportSuccessCount int `json:"importSuccessCount"`
        
        ImportReportDownloadUrl string `json:"importReportDownloadUrl"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

