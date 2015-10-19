# strm-select

A driver for the **go-strm** Go programming language library, that provides the **SELECT** command.

**SELECT** returns all rows but with only certain columns and in a certain order.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/strm-select

[![GoDoc](https://godoc.org/github.com/reiver/strm-select?status.svg)](https://godoc.org/github.com/reiver/strm-select)

## Example
```
package main

import (
	. "github.com/reiver/strm-csv"
	. "github.com/reiver/strm-select"
	. "github.com/reiver/strm-stdout"
)

func main() {
	Begin(CSV, "table.csv").
		Strm(SELECT, "time", "name").
	End(STDOUT, "tsv")
}
```

(Note that in that example dot imports were used.)

## See Also

For more information about **go-strm** and for a list of other drivers, see:
https://github.com/reiver/go-strm
