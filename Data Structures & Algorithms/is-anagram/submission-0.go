import "reflect"

func isAnagram(s string, t string) bool {

	smap := map[rune]int{}
	tmap := map[rune]int{}

	tslice := []rune(t)
	sslice := []rune(s)

	if len(tslice) != len(sslice) {
		return false
	}

	for i := 0; i < len(tslice); i++ {
		_, tok := tmap[tslice[i]]
		_, sok := smap[sslice[i]]

		if !tok {
			tmap[tslice[i]] = 1
		}

		if !sok {
			smap[sslice[i]] = 1
		}

		tmap[tslice[i]] += 1
		smap[sslice[i]] += 1
	}

	return reflect.DeepEqual(tmap, smap)
}
