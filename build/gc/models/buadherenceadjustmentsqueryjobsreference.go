package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuadherenceadjustmentsqueryjobsreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuadherenceadjustmentsqueryjobsreferenceDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Buadherenceadjustmentsqueryjobsreference
type Buadherenceadjustmentsqueryjobsreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Status - The status of the adherence adjustments query job
    Status string `json:"status"`


    // CreatedDate - The date the query job was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    

}

// String returns a JSON representation of the model
func (o *Buadherenceadjustmentsqueryjobsreference) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buadherenceadjustmentsqueryjobsreference) MarshalJSON() ([]byte, error) {
    type Alias Buadherenceadjustmentsqueryjobsreference

    if BuadherenceadjustmentsqueryjobsreferenceMarshalled {
        return []byte("{}"), nil
    }
    BuadherenceadjustmentsqueryjobsreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

