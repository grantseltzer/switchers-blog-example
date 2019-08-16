package switchers

func SwitchFunction(a, b int, c *int) string {
	switch *c {
	case a:
		return "a"
	case b:
		return "b"
	default:
		return "c"
	}
}
