package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontablerowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontablerowDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Decisiontablerow
type Decisiontablerow struct { 
    


    // Table - The decision table to which this row belongs
    Table Decisiontableversionentity `json:"table"`


    // RowIndex - The absolute index of this row in the decision table, starting at 0
    RowIndex int `json:"rowIndex"`


    // DateCreated - The date when this row was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date when this row was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // Inputs - The map input values of the row being created. At least one value must be provided. The key for this map is the column ID the row value relates
    Inputs map[string]Decisiontablerowparametervalue `json:"inputs"`


    // Outputs - The map output values of the row being created. At least one value must be provided. The key for this map is the column ID the row value relates
    Outputs map[string]Decisiontablerowparametervalue `json:"outputs"`


    

}

// String returns a JSON representation of the model
func (o *Decisiontablerow) String() string {
    
    
    
    
     o.Inputs = map[string]Decisiontablerowparametervalue{"": {}} 
     o.Outputs = map[string]Decisiontablerowparametervalue{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontablerow) MarshalJSON() ([]byte, error) {
    type Alias Decisiontablerow

    if DecisiontablerowMarshalled {
        return []byte("{}"), nil
    }
    DecisiontablerowMarshalled = true

    return json.Marshal(&struct {
        
        Table Decisiontableversionentity `json:"table"`
        
        RowIndex int `json:"rowIndex"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        Inputs map[string]Decisiontablerowparametervalue `json:"inputs"`
        
        Outputs map[string]Decisiontablerowparametervalue `json:"outputs"`
        *Alias
    }{

        


        


        


        


        


        
        Inputs: map[string]Decisiontablerowparametervalue{"": {}},
        


        
        Outputs: map[string]Decisiontablerowparametervalue{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

