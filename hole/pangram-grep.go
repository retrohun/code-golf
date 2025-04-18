package hole

import (
	"math/rand/v2"
	"slices"
	"strings"
)

var _ = answerFunc("pangram-grep", func() []Answer {
	return outputTests(
		pangramGrepTests(2, 5),
		pangramGrepTests(0, 0),
	)
})

func pangramGrepTests(l, r int) []test {
	// They all start lowercase and valid.
	pangrams := shuffle([][]byte{
		[]byte("6>_4\"gv9lb?2!ic7}=-m'fd30ph].o%@w+[8unk&t1es<az(x;${^y#)q,rj\\5/*:`"),
		[]byte(`a large fawn jumped quickly over white zinc boxes.`),
		[]byte(`all questions asked by five watched experts amaze the judge.`),
		[]byte(`a quick movement of the enemy will jeopardize six gunboats.`),
		[]byte(`back in june we delivered oxygen equipment of the same size.`),
		[]byte(`battle of thermopylae: quick javelin grazed wry xerxes.`),
		[]byte(`bored? craving a pub quiz fix? why, just come to the royal oak!`),
		[]byte(`bprsjzfwdqyaxgckilvunthemo`),
		[]byte(`brawny gods just flocked up to quiz and vex him.`),
		[]byte(`cute, kind, jovial, foxy physique, amazing beauty? wowser!`),
		[]byte(`fix problem quickly with galvanized jets.`),
		[]byte(`foxy parsons quiz and cajole the lovably dim wiki-girl.`),
		[]byte(`grumpy wizards make toxic brew for the evil queen and jack.`),
		[]byte(`hey zach, should i program a hex editor in java? why not sql or brainf--k!`),
		[]byte(`how razorback-jumping frogs can level six piqued gymnasts!`),
		[]byte(`jackie will budget for the most expensive zoology equipment.`),
		[]byte(`jack quietly moved up front and seized the big ball of wax.`),
		[]byte(`jim quickly realized that the beautiful gowns are expensive.`),
		[]byte(`just poets wax boldly as kings and queens march over fuzz.`),
		[]byte(`my faxed joke won a pager in the cable tv quiz show.`),
		[]byte(`quirky spud boys can jam after zapping five worthy polysixes.`),
		[]byte(`sixty zips were quickly picked from the woven jute bag.`),
		[]byte(`the quick brown fox jumps over the lazy dog.`),
		[]byte(`the wizard quickly jinxed the gnomes before they vaporized.`),
		[]byte(`when zombies arrive, quickly fax judge pat.`),
	})

	for i, pangram := range pangrams {
		clone := slices.Clone(pangram)

		// Replace letter `i` with a different random letter.
		old := 'a' + byte(i)
		new := 'a' + byte((i+rand.IntN(25)+1)%26)
		for j, letter := range clone {
			if letter == old {
				clone[j] = new
			}
		}

		// Replace 0-4 other letters with random letters.
		for times := rand.IntN(5); times > 0; times-- {
			old := 'a' + byte(rand.IntN(26))
			new := 'a' + byte(rand.IntN(26))
			for j, letter := range clone {
				if letter == old {
					clone[j] = new
				}
			}
		}

		pangrams = append(pangrams, clone)
	}

	// Add alphabet with one letter missing
	for del := 'a'; del <= 'z'; del++ {
		str := []byte{}
		for c := 'a'; c <= 'z'; c++ {
			if c != del {
				str = append(str, byte(c))
			}
		}
		pangrams = append(pangrams, shuffle(str))
	}

	// Uppercase random letters.
	for _, pangram := range pangrams {
		for j, letter := range pangram {
			if 'a' <= letter && letter <= 'z' && rand.IntN(2) == 0 {
				pangram[j] -= 32
			}
		}
	}

	// Insert l-r random post-'z' characters
	for i, pangram := range pangrams {
		for times := rand.IntN(r-l+1) + l; times > 0; times-- {
			c := '{' + byte(rand.IntN(4))
			pos := rand.IntN(len(pangram))

			pangram = append(pangram, 0)
			copy(pangram[pos+1:], pangram[pos:])
			pangram[pos] = c
			pangrams[i] = pangram
		}
	}

	tests := make([]test, len(pangrams))

outer:
	for i, pangram := range shuffle(pangrams) {
		str := string(pangram)
		tests[i].in = str

		for c := 'a'; c <= 'z'; c++ {
			if !strings.ContainsRune(str, c) && !strings.ContainsRune(str, c-32) {
				continue outer
			}
		}

		tests[i].out = str
	}

	return tests
}
