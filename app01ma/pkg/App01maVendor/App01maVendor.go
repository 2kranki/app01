// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Struct and Methods for App01maVendor

// Generated: Thu Nov 14, 2019 11:17


package App01maVendor

import (
	"encoding/json"
    "fmt"
    
	"net/http"
	"strconv"
	"strings"
     
	"net/url"

    
)

//============================================================================
//                             Database Interfaces
//============================================================================

type App01maVendorDbRowDeleter interface {
    // RowDelete deletes the row with keys from the provided record, rcd.
    RowDelete(rcd *App01maVendor) error
}

type App01maVendorDbRowFinder interface {
    // RowFind searches the Database for a matching row for the keys found in
    // the given record and returns the output in that same record.
    RowFind(rcd *App01maVendor) error
}

type App01maVendorDbRowFirster interface {
    // RowFirst returns the first row in the table, Vendor.
    // If there are no rows in the table, then a blank/null record is returned
    // without error.
    RowFirst(rcd *App01maVendor) error
}

type App01maVendorDbRowInserter interface {
    RowInsert(rcd *App01maVendor) error
}

type App01maVendorDbRowLaster interface {
    // RowLast returns the last row in the table, Vendor.
    // If there are no rows in the table, then a blank/null record is returned
    // without error.
    RowLast(rcd *App01maVendor) error
}

type App01maVendorDbRowNexter interface {
    // RowNext returns the next row from the row given. If row after the current
    // one does not exist, then the first row is returned.
    RowNext(rcd *App01maVendor) error
}

type App01maVendorDbRowPager interface {
    // RowPage returns a page of rows where a page size is the 'limit' parameter and
    // 'offset' is the offset into the result set ordered by the main index. Both
    // 'limit' and 'offset' are relative to 1. We return an address to the array
    // rows (structs) so that we don't have the overhead of copying them everwhere.
    RowPage(offset int, limit int) ([]App01maVendor, error)
}

type App01maVendorDbRowPrever interface {
    RowPrev(rcd *App01maVendor) error
}

type App01maVendorDbRowUpdater interface {
    RowUpdate(rcd *App01maVendor) error
}

type App01maVendorDbTableCounter interface {
    TableCount() (int, error)
}

type App01maVendorDbTableCreater interface {
    TableCreate() error
}

type App01maVendorDbTableDeleter interface {
    TableDelete() error
}

type App01maVendorDbTableScanner interface {
    // TableScan reads all the rows in the table applying a function to each of
    // them.
    TableScan(apply func (rcd App01maVendor) error) error
}

//============================================================================
//                              Table Struct
//============================================================================

type App01maVendor struct {
	Id	int64
	Name	string
	Addr1	string
	Addr2	string
	City	string
	State	string
	Zip	string
	Curbal	float64
}

type App01maVendors []*App01maVendor

type Key struct {
	Id	int64
}

type App01maVendorIndex map[Key]*App01maVendor



// NOTE: For JsonMarshal() and JsonUnmarshal() to work properly, the JSON
//  names must be defined above.

//----------------------------------------------------------------------------
//                              Compare
//----------------------------------------------------------------------------

// Compare compares our struct to another returning
// 0, 1 for equal and not equal.
func (s *App01maVendor) Compare(r *App01maVendor) int {
    // Accumulate the key value(s) in KeyNum order.
    if s.Id != r.Id {
            return 1
        }
	if s.Name != r.Name {
            return 1
        }
	if s.Addr1 != r.Addr1 {
            return 1
        }
	if s.Addr2 != r.Addr2 {
            return 1
        }
	if s.City != r.City {
            return 1
        }
	if s.State != r.State {
            return 1
        }
	if s.Zip != r.Zip {
            return 1
        }
	if s.Curbal != r.Curbal {
            return 1
        }
	return 0
}

// CompareKeys compares our struct to another using keys returning the normal
// -1, 0, 1 for less than, equal and greater than.
func (s *App01maVendor) CompareKeys(r *App01maVendor) int {
    // Accumulate the key value(s) in KeyNum order.
    // Field: Id
            if s.Id != r.Id {
                if s.Id < r.Id {
                    return -1
                } else {
                    return 1
                }
            }
	return 0
}

