package main

func main() {

}

func backspaceCompare(S string, T string) bool {
	sBsNum, tBsNum := 0, 0
	sp, tp := len(S)-1, len(T)-1

	for sp >= 0 || tp >= 0 {
		for sp >= 0 {
			if S[sp] == '#' {
				sp--
				sBsNum++
			} else if sBsNum > 0 {
				sp--
				sBsNum--
			} else {
				break
			}
		}
		for tp >= 0 {
			if T[tp] == '#' {
				tp--
				tBsNum++
			} else if tBsNum > 0 {
				tp--
				tBsNum--
			} else {
				break
			}
		}

		if sp >= 0 && tp >= 0 {
			if S[sp] != T[tp] {
				return false
			}
			sp--
			tp--
		} else if sp >= 0 || tp >= 0 {
			return false
		}
	}

	return true
}
