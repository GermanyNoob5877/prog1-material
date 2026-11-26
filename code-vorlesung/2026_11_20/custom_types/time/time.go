package time

import "fmt"

type Time struct {
	Hours   int
	Minutes int
}

//Wenn ein Custom Datentyp eine Funktion mit dem Namen "String"
//hat wird diese bei Print Statements des Datentyps aufgerufen!
func (t Time) String() string {
	result := ""

	result += fmt.Sprint(t.Hours)
	result += ":"
	result += fmt.Sprint(t.Minutes)

	return result
}
