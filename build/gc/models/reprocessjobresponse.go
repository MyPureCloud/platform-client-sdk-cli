package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReprocessjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReprocessjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Reprocessjobresponse
type Reprocessjobresponse struct { 
    


    // Name
    Name string `json:"name"`


    // Description - The description of the job.
    Description string `json:"description"`


    // DateStart - The date from which the reprocessing should begin. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The date at which the reprocessing should end. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`


    // MediaTypes - Media types used to filter interactions.
    MediaTypes []string `json:"mediaTypes"`


    // Programs - The mapped programs to be included in the job.
    Programs []string `json:"programs"`


    // Dialects - Language dialects used to filter interactions.
    Dialects []string `json:"dialects"`


    // CreatedBy - The user who created the job.
    CreatedBy Addressableentityref `json:"createdBy"`


    // DateCreated - The date the job was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // JobStatus - The status of the job.
    JobStatus string `json:"jobStatus"`


    // QueueOrder - The position of the job in the queued jobs.
    QueueOrder int `json:"queueOrder"`


    // ProcessedInteractionsCount - The amount of interactions successfully reprocessed.
    ProcessedInteractionsCount int `json:"processedInteractionsCount"`


    // FailedInteractionsCount - The amount of failed interactions.
    FailedInteractionsCount int `json:"failedInteractionsCount"`


    // TotalInteractionsCount - The amount of interactions in the job.
    TotalInteractionsCount int `json:"totalInteractionsCount"`


    

}

// String returns a JSON representation of the model
func (o *Reprocessjobresponse) String() string {
    
    
    
    
     o.MediaTypes = []string{""} 
     o.Programs = []string{""} 
     o.Dialects = []string{""} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reprocessjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Reprocessjobresponse

    if ReprocessjobresponseMarshalled {
        return []byte("{}"), nil
    }
    ReprocessjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        MediaTypes []string `json:"mediaTypes"`
        
        Programs []string `json:"programs"`
        
        Dialects []string `json:"dialects"`
        
        CreatedBy Addressableentityref `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        JobStatus string `json:"jobStatus"`
        
        QueueOrder int `json:"queueOrder"`
        
        ProcessedInteractionsCount int `json:"processedInteractionsCount"`
        
        FailedInteractionsCount int `json:"failedInteractionsCount"`
        
        TotalInteractionsCount int `json:"totalInteractionsCount"`
        *Alias
    }{

        


        


        


        


        


        
        MediaTypes: []string{""},
        


        
        Programs: []string{""},
        


        
        Dialects: []string{""},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

