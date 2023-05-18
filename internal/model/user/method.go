package user

func (m AuthUser) IsAdmin() bool {
	return m.Role >= ADMIN
}

func (m AuthUser) IsSame(ID string) bool {
	return m.ID == ID
}

func (m *AuthUser) SetRoleString(role string) {
	m.Role = roleDirectories[role]
}
