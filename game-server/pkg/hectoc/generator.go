package hectoc

import (
	"math/rand"
	"strconv"
	"time"
)

const MAX_ATTEMPTS = 100

type Hectoc struct {
	Problem string
	Solutions []string
}

var insertions = []string{"+", "-", "*", "/", "^", "(", ")", ""}

func generateSolutions(current string, remaining string, solutions *[]string, bracesParity int) {
	if len(remaining) == 0 {
		// Check if the current expression is valid
		if bracesParity > 0 {
			return
		} else {
			for i := 1; i <= bracesParity; i++ {
				current += ")"
			}
		}

		verified, _ := Verify(current)

		if verified {
			*solutions = append(*solutions, current)
		}

		return
	}

	nextDigit := remaining[0]
	newRemaining := remaining[1:]

	for _, insertion := range insertions {
		if insertion == "(" {
			bracesParity--
		} else if insertion == ")" {
			bracesParity++
		}

		newCurrent := current + insertion + string(nextDigit)

		generateSolutions(newCurrent, newRemaining, solutions, bracesParity)
	}
}

func (h *Hectoc) solve() {
	var solutions []string

	// Generate all possible solutions
	generateSolutions("", h.Problem, &solutions, 0)
	h.Solutions = solutions
}

var excluded = map[string]struct{}{
	"112117": {}, "114123": {}, "115567": {}, "115827": {}, "116567": {}, "121143": {}, "121581": {}, "131116": {}, "141171": {}, "156567": {},
	"167181": {}, "167451": {}, "171717": {}, "175117": {}, "176611": {}, "178181": {}, "178188": {}, "178881": {}, "178888": {}, "178988": {},
	"184156": {}, "185571": {}, "188788": {}, "188887": {}, "211143": {}, "211539": {}, "351117": {}, "361869": {}, "363369": {}, "366369": {},
	"383888": {}, "388838": {}, "598999": {}, "611171": {}, "611177": {}, "617667": {}, "617676": {}, "617766": {}, "633639": {}, "639669": {},
	"661667": {}, "664149": {}, "664989": {}, "666117": {}, "666161": {}, "666166": {}, "666615": {}, "666651": {}, "666661": {}, "666667": {},
	"666761": {}, "667661": {}, "675151": {}, "676111": {}, "676167": {}, "676176": {}, "676667": {}, "676761": {}, "677761": {}, "681181": {},
	"681667": {}, "711161": {}, "711781": {}, "717767": {}, "718178": {}, "718887": {}, "718888": {}, "719171": {}, "719878": {}, "745171": {},
	"747778": {}, "747787": {}, "747877": {}, "748777": {}, "761117": {}, "761161": {}, "761767": {}, "766111": {}, "766861": {}, "767717": {},
	"767761": {}, "771818": {}, "773167": {}, "773781": {}, "776761": {}, "778181": {}, "778451": {}, "778551": {}, "778978": {}, "781117": {},
	"781171": {}, "781281": {}, "781676": {}, "781718": {}, "781888": {}, "781978": {}, "781987": {}, "787888": {}, "787889": {}, "788189": {},
	"788789": {}, "788818": {}, "788878": {}, "788881": {}, "788971": {}, "789788": {}, "791887": {}, "797881": {}, "799971": {}, "817781": {},
	"817789": {}, "817881": {}, "817888": {}, "818878": {}, "819787": {}, "819877": {}, "819878": {}, "819887": {}, "838383": {}, "838588": {},
	"838858": {}, "838883": {}, "853878": {}, "858838": {}, "871888": {}, "877889": {}, "878181": {}, "878188": {}, "878538": {}, "878787": {},
	"878789": {}, "878881": {}, "878887": {}, "878988": {}, "881788": {}, "881878": {}, "881887": {}, "881987": {}, "885838": {}, "887778": {},
	"887818": {}, "887881": {}, "887888": {}, "888178": {}, "888187": {}, "888383": {}, "888717": {}, "888781": {}, "888787": {}, "888789": {},
	"888817": {}, "888861": {}, "888871": {}, "888877": {}, "897878": {}, "897887": {}, "898771": {}, "898781": {}, "898878": {}, "951999": {},
	"958999": {}, "961999": {}, "969199": {}, "969659": {}, "978788": {}, "978887": {},
}

var fallbacks = []string{"123456", "999541", "472319", "327924"}

func Generate() *Hectoc {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for attempts:=1; attempts <= 100; attempts++ {
		sequence := ""

		for i:=1; i<=6; i++ {
			digit := r.Intn(9) + 1
			sequence += strconv.Itoa(digit)
		}

		if _, found := excluded[sequence]; !found {
			h := &Hectoc{
				Problem: sequence,
			}

			h.solve()

			return h
		}
	}

	sequence := fallbacks[r.Intn(len(fallbacks))]

	h := &Hectoc{
		Problem: sequence,
	}

	h.solve()

	return h

}