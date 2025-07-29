package lr

type StringResponse []string
type AnyResponse []any

type StringArrResponse [][]string

// NoResponse marks places where we don't expect any response data, even though
// it would allow anything lightroom happened to return
type NoResponse any
