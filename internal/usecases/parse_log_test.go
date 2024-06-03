package usecases

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name             string
		filePath         string
		expectError      bool
		expectedMsgError string
	}{
		{
			name:             "should return success when parsing log successfully",
			filePath:         "../../static/logs/test_valid.log",
			expectError:      false,
			expectedMsgError: "",
		},
		{
			name:             "should return an error when unable to read the file",
			filePath:         "not_existent_file.log",
			expectError:      true,
			expectedMsgError: "failed on reading log file: open not_existent_file.log: no such file or directory",
		},
		{
			name:             "should return an error when the log cannot be parsed",
			filePath:         "../../static/logs/test_empty_file.log",
			expectError:      true,
			expectedMsgError: "failed on parsing raw log: could not parse an empty file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			usecase := NewParseLogUseCase()
			games, err := usecase.Execute(ctx, tt.filePath)

			if tt.expectError {
				require.Error(t, err)
				require.ErrorContains(t, err, tt.expectedMsgError)
			}

			if !tt.expectError {
				require.Equal(t, 3, len(games.Games))

				count := 1
				for i := range games.Games {
					expectedName := fmt.Sprintf("game_%02d", count)
					require.Equal(t, i, expectedName)
					count++
				}
			}
		})
	}
}
