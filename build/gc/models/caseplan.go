package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseplanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseplanDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Caseplan
type Caseplan struct { 
    


    // Name - The name of the Caseplan.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Starrabledivision `json:"division"`


    // Description - The description of the Caseplan.
    Description string `json:"description"`


    // ReferencePrefix - The prefix used when creating the reference for Cases from the Caseplan.
    ReferencePrefix string `json:"referencePrefix"`


    // DefaultDueDurationInSeconds - The default due duration in seconds for Cases created from the Caseplan.
    DefaultDueDurationInSeconds int `json:"defaultDueDurationInSeconds"`


    // DefaultTtlSeconds - The default TTL in seconds for Cases created from the Caseplan.
    DefaultTtlSeconds int `json:"defaultTtlSeconds"`


    // DefaultCaseOwner - The default case owner for Cases created from the Caseplan.
    DefaultCaseOwner Userreference `json:"defaultCaseOwner"`


    // Latest - The latest version of the Caseplan.
    Latest int `json:"latest"`


    // Published - The published version of the Caseplan.
    Published int `json:"published"`


    // DateCreated - The Caseplan creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The Caseplan modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DatePublished - The Caseplan publication date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DatePublished time.Time `json:"datePublished"`


    // ModifiedBy - The id of the User who modified the Caseplan.
    ModifiedBy Userreference `json:"modifiedBy"`


    // CustomerIntent - The customer intent for the Cases created from the caseplan.
    CustomerIntent Customerintentreference `json:"customerIntent"`


    // VersionState - The version state of the Caseplan.
    VersionState string `json:"versionState"`


    

}

// String returns a JSON representation of the model
func (o *Caseplan) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseplan) MarshalJSON() ([]byte, error) {
    type Alias Caseplan

    if CaseplanMarshalled {
        return []byte("{}"), nil
    }
    CaseplanMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Starrabledivision `json:"division"`
        
        Description string `json:"description"`
        
        ReferencePrefix string `json:"referencePrefix"`
        
        DefaultDueDurationInSeconds int `json:"defaultDueDurationInSeconds"`
        
        DefaultTtlSeconds int `json:"defaultTtlSeconds"`
        
        DefaultCaseOwner Userreference `json:"defaultCaseOwner"`
        
        Latest int `json:"latest"`
        
        Published int `json:"published"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        DatePublished time.Time `json:"datePublished"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        CustomerIntent Customerintentreference `json:"customerIntent"`
        
        VersionState string `json:"versionState"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

