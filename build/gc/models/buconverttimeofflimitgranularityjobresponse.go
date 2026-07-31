package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuconverttimeofflimitgranularityjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuconverttimeofflimitgranularityjobresponseDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Buconverttimeofflimitgranularityjobresponse
type Buconverttimeofflimitgranularityjobresponse struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // TimeOffLimit - The time-off limit associated with this job
    TimeOffLimit Butimeofflimitreference `json:"timeOffLimit"`


    // Status - The status of the job
    Status string `json:"status"`


    // Progress - Progress of time-off limit granularity conversion
    Progress Buconverttimeofflimitgranularityjobprogress `json:"progress"`


    // VarError - Error information. Set only when status is Error
    VarError Errorbody `json:"error"`


    

}

// String returns a JSON representation of the model
func (o *Buconverttimeofflimitgranularityjobresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buconverttimeofflimitgranularityjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Buconverttimeofflimitgranularityjobresponse

    if BuconverttimeofflimitgranularityjobresponseMarshalled {
        return []byte("{}"), nil
    }
    BuconverttimeofflimitgranularityjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        TimeOffLimit Butimeofflimitreference `json:"timeOffLimit"`
        
        Status string `json:"status"`
        
        Progress Buconverttimeofflimitgranularityjobprogress `json:"progress"`
        
        VarError Errorbody `json:"error"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

