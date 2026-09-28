package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuadherenceadjustmentsqueryjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuadherenceadjustmentsqueryjobDud struct { 
    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Buadherenceadjustmentsqueryjob
type Buadherenceadjustmentsqueryjob struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Status - The status of the adherence adjustments query job
    Status string `json:"status"`


    // DownloadUrl - A URL to fetch results of the job. Only set if status == 'Complete'
    DownloadUrl string `json:"downloadUrl"`


    // VarError - Error details if status == 'Error'
    VarError Errorbody `json:"error"`


    // Result - Schema template for deserializing data returned from the downloadUrl. Will always be null on the response
    Result Adherenceadjustmentslisting `json:"result"`


    

}

// String returns a JSON representation of the model
func (o *Buadherenceadjustmentsqueryjob) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buadherenceadjustmentsqueryjob) MarshalJSON() ([]byte, error) {
    type Alias Buadherenceadjustmentsqueryjob

    if BuadherenceadjustmentsqueryjobMarshalled {
        return []byte("{}"), nil
    }
    BuadherenceadjustmentsqueryjobMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        VarError Errorbody `json:"error"`
        
        Result Adherenceadjustmentslisting `json:"result"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