//----------------------------------------------------------------------------
//                             Empty
//----------------------------------------------------------------------------

// Empty resets the struct values to their null values.
func (s *App01maVendor) Empty() {
var i64     int64
var f64     float64
var str     string


s.Id = i64
    s.Name = str
    s.Addr1 = str
    s.Addr2 = str
    s.City = str
    s.State = str
    s.Zip = str
    s.Curbal = f64
    
}

//----------------------------------------------------------------------------
//                      Fields to URL Value String
//----------------------------------------------------------------------------

// FieldsToValue creates a URL Value map from the the table's field(s).
func (s *App01maVendor) FieldsToValue() string {
    var wrk string

    v := url.Values{}
    // Accumulate the value(s) from the fields.
    // Field: Id
            	wrk = fmt.Sprintf("%d", s.Id)
v.Add("Id", wrk)
	// Field: Name
            	wrk = s.Name
v.Add("Name", wrk)
	// Field: Addr1
            	wrk = s.Addr1
v.Add("Addr1", wrk)
	// Field: Addr2
            	wrk = s.Addr2
v.Add("Addr2", wrk)
	// Field: City
            	wrk = s.City
v.Add("City", wrk)
	// Field: State
            	wrk = s.State
v.Add("State", wrk)
	// Field: Zip
            	wrk = s.Zip
v.Add("Zip", wrk)
	// Field: Curbal
            	{
		s := fmt.Sprintf("%.4f", s.Curbal)
		wrk = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	}
v.Add("Curbal", wrk)
	return v.Encode()
}

//----------------------------------------------------------------------------
//                  		JSON Marshal
//----------------------------------------------------------------------------

func (d *App01maVendor) JsonMarshal() ([]byte, error) {
	var err         error
    var text        []byte

    if text, err = json.Marshal(d); err != nil {
		return nil, fmt.Errorf("Error: marshalling json: %s : %v", err, d)
	}

	return text, err
}

//----------------------------------------------------------------------------
//                             JSON Unmarshal
//----------------------------------------------------------------------------

func (d *App01maVendor) JsonUnmarshal(text []byte) error {
	var err         error

	if err = json.Unmarshal(text, d); err != nil {
		return fmt.Errorf("Error: unmarshalling json: %s : %s", err, text)
	}

	return err
}

//----------------------------------------------------------------------------
//                      Set Keys from a Slice of Strings
//----------------------------------------------------------------------------

// SetKeysFromStrings creates a URL Value map from the table's key(s). The slice
// is in field order within the struct, not sorted by field name.
func (s *App01maVendor) Key() Key {
    var k       Key

    k.Id = s.Id
	return k
}

//----------------------------------------------------------------------------
//                      Keys to URL Value String
//----------------------------------------------------------------------------

// KeysToValue creates a URL Value map from the table's key(s).
func (s *App01maVendor) KeysToValue() string {
    var wrk string

    v := url.Values{}
    // Accumulate the key value(s) in KeyNum order.
    // Field: Id
            	wrk = fmt.Sprintf("%d", s.Id)
v.Add(fmt.Sprintf("key%d", 1-1), wrk)
	return v.Encode()
}

//----------------------------------------------------------------------------
//                             List Output
//----------------------------------------------------------------------------

func (s *App01maVendor) ListOutput() string {
	var str strings.Builder
    var wrk string

    if s == nil {
        return ""
    }

    // Field: Id
            str.WriteString("<td>")
            wrk = fmt.Sprintf("<a href=\"/Vendor/find?%s\">", s.KeysToValue())
                str.WriteString(wrk)
            	wrk = fmt.Sprintf("%d", s.Id)
str.WriteString(wrk)
            //str.WriteString("\n")
            str.WriteString("</a>" )
            str.WriteString("</td>\n")
        // Field: Name
            str.WriteString("<td>")
            	wrk = s.Name
str.WriteString(wrk)
            //str.WriteString("\n")
            str.WriteString("</td>\n")
        return str.String()
}

//----------------------------------------------------------------------------
//                  Request Form Value(s) to Struct
//----------------------------------------------------------------------------

