package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerycriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerycriteriaDud struct { 
    


    


    


    


    


    

}

// Querycriteria - A criteria type that can be used in tandem with other criteria type to create queries of executionData
type Querycriteria struct { 
    // CriteriaKey - The is the name of the criteria that can be queried.
    CriteriaKey string `json:"criteriaKey"`


    // CriteriaGroups - The executionData type that this criteria item can be used on.
    CriteriaGroups []string `json:"criteriaGroups"`


    // Description - The is the description of the criteria.
    Description string `json:"description"`


    // Operators - A list of operators that can be used on this criteria.
    Operators []string `json:"operators"`


    // DataType - The type of data for the criteria (string, int, etc).
    DataType string `json:"dataType"`


    // CategoryInfo - A logical grouping and display order for this item.
    CategoryInfo Criteriacategoryinfo `json:"categoryInfo"`

}

// String returns a JSON representation of the model
func (o *Querycriteria) String() string {
    
     o.CriteriaGroups = []string{""} 
    
     o.Operators = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querycriteria) MarshalJSON() ([]byte, error) {
    type Alias Querycriteria

    if QuerycriteriaMarshalled {
        return []byte("{}"), nil
    }
    QuerycriteriaMarshalled = true

    return json.Marshal(&struct {
        
        CriteriaKey string `json:"criteriaKey"`
        
        CriteriaGroups []string `json:"criteriaGroups"`
        
        Description string `json:"description"`
        
        Operators []string `json:"operators"`
        
        DataType string `json:"dataType"`
        
        CategoryInfo Criteriacategoryinfo `json:"categoryInfo"`
        *Alias
    }{

        


        
        CriteriaGroups: []string{""},
        


        


        
        Operators: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

