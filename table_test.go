package index_test

type testTable struct {
	ltr       bool
	exclusive bool
	indexOf   string
	occur     uint
	value     string
	expected  int
	errMsg    string
}

func indexCreationTestTable() []testTable {
	return []testTable{
		{ltr: true, exclusive: true, indexOf: ".", occur: 1},
		{ltr: true, exclusive: true, indexOf: "..", occur: 1},
		{ltr: true, exclusive: true, indexOf: "-", occur: 2},
		{ltr: true, exclusive: true, indexOf: "_", occur: 3},
		{ltr: true, exclusive: true, indexOf: "-_-", occur: 4},
		{ltr: true, exclusive: true, indexOf: "-._", occur: 5},
		{ltr: true, exclusive: true, indexOf: "--", occur: 6},
		{ltr: true, exclusive: true, indexOf: "__", occur: 7},

		{ltr: true, exclusive: false, indexOf: "a", occur: 8},
		{ltr: true, exclusive: false, indexOf: "b.", occur: 9},
		{ltr: true, exclusive: false, indexOf: ".y", occur: 10},
		{ltr: true, exclusive: false, indexOf: "xyz", occur: 11},
		{ltr: true, exclusive: false, indexOf: "AbC", occur: 22},
		{ltr: true, exclusive: false, indexOf: "p", occur: 30},
		{ltr: true, exclusive: false, indexOf: "so", occur: 33},

		{ltr: false, exclusive: true, indexOf: "1", occur: 44},
		{ltr: false, exclusive: true, indexOf: "2", occur: 23},
		{ltr: false, exclusive: true, indexOf: "24", occur: 19},
		{ltr: false, exclusive: true, indexOf: "234", occur: 21},
		{ltr: false, exclusive: true, indexOf: "5234", occur: 23},
		{ltr: false, exclusive: true, indexOf: "90", occur: 55},
		{ltr: false, exclusive: true, indexOf: "01", occur: 15},
		{ltr: false, exclusive: true, indexOf: "11", occur: 2},

		{ltr: false, exclusive: false, indexOf: "a01-.", occur: 1},
		{ltr: false, exclusive: false, indexOf: "b02-.", occur: 1},
		{ltr: false, exclusive: false, indexOf: "_ab2.", occur: 2},
		{ltr: false, exclusive: false, indexOf: "as29_df", occur: 1},
		{ltr: false, exclusive: false, indexOf: "12-w1342", occur: 2},
		{ltr: false, exclusive: false, indexOf: "a01-ww", occur: 3},
		{ltr: false, exclusive: false, indexOf: "xw2a01-", occur: 1},
		{ltr: false, exclusive: false, indexOf: "o1234", occur: 1},
		{ltr: false, exclusive: false, indexOf: "123400_", occur: 2},
	}
}
func indexCreationErrorTestTable() []testTable {
	errMsg := "indexOf cannot be empty"
	emptyIndexOf := []testTable{
		{ltr: true, exclusive: true, indexOf: "", occur: 1, value: ".wow", errMsg: errMsg},
		{ltr: true, exclusive: true, indexOf: "", occur: 2, value: "www.begopher.com", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "", occur: 2, value: "wwww.x.com", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "", occur: 3, value: "wwww.xww.com", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "", occur: 1, value: "wwww.xww.com", errMsg: errMsg},
		{ltr: false, exclusive: true, indexOf: "", occur: 1, value: "wwww.x.com", errMsg: errMsg},
		{ltr: false, exclusive: true, indexOf: "", occur: 2, value: "wwww.x.com", errMsg: errMsg},
		{ltr: false, exclusive: false, indexOf: "", occur: 1, value: "wwww.xww.com", errMsg: errMsg},
		{ltr: false, exclusive: false, indexOf: "", occur: 2, value: "wwwow.xww.com", errMsg: errMsg},
	}
	errMsg = "occur value must be bigger than zero"
	zeroOccur := []testTable{
		{ltr: true, exclusive: true, indexOf: "kk", occur: 0, value: ".wow", errMsg: errMsg},
		{ltr: true, exclusive: true, indexOf: ".", occur: 0, value: "www.begopher.com", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "x", occur: 0, value: "wwww.x.com", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "ww", occur: 0, value: "wwww.xww.com", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "o", occur: 0, value: "wwww.xww.com", errMsg: errMsg},
		{ltr: false, exclusive: true, indexOf: "m", occur: 0, value: "wwww.x.com", errMsg: errMsg},
		{ltr: false, exclusive: true, indexOf: "j", occur: 0, value: "wwww.x.com", errMsg: errMsg},
		{ltr: false, exclusive: false, indexOf: "xyz", occur: 0, value: "wwww.xww.com", errMsg: errMsg},
		{ltr: false, exclusive: false, indexOf: "xw", occur: 0, value: "wwwow.xww.com", errMsg: errMsg},
	}
	var testTable []testTable
	testTable = append(testTable, emptyIndexOf...)
	testTable = append(testTable, zeroOccur...)
	return testTable
}
func strictIndexTestTable() []testTable {
	return []testTable{
		{ltr: true, exclusive: true, indexOf: ".", occur: 1, value: ".wow", expected: 0},
		{ltr: true, exclusive: true, indexOf: ".", occur: 1, value: "www.begopher.com", expected: 3},
		{ltr: true, exclusive: true, indexOf: ".", occur: 2, value: "www.begopher.com", expected: 12},
		{ltr: true, exclusive: true, indexOf: "ww", occur: 1, value: "wwww.x.com", expected: 0},
		{ltr: true, exclusive: true, indexOf: "ww", occur: 2, value: "wwww.x.com", expected: 2},
		{ltr: true, exclusive: true, indexOf: "ww", occur: 3, value: "wwww.xww.com", expected: 6},
		{ltr: true, exclusive: true, indexOf: "ww", occur: 4, value: "wwww.xww.cww", expected: 10},
		{ltr: true, exclusive: true, indexOf: "o", occur: 1, value: "wwww.xww.com", expected: 10},

		{ltr: true, exclusive: false, indexOf: ".", occur: 1, value: ".wow", expected: 1},
		{ltr: true, exclusive: false, indexOf: ".", occur: 1, value: "www.begopher.com", expected: 4},
		{ltr: true, exclusive: false, indexOf: ".", occur: 2, value: "www.begopher.com", expected: 13},
		{ltr: true, exclusive: false, indexOf: "ww", occur: 1, value: "wwww.x.com", expected: 2},
		{ltr: true, exclusive: false, indexOf: "ww", occur: 2, value: "wwww.x.com", expected: 4},
		{ltr: true, exclusive: false, indexOf: "ww", occur: 3, value: "wwww.xww.com", expected: 8},
		{ltr: true, exclusive: false, indexOf: "o", occur: 1, value: "wwww.xww.com", expected: 11},

		{ltr: false, exclusive: true, indexOf: ".", occur: 1, value: ".abc", expected: 0},
		{ltr: false, exclusive: true, indexOf: ".", occur: 1, value: "www.begopher.com", expected: 12},
		{ltr: false, exclusive: true, indexOf: ".", occur: 2, value: "www.begopher.com", expected: 3},
		{ltr: false, exclusive: true, indexOf: "ww", occur: 1, value: "wwww.x.com", expected: 2},
		{ltr: false, exclusive: true, indexOf: "ww", occur: 2, value: "wwww.x.com", expected: 0},
		{ltr: false, exclusive: true, indexOf: "xw", occur: 1, value: "wwww.xww.com", expected: 5},
		{ltr: false, exclusive: true, indexOf: "o", occur: 1, value: "wwww.xww.como", expected: 12},
		{ltr: false, exclusive: true, indexOf: "o", occur: 2, value: "wwww.xww.como", expected: 10},

		{ltr: false, exclusive: false, indexOf: ".", occur: 1, value: ".wow", expected: 1},
		{ltr: false, exclusive: false, indexOf: ".", occur: 1, value: "www.begopher.com", expected: 13},
		{ltr: false, exclusive: false, indexOf: ".", occur: 2, value: "www.begopher.com", expected: 4},
		{ltr: false, exclusive: false, indexOf: "ww", occur: 1, value: "wwww.x.com", expected: 4},
		{ltr: false, exclusive: false, indexOf: "ww", occur: 2, value: "wwww.x.com", expected: 2},
		{ltr: false, exclusive: false, indexOf: "ww", occur: 3, value: "wwww.xww.com", expected: 2},
		{ltr: false, exclusive: false, indexOf: "xw", occur: 1, value: "wwww.xww.com", expected: 7},
		{ltr: false, exclusive: false, indexOf: "o", occur: 1, value: "wwwow.xww.com", expected: 12},
		{ltr: false, exclusive: false, indexOf: "o", occur: 2, value: "wwwow.xww.com", expected: 4},
	}
}
func looseIndexTestTable() []testTable {
	return []testTable{}
}

