package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion GetStringsBetweenEmptyLines.
PUNKTE: 15
BEWERTUNG:
*/

// GetStringsBetweenEmptyLines erwartet eine Liste und.
// liefert die Teilliste, die zwischen dem ersten und dem letzten leeren String liegt.
// Die beiden leeren Strings sollen nicht zum Ergebnis gehören.
// Falls die Liste keine zwei leeren Strings enthält, soll die leere Liste geliefert werden.
func GetStringsBetweenEmptyLines(list []string) []string {
	// BEGIN-SOLUTION
	firstpos := -1
	lastpos := -1
	for i, s := range list {
		if s == "" {
			if firstpos == -1 {
				firstpos = i
			} else {
				lastpos = i
			}
		}
	}
	if lastpos <= firstpos {
		return []string{}
	}
	// END-SOLUTION
	return list[firstpos+1 : lastpos]
}
