package hole

import "strings"

var pieceMap = map[rune]rune{
	'r': '♜', 'n': '♞', 'b': '♝', 'q': '♛', 'k': '♚', 'p': '♟',
	'R': '♖', 'N': '♘', 'B': '♗', 'Q': '♕', 'K': '♔', 'P': '♙',
}

var _ = answerFunc("forsyth-edwards-notation", func() []Answer {
	args := shuffle([]string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		"K7/2bb4/1bbb4/3k4/8/8/8/8 w - - 69 272",
		"1n1Rkb1r/p4ppp/4q3/4p1B1/4P3/8/PPP2PPP/2K5 b k - 1 17",
		strings.Join(shuffle([]string{"2k1N3", "7r", "8", "2B5", "3rb3", "4n3", "1R6", "4K3"}), "/") + " w - - 0 50",
	})

	tests := make([]test, len(args))

	for i, arg := range args {
		var buf strings.Builder

		for i, c := range arg {
			if c == ' ' {
				break
			} else if c == '/' {
				buf.WriteByte('\n')
			} else if '1' <= c && c <= '8' {
				if arg[i+1] >= 'A' {
					for n := c - '0'; n > 0; n-- {
						buf.WriteByte(' ')
					}
				}
			} else {
				buf.WriteRune(pieceMap[c])
			}
		}

		tests[i] = test{in: arg, out: buf.String()}
	}

	return outputTestsWithSep("\n\n", tests)
})