func indexNotFoundTestTable() []testTable {
	errMsg := "No match found"
	return []testTable{
		{ltr: true, exclusive: true, indexOf: "kk", occur: 1, value: ".wow", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "moc", occur: 1, value: "wwww.x.com", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "wx", occur: 1, value: "wwww.xww.com", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "oo", occur: 1, value: "wwww.xww.com", errMsg: errMsg},
		{ltr: false, exclusive: true, indexOf: "j", occur: 1, value: "www.x.com", errMsg: errMsg},
		//// occur above what exist
		{ltr: true, exclusive: true, indexOf: ".", occur: 3, value: "www.begopher.com", errMsg: errMsg},
		{ltr: false, exclusive: true, indexOf: "w", occur: 4, value: "www.x.com", errMsg: errMsg},
		{ltr: false, exclusive: false, indexOf: "xyz", occur: 1, value: "www.xww.com", errMsg: errMsg},
		{ltr: false, exclusive: false, indexOf: "xw", occur: 2, value: "wwow.xww.com", errMsg: errMsg},
	}

}
func emptyValueTestTable() []testTable {
	errMsg := "Value cannot be empty"
	return []testTable{
		{ltr: true, exclusive: true, indexOf: "kk", occur: 1, value: "", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "moc", occur: 1, value: "", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "wx", occur: 1, value: "", errMsg: errMsg},
		{ltr: true, exclusive: false, indexOf: "oo", occur: 1, value: "", errMsg: errMsg},
		{ltr: false, exclusive: true, indexOf: "j", occur: 1, value: "", errMsg: errMsg},
		//// occur above what exist
		{ltr: true, exclusive: true, indexOf: ".", occur: 3, value: "", errMsg: errMsg},
		{ltr: false, exclusive: true, indexOf: "w", occur: 4, value: "", errMsg: errMsg},
		{ltr: false, exclusive: false, indexOf: "xyz", occur: 1, value: "", errMsg: errMsg},
		{ltr: false, exclusive: false, indexOf: "xw", occur: 2, value: "", errMsg: errMsg},
	}

}
