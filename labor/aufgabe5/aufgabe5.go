package aufgabe5

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ArrayAvg.
PUNKTE: 5
BEWERTUNG:
*/

// ArrayAvg erwartet eine int-Slice und liefert den Durchschnittswert aller Elemente.
// Ist die liste leer, soll 0 geliefert werden.
func ArrayAvg(list []float64) float64 {
	// BEGIN-SOLUTION
	if len(list) == 0 {
		return 0
	}
	sum := 0.0
	for _, el := range list {
		sum += el
	}
	return sum / float64(len(list))
	// END-SOLUTION
}
