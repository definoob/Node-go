package services_test

import (
	"testing"
	"time"

	"github.com/vordev/VOR/core/internal/cltest"
	"github.com/vordev/VOR/core/services"
	"github.com/vordev/VOR/core/store/models"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreReaper_ReapSessions(t *testing.T) {
	t.Parallel()

	store, cleanup := cltest.NewStore(t)
	defer cleanup()

	r := services.NewStoreReaper(store)
	defer r.Stop()

	tests := []struct {
		name     string
		lastUsed time.Time
		wantReap bool
	}{
		{"current", time.Now(), false},
		{"expired", time.Now().Add(-store.Config.SessionTimeout().Duration()), false},
		{"almost stale", time.Now().Add(-store.Config.ReaperExpiration().Duration()), false},
		{"stale", time.Now().Add(-store.Config.ReaperExpiration().Duration()).
			Add(-store.Config.SessionTimeout().Duration()), true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer store.ORM.ClearSessions()

			session := cltest.NewSession(test.name)
			session.LastUsed = test.lastUsed
			require.NoError(t, store.SaveSession(&session))

			r.WakeUp()

			if test.wantReap {
				gomega.NewGomegaWithT(t).Eventually(func() []models.Session {
					sessions, err := store.Sessions(0, 10)
					assert.NoError(t, err)
					return sessions
				}).Should(gomega.HaveLen(0))
			} else {
				gomega.NewGomegaWithT(t).Consistently(func() []models.Session {
					sessions, err := store.Sessions(0, 10)
					assert.NoError(t, err)
					return sessions
				}).Should(gomega.HaveLen(1))
			}
		})
	}
}
