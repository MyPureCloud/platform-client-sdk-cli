package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatereprocessjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatereprocessjobrequestDud struct { 
    


    


    


    


    


    


    

}

// Createreprocessjobrequest
type Createreprocessjobrequest struct { 
    // Name - The name given for the job.
    Name string `json:"name"`


    // Description - The description of the job.
    Description string `json:"description"`


    // DateStart - The start date for the search criteria. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The end date for the search criteria. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`


    // Programs - A list of program IDs to filter conversations by.
    Programs []string `json:"programs"`


    // MediaTypes - The types of conversations to reprocess.
    MediaTypes []string `json:"mediaTypes"`


    // Dialects - The dialects to filter by the conversations.
    Dialects []string `json:"dialects"`

}

// String returns a JSON representation of the model
func (o *Createreprocessjobrequest) String() string {
    
    
    
    
     o.Programs = []string{""} 
     o.MediaTypes = []string{""} 
     o.Dialects = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createreprocessjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Createreprocessjobrequest

    if CreatereprocessjobrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatereprocessjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        Programs []string `json:"programs"`
        
        MediaTypes []string `json:"mediaTypes"`
        
        Dialects []string `json:"dialects"`
        *Alias
    }{

        


        


        


        


        
        Programs: []string{""},
        


        
        MediaTypes: []string{""},
        


        
        Dialects: []string{""},
        

        Alias: (*Alias)(u),
    })
}

