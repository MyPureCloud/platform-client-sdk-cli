package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntegrationexportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntegrationexportDud struct { 
    


    

}

// Integrationexport
type Integrationexport struct { 
    // Integration - The aws-s3-recording-bulk-actions-integration that the policy uses for exports.
    Integration Domainentityref `json:"integration"`


    // ShouldExportScreenRecordings - True if the policy should export screen recordings in addition to the other conversation media. Default = true
    ShouldExportScreenRecordings bool `json:"shouldExportScreenRecordings"`

}

// String returns a JSON representation of the model
func (o *Integrationexport) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Integrationexport) MarshalJSON() ([]byte, error) {
    type Alias Integrationexport

    if IntegrationexportMarshalled {
        return []byte("{}"), nil
    }
    IntegrationexportMarshalled = true

    return json.Marshal(&struct {
        
        Integration Domainentityref `json:"integration"`
        
        ShouldExportScreenRecordings bool `json:"shouldExportScreenRecordings"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

