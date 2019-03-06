package storage

import (

	// This line is must for working MySQL database
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/team-monitoring/comedian/model"
)

// CreateChannel creates standup entry in database
func (m *MySQL) CreateChannel(c model.Channel) (model.Channel, error) {
	res, err := m.conn.Exec(
		"INSERT INTO `channels` (team_id, channel_name, channel_id, channel_standup_time) VALUES (?, ?, ?, ?)",
		c.TeamID, c.ChannelName, c.ChannelID, 0,
	)
	if err != nil {
		return c, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return c, err
	}
	c.ID = id

	return c, nil
}

// UpdateChannel updates Channel entry in database
func (m *MySQL) UpdateChannel(ch model.Channel) (model.Channel, error) {
	_, err := m.conn.Exec(
		"UPDATE `channels` SET channel_standup_time=?  WHERE id=?",
		ch.StandupTime, ch.ID,
	)
	if err != nil {
		return ch, err
	}
	var i model.Channel
	err = m.conn.Get(&i, "SELECT * FROM `channels` WHERE id=?", ch.ID)
	return i, err
}

//ListChannels returns list of channels
func (m *MySQL) ListChannels() ([]model.Channel, error) {
	channels := []model.Channel{}
	err := m.conn.Select(&channels, "SELECT * FROM `channels`")
	return channels, err
}

// SelectChannel selects Channel entry from database
func (m *MySQL) SelectChannel(channelID string) (model.Channel, error) {
	var c model.Channel
	err := m.conn.Get(&c, "SELECT * FROM `channels` WHERE channel_id=?", channelID)
	if err != nil {
		return c, err
	}
	return c, err
}

// GetChannel selects Channel entry from database with specific id
func (m *MySQL) GetChannel(id int64) ([]model.Channel, error) {
	var c []model.Channel
	err := m.conn.Select(&c, "SELECT * FROM `channels` where id=?", id)
	if err != nil {
		return c, err
	}
	return c, err
}

// DeleteChannel deletes Channel entry from database
func (m *MySQL) DeleteChannel(id int64) error {
	_, err := m.conn.Exec("DELETE FROM `channels` WHERE id=?", id)
	return err
}
