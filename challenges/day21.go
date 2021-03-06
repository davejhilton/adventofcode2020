package challenges

import (
	"fmt"
	"sort"
	"strings"

	"github.com/davejhilton/adventofcode2020/log"
)

func day21_part1(input []string) (string, error) {
	foods := day21_parse(input)

	log.Debugln("FOODS: ")
	for _, f := range foods {
		log.Debugf("  - %v\n", f)
	}

	log.Debug("\n\n")
	suspects := make(map[string][]string)
	culprits := make(map[string]string)

	var loop int
	for loop < 10 && (len(suspects) > 0 || len(culprits) == 0) {
		log.Debugf("LOOP: %d\n", loop)
		log.Debugln("  SUSPECTS: ")
		for a, in := range suspects {
			log.Debugf("    - %-10s: %s\n", a, in)
		}
		log.Debugln("\n  CULPRITS: ")
		for i, a := range culprits {
			log.Debugf("    - %-10s: %s\n", a, i)
		}
		log.Debugln("\n------------------")
		for _, f := range foods {
			for _, a := range f.Allergens {
				if _, ok := suspects[a]; !ok {
					known := false
					for _, al := range culprits {
						if al == a {
							known = true
							break
						}
					}
					if !known {
						suspects[a] = make([]string, len(f.Ingredients))
						copy(suspects[a], f.Ingredients)
					}
				} else {
					if len(suspects[a]) == 1 {
						culprits[suspects[a][0]] = a
						delete(suspects, a)
					} else {
						stillSuspect := make([]string, 0)
						for _, i := range suspects[a] {
							if _, ok := culprits[i]; !ok {
								if f.HasIngredient(i) {
									stillSuspect = append(stillSuspect, i)
								}
							}
						}
						suspects[a] = stillSuspect
					}
				}
			}
		}
		loop++
	}

	harmless := make(map[string]int)
	sum := 0
	for _, f := range foods {
		for _, in := range f.Ingredients {
			if _, ok := culprits[in]; !ok {
				harmless[in] += 1
				sum++
			}
		}
	}

	log.Debugln("HARMLESS: ")
	for i, c := range harmless {
		log.Debugf("  - %-10s: %d\n", i, c)
	}
	return fmt.Sprintf("%d", sum), nil
}

func day21_part2(input []string) (string, error) {
	foods := day21_parse(input)

	log.Debugln("FOODS: ")
	for _, f := range foods {
		log.Debugf("  - %v\n", f)
	}

	log.Debug("\n\n")
	suspects := make(map[string][]string)
	culprits := make(map[string]string)

	var loop int
	for loop < 100 && (len(suspects) > 0 || len(culprits) == 0) {
		log.Debugf("LOOP: %d\n", loop)
		log.Debugln("  SUSPECTS: ")
		for a, in := range suspects {
			log.Debugf("    - %-10s: %s\n", a, in)
		}
		log.Debugln("\n  CULPRITS: ")
		for a, i := range culprits {
			log.Debugf("    - %-10s: %s\n", a, i)
		}
		log.Debugln("\n------------------")
		for _, f := range foods {
			for _, a := range f.Allergens {
				if _, ok := suspects[a]; !ok {
					if _, ok2 := culprits[a]; !ok2 {
						suspects[a] = make([]string, len(f.Ingredients))
						copy(suspects[a], f.Ingredients)
					}
				} else {
					if len(suspects[a]) == 1 {
						culprits[a] = suspects[a][0]
						delete(suspects, a)
					} else {
						stillSuspect := make([]string, 0)
						for _, i := range suspects[a] {
							culprit := false
							for _, in := range culprits {
								if in == i {
									culprit = true
									break
								}
							}
							if !culprit && f.HasIngredient(i) {
								stillSuspect = append(stillSuspect, i)
							}
						}
						suspects[a] = stillSuspect
					}
				}
			}
		}
		loop++
	}

	allergens := make([]string, 0, len(culprits))
	for a := range culprits {
		allergens = append(allergens, a)
	}
	sort.Strings(allergens)

	var b strings.Builder
	for i, a := range allergens {
		b.WriteString(culprits[a])
		if i != len(allergens)-1 {
			b.WriteString(",")
		}
	}
	return fmt.Sprintf("%s", b.String()), nil
}

func day21_parse(input []string) []day21_food {
	foods := make([]day21_food, 0, len(input))

	for _, line := range input {
		halves := strings.Split(line, " (contains ")
		halves[1] = strings.TrimSuffix(halves[1], ")")
		foods = append(foods, day21_food{
			Ingredients: strings.Split(halves[0], " "),
			Allergens:   strings.Split(halves[1], ", "),
		})
	}
	return foods
}

type day21_food struct {
	Ingredients []string
	Allergens   []string
}

func (f day21_food) ListsAllergen(a string) bool {
	for _, ag := range f.Allergens {
		if ag == a {
			return true
		}
	}
	return false
}

func (f day21_food) HasIngredient(in string) bool {
	for _, ing := range f.Ingredients {
		if ing == in {
			return true
		}
	}
	return false
}

func init() {
	registerChallengeFunc(21, 1, "day21.txt", day21_part1)
	registerChallengeFunc(21, 2, "day21.txt", day21_part2)
}
