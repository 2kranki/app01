// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// ioApp01ms contains all the functions
// and data to interact with the SQL Database.

// Generated: Tue Nov 19, 2019 15:46

package App01msCustomer

import (
    "fmt"
    "strings"
	"testing"
)

//============================================================================
//                              Tests
//============================================================================

func TestTestDataApp01msCustomer(t *testing.T) {
    var chr         rune
    var str         string
    var i64         int64
    var f64         float64

    t.Logf("Test.TestData()...\n")
    i64 = 1
    f64 = float64(i64)
    chr = rune(i64 + 65)
    str = string(chr)
    t.Logf("\t i64 = %d\n", i64)
    t.Logf("\t chr = %c\n", chr)
    t.Logf("\t str = (%d)%s\n", len(str), str)

    rcd := NewApp01msCustomer()
    if rcd == nil {
        t.Fatalf("Error: Could not create rcd!\n\n\n")
    }
    rcd.TestData(1)

    
        
            
                if rcd.Num != i64 {
                    t.Fatalf("Error: Invalid data for rcd.Num of %d!\n\n\n", rcd.Num)
                }
            
        
    
        
            if rcd.Name != string(chr) {
                t.Fatalf("Error: Invalid data for rcd.Name of (%d)%s!\n\n\n",
                                    len(rcd.Name), rcd.Name)
            }
        
    
        
            if rcd.Addr1 != string(chr) {
                t.Fatalf("Error: Invalid data for rcd.Addr1 of (%d)%s!\n\n\n",
                                    len(rcd.Addr1), rcd.Addr1)
            }
        
    
        
            if rcd.Addr2 != string(chr) {
                t.Fatalf("Error: Invalid data for rcd.Addr2 of (%d)%s!\n\n\n",
                                    len(rcd.Addr2), rcd.Addr2)
            }
        
    
        
            if rcd.City != string(chr) {
                t.Fatalf("Error: Invalid data for rcd.City of (%d)%s!\n\n\n",
                                    len(rcd.City), rcd.City)
            }
        
    
        
            if rcd.State != string(chr) {
                t.Fatalf("Error: Invalid data for rcd.State of (%d)%s!\n\n\n",
                                    len(rcd.State), rcd.State)
            }
        
    
        
            if rcd.Zip != string(chr) {
                t.Fatalf("Error: Invalid data for rcd.Zip of (%d)%s!\n\n\n",
                                    len(rcd.Zip), rcd.Zip)
            }
        
    
        
            if rcd.Curbal != f64 {
                t.Fatalf("Error: Invalid data for rcd.Curbal of %f!\n\n\n", rcd.Curbal)
            }
        
    


    t.Logf("Test.TestData() - End of Test\n\n\n")
}

func TestToStringApp01msCustomer(t *testing.T) {
    var str         string
    var strRcd      string

    t.Logf("Test.ToStrings()...\n")

    rcd := NewApp01msCustomer()
    if rcd == nil {
        t.Fatalf("Error: Could not create rcd!\n\n\n")
    }
    rcd.TestData(1)

    strRcd = rcd.ToString("Num")
        	str = fmt.Sprintf("%d", rcd.Num)

        if str != strRcd {
            t.Fatalf("Error: Invalid data for %s!\n\n\n", "Num")
        }
    strRcd = rcd.ToString("Name")
        	str = rcd.Name

        if str != strRcd {
            t.Fatalf("Error: Invalid data for %s!\n\n\n", "Name")
        }
    strRcd = rcd.ToString("Addr1")
        	str = rcd.Addr1

        if str != strRcd {
            t.Fatalf("Error: Invalid data for %s!\n\n\n", "Addr1")
        }
    strRcd = rcd.ToString("Addr2")
        	str = rcd.Addr2

        if str != strRcd {
            t.Fatalf("Error: Invalid data for %s!\n\n\n", "Addr2")
        }
    strRcd = rcd.ToString("City")
        	str = rcd.City

        if str != strRcd {
            t.Fatalf("Error: Invalid data for %s!\n\n\n", "City")
        }
    strRcd = rcd.ToString("State")
        	str = rcd.State

        if str != strRcd {
            t.Fatalf("Error: Invalid data for %s!\n\n\n", "State")
        }
    strRcd = rcd.ToString("Zip")
        	str = rcd.Zip

        if str != strRcd {
            t.Fatalf("Error: Invalid data for %s!\n\n\n", "Zip")
        }
    strRcd = rcd.ToString("Curbal")
        	{
		s := fmt.Sprintf("%.4f", rcd.Curbal)
		str = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	}

        if str != strRcd {
            t.Fatalf("Error: Invalid data for %s!\n\n\n", "Curbal")
        }
    


    t.Logf("Test.ToStrings() - End of Test\n\n\n")
}

func TestToStringsApp01msCustomer(t *testing.T) {
    var strs        []string
    var str         string
    var offset      int

    t.Logf("Test.ToStrings()...\n")

    rcd := NewApp01msCustomer()
    if rcd == nil {
        t.Fatalf("Error: Could not create rcd!\n\n\n")
    }
    rcd.TestData(1)

    strs = rcd.ToStrings()
    
        offset = 0
        	str = fmt.Sprintf("%d", rcd.Num)

        if str != strs[offset] {
            t.Fatalf("Error: Invalid data for %s of %s!\n\n\n",
                                    "Num", strs[offset])
        }
    
        offset = 1
        	str = rcd.Name

        if str != strs[offset] {
            t.Fatalf("Error: Invalid data for %s of %s!\n\n\n",
                                    "Name", strs[offset])
        }
    
        offset = 2
        	str = rcd.Addr1

        if str != strs[offset] {
            t.Fatalf("Error: Invalid data for %s of %s!\n\n\n",
                                    "Addr1", strs[offset])
        }
    
        offset = 3
        	str = rcd.Addr2

        if str != strs[offset] {
            t.Fatalf("Error: Invalid data for %s of %s!\n\n\n",
                                    "Addr2", strs[offset])
        }
    
        offset = 4
        	str = rcd.City

        if str != strs[offset] {
            t.Fatalf("Error: Invalid data for %s of %s!\n\n\n",
                                    "City", strs[offset])
        }
    
        offset = 5
        	str = rcd.State

        if str != strs[offset] {
            t.Fatalf("Error: Invalid data for %s of %s!\n\n\n",
                                    "State", strs[offset])
        }
    
        offset = 6
        	str = rcd.Zip

        if str != strs[offset] {
            t.Fatalf("Error: Invalid data for %s of %s!\n\n\n",
                                    "Zip", strs[offset])
        }
    
        offset = 7
        	{
		s := fmt.Sprintf("%.4f", rcd.Curbal)
		str = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	}

        if str != strs[offset] {
            t.Fatalf("Error: Invalid data for %s of %s!\n\n\n",
                                    "Curbal", strs[offset])
        }
    


    t.Logf("Test.ToStrings() - End of Test\n\n\n")
}

