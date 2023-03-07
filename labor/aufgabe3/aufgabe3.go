package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountGreater.
Die Funktion muss rekursiv sein.
PUNKTE: 10
BEWERTUNG:
*/

// CountGreater erwartet eine Liste von Zahlen und eine Zahl x.
// Sie liefert die Anzahl der Zahlen darin, die größer als x sind.
func CountGreater(list []int, x int) int {
	// BEGIN-SOLUTION
	if len(list) == 0 {
		return 0
	}
	head, result := list[0], CountGreater(list[1:], x)
	if head > x {
		result++
	}
	return result
	// END-SOLUTION
}