// VendorRequest2Struct converts the form values to a struct. FormValue(s) are available
// for both, GET and POST.  It is just that all your parameters are present in the URL if you use
// GET.  In general, you should use POST with this function for security reasons.
func (s *App01maVendor) Request2Struct(r *http.Request) error {
    var err         error
    var str         string

    

    s.Empty()
    str = r.FormValue("Id")
        	s.Id, _ = strconv.ParseInt(str,0,64)
str = r.FormValue("Name")
        	s.Name = str
str = r.FormValue("Addr1")
        	s.Addr1 = str
str = r.FormValue("Addr2")
        	s.Addr2 = str
str = r.FormValue("City")
        	s.City = str
str = r.FormValue("State")
        	s.State = str
str = r.FormValue("Zip")
        	s.Zip = str
str = r.FormValue("Curbal")
        		s.Curbal, _ = strconv.ParseFloat(str, 64)


    
    return err
}

//----------------------------------------------------------------------------
//                      Set Keys from a Slice of Strings
//----------------------------------------------------------------------------

// SetKeysFromStrings creates a URL Value map from the table's key(s). The slice
// is in field order within the struct, not sorted by field name.
func (s *App01maVendor) SetKeysFromStrings(strs []string) error {

    if len(strs) != 1 {
        return fmt.Errorf("Error - Invalid key count of %d, need %d!\n", len(strs), 1)
    }

    // Accumulate the key value(s) in KeyNum order.
    	s.Id, _ = strconv.ParseInt(strs[0],0,64)

	return nil
}

//----------------------------------------------------------------------------
//                             Test Data
//----------------------------------------------------------------------------

// TestData takes the given integer and uses it to fill most of the fields in
// with data derived from it. 'i' is relative to zero.
func (s *App01maVendor) TestData(i int) {
    var chr     rune
var i64     int64
var f64     float64
var str     string

    if i < 27 {
        chr = rune(65 + i)      // A
    } else if i < 55 {
        chr = rune(97 + i)      // a
    } else {
        chr = rune(65)          // A
    }

    i64 = int64(i)
f64 = float64(i)
str = string(chr)


    s.Id = i64
            s.Id++       // auto-increment fields are relative to one not zero
            s.Name = str
        s.Addr1 = str
        s.Addr2 = str
        s.City = str
        s.State = str
        s.Zip = str
        s.Curbal = f64
        
}

//----------------------------------------------------------------------------
//                             To String
//----------------------------------------------------------------------------

// ToString converts a record's field to a string.
func (s *App01maVendor) ToString(TitledName string) string {
    var str     string

    switch TitledName {
    
    case "Id":
        	str = fmt.Sprintf("%d", s.Id)

    case "Name":
        	str = s.Name

    case "Addr1":
        	str = s.Addr1

    case "Addr2":
        	str = s.Addr2

    case "City":
        	str = s.City

    case "State":
        	str = s.State

    case "Zip":
        	str = s.Zip

    case "Curbal":
        	{
		s := fmt.Sprintf("%.4f", s.Curbal)
		str = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	}

	default:
	    str = ""
	}

	return str
}

//----------------------------------------------------------------------------
//                             To Strings
//----------------------------------------------------------------------------

// ToStrings converts a record to an array of strings acceptable to CSV and
// other conversion packages.
func (s *App01maVendor) ToStrings() []string {
    var strs    []string
    var str     string

    
        	str = fmt.Sprintf("%d", s.Id)

        strs = append(strs, str)
        	str = s.Name

        strs = append(strs, str)
        	str = s.Addr1

        strs = append(strs, str)
        	str = s.Addr2

        strs = append(strs, str)
        	str = s.City

        strs = append(strs, str)
        	str = s.State

        strs = append(strs, str)
        	str = s.Zip

        strs = append(strs, str)
        	{
		s := fmt.Sprintf("%.4f", s.Curbal)
		str = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	}

        strs = append(strs, str)

	return strs
}

//----------------------------------------------------------------------------
//                             New Struct
//----------------------------------------------------------------------------

// VendorNew creates a new empty struct.
func NewApp01maVendor() *App01maVendor {
    return &App01maVendor{}
}

func NewApp01maVendors() *App01maVendors {
    return &App01maVendors{}
}


