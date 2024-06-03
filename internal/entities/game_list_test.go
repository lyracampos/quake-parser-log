package entities

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddGame(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "should add two games to list successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewGameList()
			list.AddGame(NewGame())
			list.AddGame(NewGame())

			require.Equal(t, 2, len(list.Games))
			count := 1
			for i := range list.Games {
				expectedName := fmt.Sprintf("game_%02d", count)
				require.Equal(t, i, expectedName)
				count++
			}
		})
	}
}
