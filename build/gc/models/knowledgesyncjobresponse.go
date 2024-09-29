package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesyncjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesyncjobresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgesyncjobresponse
type Knowledgesyncjobresponse struct { 
    // Id - Id of the sync job.
    Id string `json:"id"`


    // UploadKey
    UploadKey string `json:"uploadKey"`


    // Status - The status of the export job.
    Status string `json:"status"`


    // Report - Report of the sync job
    Report Knowledgesyncjobreport `json:"report"`


    // KnowledgeBase - Knowledge base which document export belongs to.
    KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`


    // DateCreated - The timestamp of when the export began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The timestamp of when the export stopped. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // CreatedBy - The user who created the operation
    CreatedBy Userreference `json:"createdBy"`


    // DownloadURL - The URL of the location at which the caller can download the sync file, when available.
    DownloadURL string `json:"downloadURL"`


    // FailedEntitiesURL - The URL of the location at which the caller can download the entities in json format that failed during the sync.
    FailedEntitiesURL string `json:"failedEntitiesURL"`


    // Source - Source of the sync job.
    Source Knowledgeoperationsource `json:"source"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgesyncjobresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesyncjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesyncjobresponse

    if KnowledgesyncjobresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesyncjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        UploadKey string `json:"uploadKey"`
        
        Status string `json:"status"`
        
        Report Knowledgesyncjobreport `json:"report"`
        
        KnowledgeBase Knowledgebasereference `json:"knowledgeBase"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        DownloadURL string `json:"downloadURL"`
        
        FailedEntitiesURL string `json:"failedEntitiesURL"`
        
        Source Knowledgeoperationsource `json:"source"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

