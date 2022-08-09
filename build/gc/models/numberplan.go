package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NumberplanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NumberplanDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Numberplan
type Numberplan struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // Match
    Match string `json:"match"`


    // NormalizedFormat
    NormalizedFormat string `json:"normalizedFormat"`


    // Priority
    Priority int `json:"priority"`


    // Numbers
    Numbers []Number `json:"numbers"`


    // DigitLength
    DigitLength Digitlength `json:"digitLength"`


    // Classification
    Classification string `json:"classification"`


    // MatchType
    MatchType string `json:"matchType"`


    

}

// String returns a JSON representation of the model
func (o *Numberplan) String() string {
    
    
    
    
    
    
    
     o.Numbers = []Number{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Numberplan) MarshalJSON() ([]byte, error) {
    type Alias Numberplan

    if NumberplanMarshalled {
        return []byte("{}"), nil
    }
    NumberplanMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        Match string `json:"match"`
        
        NormalizedFormat string `json:"normalizedFormat"`
        
        Priority int `json:"priority"`
        
        Numbers []Number `json:"numbers"`
        
        DigitLength Digitlength `json:"digitLength"`
        
        Classification string `json:"classification"`
        
        MatchType string `json:"matchType"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Numbers: []Number{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

