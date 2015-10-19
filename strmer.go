package verboten


import (
	"github.com/reiver/go-strm/driver"
)


const (
	// SELECT is a (string) constant that this Beginner driver
	// is registered under.
	SELECT = "SELECT"

	defaultLimit = 5
)


func init() {
	strmDriver := newStrmer()

	strmdriver.RegisterStrmer(SELECT, strmDriver)
}


type internalStrmer struct{}


func newStrmer() strmdriver.Strmer {
	strmDriver := internalStrmer{

	}

	return &strmDriver
}



func (strmDriver *internalStrmer) Strm(src <-chan []interface{}, dst chan<- []interface{}, args ...interface{}) {

	// Parse args.
	wantedColumns := []string{}
	for _, arg := range args {
		column, ok := arg.(string)
		if !ok {
//@TODO: Throw something better.
			panic("Wrong type for column name.")
		}

		wantedColumns = append(wantedColumns, column)
	}

	// Execute.
	//
	// Assume the first row is the header which containes the column names.
	columnMapping := map[string]int{}
	first := true
	i :=0
	for row := range src {
		if first {
			first = false

			header := row
			for ii, datum := range header {
				columnName, ok := datum.(string)
				if !ok {
//@TODO: Throw an error
					panic("Not a column name.")
				}

				columnMapping[columnName] = ii
			}

			newRow := make([]interface{}, len(wantedColumns))
			for ii, wantedColumn := range wantedColumns {
				newRow[ii] = wantedColumn
			}
			dst <- newRow

		} else {
			newRow := make([]interface{}, len(wantedColumns))
			for ii, wantedColumn := range wantedColumns {

				newRow[ii] = row[columnMapping[wantedColumn]]
			}
			dst <- newRow
		}

		i++
	}
	close(dst)
}
