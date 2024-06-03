package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddTotalKills(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "should increase game's total kills successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.AddTotalKills()
			game.AddTotalKills()
			game.AddTotalKills()

			require.Equal(t, 3, game.TotalKills)
		})
	}
}

func TestAddPlayer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "should add three playes on game's players list successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.AddPlayer("Isgalamido")
			game.AddPlayer("Dono da Bola")
			game.AddPlayer("Mocinha")

			require.Equal(t, 3, len(game.Players))

			for i, player := range game.Players {
				require.Equal(t, game.Players[i], player)
			}
		})
	}
}

func TestIsPlayerOnList(t *testing.T) {
	tests := []struct {
		name           string
		player         string
		expectedResult bool
	}{
		{
			name:           "should return true when player is on the list",
			player:         "Isgalamido",
			expectedResult: true,
		},
		{
			name:           "should return false when player isn't on the list",
			player:         "Mocinha",
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.AddPlayer("Isgalamido")
			result := game.IsPlayerOnList(tt.player)

			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestAddPlayerKill(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "should add player's kills successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.AddPlayerKill("Isgalamido")
			game.AddPlayerKill("Isgalamido")
			game.AddPlayerKill("Isgalamido")

			require.Equal(t, 3, game.PlayerKills["Isgalamido"])
			require.Equal(t, 0, game.PlayerKills["Mocinha"])
		})
	}
}

func TestRemovePlayerKill(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "should remove player's kills successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.RemovePlayerKill("Isgalamido")
			game.RemovePlayerKill("Isgalamido")
			game.RemovePlayerKill("Isgalamido")

			require.Equal(t, -3, game.PlayerKills["Isgalamido"])
			require.Equal(t, 0, game.PlayerKills["Mocinha"])
		})
	}
}

func TestAddKillMeans(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "should add kiil's means successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.AddKillMeans("MOD_TRIGGER_HURT")
			game.AddKillMeans("MOD_TRIGGER_HURT")
			game.AddKillMeans("MOD_TRIGGER_HURT")
			game.AddKillMeans("MOD_ROCKET_SPLASH")
			game.AddKillMeans("MOD_ROCKET_SPLASH")

			require.Equal(t, 3, game.KillByMeans["MOD_TRIGGER_HURT"])
			require.Equal(t, 2, game.KillByMeans["MOD_ROCKET_SPLASH"])
		})
	}
}
