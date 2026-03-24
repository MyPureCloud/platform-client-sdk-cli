package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Case
type Case struct { 
    


    // Name - The name of the Case.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Starrabledivision `json:"division"`


    // Version - The version of the Case.
    Version int `json:"version"`


    // Reference - The reference identifier of the Case.
    Reference string `json:"reference"`


    // Caseplan - The Caseplan the case was created from.
    Caseplan Caseplanreference `json:"caseplan"`


    // Summary - Overview information for the Case.
    Summary string `json:"summary"`


    // Owner - The owner of the Case.
    Owner Userreference `json:"owner"`


    // Status - The status of the Case.
    Status string `json:"status"`


    // Priority - The priority of the Case.
    Priority string `json:"priority"`


    // DateDue - The due date of the Case. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateDue time.Time `json:"dateDue"`


    // DateStarted - The start time of the Case. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStarted time.Time `json:"dateStarted"`


    // DateClosed - The completion time of the Case. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateClosed time.Time `json:"dateClosed"`


    // DateCreated - The date the Case was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date the Case was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The id of the User who modified the Case.
    ModifiedBy Userreference `json:"modifiedBy"`


    // ExternalContact - The External Contact associated with the Case.
    ExternalContact Caseexternalcontactreference `json:"externalContact"`


    // CustomerIntent - The customer intent for the Case.
    CustomerIntent Customerintentreference `json:"customerIntent"`


    // CreationStatus - The creation status of the Case
    CreationStatus string `json:"creationStatus"`


    // TtlSeconds - The time-to-live in seconds for the lifetime of the Case.
    TtlSeconds int `json:"ttlSeconds"`


    

}

// String returns a JSON representation of the model
func (o *Case) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Case) MarshalJSON() ([]byte, error) {
    type Alias Case

    if CaseMarshalled {
        return []byte("{}"), nil
    }
    CaseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Starrabledivision `json:"division"`
        
        Version int `json:"version"`
        
        Reference string `json:"reference"`
        
        Caseplan Caseplanreference `json:"caseplan"`
        
        Summary string `json:"summary"`
        
        Owner Userreference `json:"owner"`
        
        Status string `json:"status"`
        
        Priority string `json:"priority"`
        
        DateDue time.Time `json:"dateDue"`
        
        DateStarted time.Time `json:"dateStarted"`
        
        DateClosed time.Time `json:"dateClosed"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        ExternalContact Caseexternalcontactreference `json:"externalContact"`
        
        CustomerIntent Customerintentreference `json:"customerIntent"`
        
        CreationStatus string `json:"creationStatus"`
        
        TtlSeconds int `json:"ttlSeconds"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

