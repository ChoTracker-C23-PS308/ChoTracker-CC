package user

const (
	DEFAULT uint32 = 1
	ADMIN   uint32 = 2
)

var (
	roleDirectories map[string]uint32 = map[string]uint32{
		"":      DEFAULT,
		"admin": ADMIN,
	}
)
