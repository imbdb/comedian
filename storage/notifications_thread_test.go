package storage

import (
	"testing"
	"time"

	"github.com/maddevsio/comedian/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotification(t *testing.T) {
	var db = setupDB()
	n := model.NotificationThread{
		ChatID:           int64(1),
		Username:         "User1",
		NotificationTime: time.Now(),
		ReminderCounter:  0,
	}

	timeTest := n.NotificationTime

	notification, err := db.CreateNotificationThread(n)
	require.NoError(t, err)
	assert.Equal(t, int64(1), notification.ChatID)
	assert.Equal(t, "User1", notification.Username)
	assert.Equal(t, timeTest, notification.NotificationTime)
	assert.Equal(t, 0, notification.ReminderCounter)

	notification2, err := db.CreateNotificationThread(n)
	require.NoError(t, err)

	notifications, err := db.ListNotificationsThread(notification2.ChatID)
	require.NoError(t, err)
	assert.Equal(t, 2, len(notifications))

	err = db.DeleteNotificationThread(notification2.ID)
	require.NoError(t, err)

	err = db.DeleteNotificationThread(notification.ID)
	require.NoError(t, err)

	notifications, err = db.ListNotificationsThread(notification2.ChatID)
	require.NoError(t, err)
	assert.Equal(t, 0, len(notifications))

	n = model.NotificationThread{
		ChatID:           int64(1),
		Username:         "User2",
		NotificationTime: time.Now(),
		ReminderCounter:  0,
	}

	nt, err := db.CreateNotificationThread(n)
	require.NoError(t, err)

	err = db.UpdateNotificationThread(nt.ID, nt.ChatID, time.Now())
	require.NoError(t, err)

	notifications, err = db.ListNotificationsThread(nt.ChatID)
	require.NoError(t, err)
	for _, thread := range notifications {
		assert.Equal(t, 1, thread.ReminderCounter)
	}

	err = db.DeleteNotificationThread(nt.ID)
	require.NoError(t, err)
}
