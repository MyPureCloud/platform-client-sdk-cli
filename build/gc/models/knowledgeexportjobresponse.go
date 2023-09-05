package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeexportjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeexportjobresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgeexportjobresponse
type Knowledgeexportjobresponse struct { 
    // Id - Id of the export job.
    Id string `json:"id"`


    // DownloadURL - The URL of the location at which the caller can download the export file, when available.
    DownloadURL string `json:"downloadURL"`


    // FileType - File type of the document
    FileType string `json:"fileType"`


    // CountDocumentProcessed - The current count of the number of records processed.
    CountDocumentProcessed int `json:"countDocumentProcessed"`


    // ExportFilter - Filters to narrow down what to export.
    ExportFilter Knowledgeexportjobfilter `json:"exportFilter"`


    // Status - The status of the export job.
    Status string `json:"status"`


    // KnowledgeBase - Knowledge base which document export belongs to.
    KnowledgeBase Knowledgebase `json:"knowledgeBase"`


    // CreatedBy - The user who created the operation
    CreatedBy Userreference `json:"createdBy"`


    // DateCreated - The timestamp of when the export began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The timestamp of when the export stopped. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ErrorInformation - Any error information, or null of the processing is not in failed state.
    ErrorInformation Errorbody `json:"errorInformation"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgeexportjobresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeexportjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeexportjobresponse

    if KnowledgeexportjobresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeexportjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DownloadURL string `json:"downloadURL"`
        
        FileType string `json:"fileType"`
        
        CountDocumentProcessed int `json:"countDocumentProcessed"`
        
        ExportFilter Knowledgeexportjobfilter `json:"exportFilter"`
        
        Status string `json:"status"`
        
        KnowledgeBase Knowledgebase `json:"knowledgeBase"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ErrorInformation Errorbody `json:"errorInformation"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

