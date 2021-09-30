package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinerexecuterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinerexecuterequestDud struct { 
    


    


    


    


    

}

// Minerexecuterequest
type Minerexecuterequest struct { 
    // DateStart - Start date for the date range to mine. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`


    // DateEnd - End date for the date range to mine. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEnd time.Time `json:"dateEnd"`


    // UploadKey - Location of input conversations.
    UploadKey string `json:"uploadKey"`


    // MediaType - Media type for filtering conversations.
    MediaType string `json:"mediaType"`


    // QueueIds - List of queue IDs for filtering conversations.
    QueueIds []string `json:"queueIds"`

}

// String returns a JSON representation of the model
func (o *Minerexecuterequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.QueueIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Minerexecuterequest) MarshalJSON() ([]byte, error) {
    type Alias Minerexecuterequest

    if MinerexecuterequestMarshalled {
        return []byte("{}"), nil
    }
    MinerexecuterequestMarshalled = true

    return json.Marshal(&struct { 
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        UploadKey string `json:"uploadKey"`
        
        MediaType string `json:"mediaType"`
        
        QueueIds []string `json:"queueIds"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        QueueIds: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

