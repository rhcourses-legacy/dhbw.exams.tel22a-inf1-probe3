package aufgabe7

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Join.
PUNKTE: 10
BEWERTUNG:
*/

// Join erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen aus l1,
// die in wenigstens einer der beiden Listen vorhanden sind.
// Gehen Sie davon aus, dass jedes Element in jeder Liste höchstens einmal vorkommt.
// Die Ergebnisliste soll keine Duplikate enthalten.
func Join(l1, l2 []int) []int {
	// BEGIN-SOLUTION
	result := append([]int{}, l1...)
	for _, el2 := range l2 {
		found := false
		for _, elr := range result {
			if el2 == elr {
				found = true
			}
		}
		if !found {
			result = append(result, el2)
		}
	}
	return result
	// END-SOLUTION
}
