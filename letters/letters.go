// This package contains the block letter representation of characters
package letters

import (
	"fmt"
)

type BlockLetter struct {
	Contents [6]string		// Contents will have the content representation of a block-letter (All block-letters have height 6 for now.  Element 0 is the top part of the block-letter, Element 1 the next, and so on. 
}

func NewBlockLetter() *BlockLetter {
	return &BlockLetter{}
}

// A BlockWord is a list of BlockLetters
type BlockWord struct {
	Contents [6]string
	Letters []*BlockLetter
}

func NewBlockWord() *BlockWord {
	bw := &BlockWord{}
	bw.Letters = make([]*BlockLetter, 0)
	return bw
}

func NewBlockWordFromString(word string) *BlockWord {
	bw := NewBlockWord()
	for i:=0; i< len(word); i++ {
		if letter, exists := BlockLetterMap[string(word[i])]; exists {
			bw.Letters = append(bw.Letters, letter)
		} else {
			panic(fmt.Sprintf("Do not have support for letter %s", string(word[i])))
		}
	}
	return bw
}

func (bw *BlockWord) clearContents() {
	bw.Contents[0] = ""
	bw.Contents[1] = ""
	bw.Contents[2] = ""
	bw.Contents[3] = ""
	bw.Contents[4] = ""
	bw.Contents[5] = ""
}

func (bw *BlockWord) updateContents() {
	bw.clearContents()
	for _, letter := range bw.Letters {
		bw.Contents[0] += letter.Contents[0]
		bw.Contents[1] += letter.Contents[1]
		bw.Contents[2] += letter.Contents[2]
		bw.Contents[3] += letter.Contents[3]
		bw.Contents[4] += letter.Contents[4]
		bw.Contents[5] += letter.Contents[5]
	}
}

func (bw *BlockWord) Print() {
	bw.updateContents()
	for i := range bw.Contents {
		fmt.Println(bw.Contents[i])
	}
}

// BlockLetterMap will have the letter => block letter representation (ie: "a" => "a" BlockLetter)
var BlockLetterMap = make(map[string]*BlockLetter, 26)

func InitializeBlockLetterMap () {
   
	a := NewBlockLetter()
	a.Contents[0] = ` ______ `
	a.Contents[1] = `|  __  |`
	a.Contents[2] = `| |__| |`
	a.Contents[3] = `|  __  |`
	a.Contents[4] = `| |  | |`
	a.Contents[5] = `|_|  |_|`
	BlockLetterMap["a"] = a

	b := NewBlockLetter()
	b.Contents[0] = ` _______ `
	b.Contents[1] = `|  __   \`
	b.Contents[2] = `| |__| _/`
	b.Contents[3] = `|  __ |_ `
	b.Contents[4] = `| |__|  \`
	b.Contents[5] = `|_______/`
	BlockLetterMap["b"] = b

	c := NewBlockLetter()
	c.Contents[0] = ` _______ `
	c.Contents[1] = `|  _____|`
	c.Contents[2] = `| |      `
	c.Contents[3] = `| |      `
	c.Contents[4] = `| |_____ `
	c.Contents[5] = `|_______|`
	BlockLetterMap["c"] = c
	
	d := NewBlockLetter()
	d.Contents[0] = ` ______  `
	d.Contents[1] = `|  ___ \ `
	d.Contents[2] = `| |   \ \`
	d.Contents[3] = `| |   | |`
	d.Contents[4] = `| |___/ /`
	d.Contents[5] = `|______/ `
	BlockLetterMap["d"] = d

	e := NewBlockLetter()
	e.Contents[0] = ` ______ `
	e.Contents[1] = `|  ____|`
	e.Contents[2] = `| |____ `
	e.Contents[3] = `|  ____|`
	e.Contents[4] = `| |____ `
	e.Contents[5] = `|______|`
	BlockLetterMap["e"] = e
	
	f := NewBlockLetter()
	f.Contents[0] = ` ______ `
	f.Contents[1] = `|  ____|`
	f.Contents[2] = `| |____ `
	f.Contents[3] = `|  ____|`
	f.Contents[4] = `| |     `
	f.Contents[5] = `|_|     `
	BlockLetterMap["f"] = f

	g := NewBlockLetter()
	g.Contents[0] = ` ________ `
	g.Contents[1] = `|  ______|`
	g.Contents[2] = `| |  ____ `
	g.Contents[3] = `| | |__  |`
	g.Contents[4] = `| |____| |`
	g.Contents[5] = `|________|`
	BlockLetterMap["g"] = g

	h := NewBlockLetter()
	h.Contents[0] = ` _     _ `
	h.Contents[1] = `| |   | |`
	h.Contents[2] = `| |___| |`
	h.Contents[3] = `|  ___  |`
	h.Contents[4] = `| |   | |`
	h.Contents[5] = `|_|   |_|`
	BlockLetterMap["h"] = h

	i := NewBlockLetter()
	i.Contents[0] = ` _______ `
	i.Contents[1] = `|__   __|`
	i.Contents[2] = `   | |   `
	i.Contents[3] = `   | |   `
	i.Contents[4] = ` __| |__ `
	i.Contents[5] = `|_______|`
	BlockLetterMap["i"] = i

	j := NewBlockLetter()
	j.Contents[0] = ` _______ `
	j.Contents[1] = `|__   __|`
	j.Contents[2] = `   | |   `
	j.Contents[3] = `   | |   `
	j.Contents[4] = ` __| |   `
	j.Contents[5] = `|____|   `
	BlockLetterMap["j"] = j
  
  k := NewBlockLetter()
	k.Contents[0] = ` _     _ `
	k.Contents[1] = `| |   / /`
	k.Contents[2] = `| |_ / / `
	k.Contents[3] = `|  _ <   `
	k.Contents[4] = `| |  \ \ `
	k.Contents[5] = `|_|   \_\`
	BlockLetterMap["k"] = k

  l := NewBlockLetter()
	l.Contents[0] = ` _       `
	l.Contents[1] = `| |      `
	l.Contents[2] = `| |      `
	l.Contents[3] = `| |      `
	l.Contents[4] = `| |_____ `
	l.Contents[5] = `|_______|`
	BlockLetterMap["l"] = l
  
  m := NewBlockLetter()
	m.Contents[0] = ` _     _ `
	m.Contents[1] = `| \   / |`
	m.Contents[2] = `|  \_/  |`
	m.Contents[3] = `| |\_/| |`
	m.Contents[4] = `| |   | |`
	m.Contents[5] = `|_|   |_|`
	BlockLetterMap["m"] = m

	u := NewBlockLetter()
	u.Contents[0] = ` _     _ `
	u.Contents[1] = `| |   | |`
	u.Contents[2] = `| |   | |`
	u.Contents[3] = `| |   | |`
	u.Contents[4] = `| |___| |`
	u.Contents[5] = `|_______|`
	BlockLetterMap["u"] = u

	q := NewBlockLetter()
	q.Contents[0] = ` _______ `
	q.Contents[1] = `|  ___  |`
	q.Contents[2] = `| |   | |`
	q.Contents[3] = `| | |\| |`
	q.Contents[4] = `| |_\ \ |`
	q.Contents[5] = `|____\_\|`
	BlockLetterMap["q"] = q
  
	z := NewBlockLetter()
	z.Contents[0] = ` _______ `
	z.Contents[1] = `|_____  |`
	z.Contents[2] = `     / / `
	z.Contents[3] = `   / /   `
	z.Contents[4] = ` / /____ `
	z.Contents[5] = `|_______|`
	BlockLetterMap["z"] = z
}
