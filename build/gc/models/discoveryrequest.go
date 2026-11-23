package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DiscoveryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DiscoveryrequestDud struct { 
    


    


    


    


    

}

// Discoveryrequest
type Discoveryrequest struct { 
    // Language - The language code for the discovery job.
    Language string `json:"language"`


    // DateStart - The start date for the discovery job in ISO 8601 format.
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The end date for the discovery job in ISO 8601 format.
    DateEnd time.Time `json:"dateEnd"`


    // QueueId - The queue ID to include in the discovery job.
    QueueId string `json:"queueId"`


    // MediaTypes - The media types to include in the discovery job.
    MediaTypes []string `json:"mediaTypes"`

}

// String returns a JSON representation of the model
func (o *Discoveryrequest) String() string {
    
    
    
    
     o.MediaTypes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Discoveryrequest) MarshalJSON() ([]byte, error) {
    type Alias Discoveryrequest

    if DiscoveryrequestMarshalled {
        return []byte("{}"), nil
    }
    DiscoveryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Language string `json:"language"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        QueueId string `json:"queueId"`
        
        MediaTypes []string `json:"mediaTypes"`
        *Alias
    }{

        


        


        


        


        
        MediaTypes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

