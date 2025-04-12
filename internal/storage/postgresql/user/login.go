package psqlUser

import "context"

func (r *Repository) Login(ctx context.Context, email string) error {

	query := `UPDATE users
			  SET last_login = NOW()
			  WHERE email = $1;`

	_, err := r.db.ExecContext(ctx, query, email)
	if err != nil {
		return err
	}
	return nil
}
