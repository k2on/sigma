package sigma

func (c *realClient) Handles() ([]Handle, error) {
	rows, err := c.runSQL(`
		SELECT handle.ROWID, handle.id
		from handle
	`)
	if err != nil { return []Handle{}, err }
	defer rows.Close()

	handles := []Handle{}

	for rows.Next() {
		var id int
		var identifier string
		err = rows.Scan(&id, &identifier)
		if err != nil { return []Handle{}, err }
		handles = append(handles, Handle{
			ID: id,
			Identifier: identifier,
		})
	}
	return handles, nil

}
