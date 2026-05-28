package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryopportunityenrollmentmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryopportunityenrollmentmetadataDud struct { 
    ModifiedBy Userreference `json:"modifiedBy"`


    DateModified time.Time `json:"dateModified"`


    CreatedBy Userreference `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    


    


    

}

// Queryopportunityenrollmentmetadata
type Queryopportunityenrollmentmetadata struct { 
    


    


    


    


    // Version - The version of the associated entity.  Used to prevent conflicts on concurrent edits
    Version int `json:"version"`


    // ReviewedBy - The user who reviewed the enrollment
    ReviewedBy Userreference `json:"reviewedBy"`


    // DateReviewed - The date the enrollment was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateReviewed time.Time `json:"dateReviewed"`

}

// String returns a JSON representation of the model
func (o *Queryopportunityenrollmentmetadata) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryopportunityenrollmentmetadata) MarshalJSON() ([]byte, error) {
    type Alias Queryopportunityenrollmentmetadata

    if QueryopportunityenrollmentmetadataMarshalled {
        return []byte("{}"), nil
    }
    QueryopportunityenrollmentmetadataMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        ReviewedBy Userreference `json:"reviewedBy"`
        
        DateReviewed time.Time `json:"dateReviewed"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

