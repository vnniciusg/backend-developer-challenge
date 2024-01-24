package utils

func QueryAppend(query *string, args *[]interface{}, argsCount *int, separator, field string, value interface{}) {
	if *argsCount > 0 {
		*query += separator
	}
	*query += " " + field
	*args = append(*args, value)

	*argsCount++
}
