package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TerminatejobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TerminatejobDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Terminatejob
type Terminatejob struct { 
    


    // Status - The status of the Job.
    Status string `json:"status"`


    // JobType - The type of the Job.
    JobType string `json:"jobType"`


    // CreatedBy - The ID of the User who created this Job.
    CreatedBy Userreference `json:"createdBy"`


    // DateCreated - The job creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The job modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Terminatejob) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Terminatejob) MarshalJSON() ([]byte, error) {
    type Alias Terminatejob

    if TerminatejobMarshalled {
        return []byte("{}"), nil
    }
    TerminatejobMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        JobType string `json:"jobType"`
        
        CreatedBy Userreference `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

