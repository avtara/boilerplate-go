package queries

const (
	GetLastLoginByUsernameOrEmail = `
		SELECT last_login FROM accounts WHERE username = $1 OR email = $2 LIMIT 1
	`
)
