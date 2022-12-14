package db

import (
	"database/sql"

	"github.com/rajatxs/go-hyper/peer"
)

// Returns RawKeyBundle from database by given address
func (d *Database) GetPeerRawKeyBundle(addr string) (rkb *peer.RawKeyBundle, err error) {
	row := d.instance.QueryRow(
		"SELECT addr, auth_key, exch_key FROM peer_keys WHERE addr = ?;",
		addr)

	if err = row.Err(); err != nil {
		return nil, err
	}

	rkb = &peer.RawKeyBundle{}
	if err = row.Scan(
		&rkb.Address,
		&rkb.AuthKey,
		&rkb.ExchangeKey); err != nil {
		return nil, err
	}

	return rkb, nil
}

// Writes given RawKeyBundle to database
func (d *Database) SavePeerRawKeyBundle(rpo *peer.RawKeyBundle) (sql.Result, error) {
	if stmt, err := d.instance.Prepare(
		"INSERT INTO peer_keys (addr, auth_key, exch_key) VALUES (?, ?, ?);"); err != nil {
		return nil, err
	} else {
		return stmt.Exec(
			rpo.Address,
			rpo.AuthKey,
			rpo.ExchangeKey)
	}
}

// Check availability of peer in database by given address
func (d *Database) HasPeerKeyBundle(addr string) (bool, error) {
	var count int
	row := d.instance.QueryRow(
		"SELECT COUNT(addr) AS count FROM peer_keys WHERE addr = ?;",
		addr)

	if err := row.Scan(&count); err != nil {
		return false, nil
	}

	return count == 1, nil
}

// Deletes single peer record by given address
func (d *Database) DeletePeerKeyBundle(addr string) (sql.Result, error) {
	if stmt, err := d.instance.Prepare("DELETE FROM peer_keys WHERE addr = ?;"); err != nil {
		return nil, err
	} else {
		return stmt.Exec(addr)
	}
}
