package chain

func ChanToSlice(c chan Document) []Document {
    s := make([]Document, 0)
    for i := range c {
        s = append(s, i)
    }
    return s
}
