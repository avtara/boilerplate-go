package queries

const (
	GetLastLoginByUsernameOrEmail = `
		SELECT last_login FROM accounts WHERE username = $1 OR email = $2 LIMIT 1
	`

	CreateAccount = `
		INSERT INTO accounts (username, password, email) VALUES ($1, $2, $3) RETURNING user_id
	`
)
