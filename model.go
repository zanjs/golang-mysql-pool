package main

// ModelUser is ...
func ModelUser() (record map[string]string) {
	sql := "SELECT * FROM user"
	// rows := RSQL(sql)
	// rows, err := DB.Query(sql)
	// defer rows.Close()

	// CheckErr(err)

	record = RSQL(sql)

	// fmt.Println(reflect.TypeOf(rows))
	// fmt.Println(reflect.TypeOf(record))

	// fmt.Println(record)

	return record
}
