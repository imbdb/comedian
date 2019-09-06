package storage

import (
	"testing"
	"time"

	"github.com/maddevsio/comedian/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotification(t *testing.T) {
	n := model.NotificationThread{
		ChannelID:        "1",
		UserID:           "1",
		RealName:         "User1",
		NotificationTime: int64(2),
		ReminderCounter:  0,
	}

	notification, err := db.CreateNotificationThread(n)
	require.NoError(t, err)
	assert.Equal(t, "1", notification.ChannelID)
	assert.Equal(t, "1", notification.UserID)
	assert.Equal(t, "User1", notification.RealName)
	assert.Equal(t, int64(2), notification.NotificationTime)
	assert.Equal(t, 0, notification.ReminderCounter)

	notification2, err := db.CreateNotificationThread(n)
	require.NoError(t, err)

	notifications, err := db.ListNotificationsThread(notification2.ChannelID)
	require.NoError(t, err)
	assert.Equal(t, 2, len(notifications))

	err = db.DeleteNotificationThread(notification2.ID)
	require.NoError(t, err)

	err = db.DeleteNotificationThread(notification.ID)
	require.NoError(t, err)

	notifications, err = db.ListNotificationsThread(notification2.ChannelID)
	require.NoError(t, err)
	assert.Equal(t, 0, len(notifications))

	n = model.NotificationThread{
		ChannelID:        "1",
		RealName:         "User2",
		NotificationTime: int64(2),
		ReminderCounter:  0,
	}

	nt, err := db.CreateNotificationThread(n)
	require.NoError(t, err)

	err = db.UpdateNotificationThread(nt.ID, nt.ChannelID, time.Now())
	require.NoError(t, err)

	notifications, err = db.ListNotificationsThread(nt.ChannelID)
	require.NoError(t, err)
	for _, thread := range notifications {
		assert.Equal(t, 1, thread.ReminderCounter)
	}

	err = db.DeleteNotificationThread(nt.ID)
	require.NoError(t, err)
}
