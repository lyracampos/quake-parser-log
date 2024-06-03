package usecases

import (
	"context"
	"fmt"
	"os"

	"github.com/lyracampos/quake-log-parser/internal/entities"
	"github.com/lyracampos/quake-log-parser/internal/parser"
)

type ParseLogUseCase struct {
}

func NewParseLogUseCase() *ParseLogUseCase {
	return &ParseLogUseCase{}
}

func (u *ParseLogUseCase) Execute(ctx context.Context, filePath string) (*entities.GameList, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return &entities.GameList{}, fmt.Errorf("failed on reading log file: %w", err)
	}

	games, err := parser.ParseLog(string(content))
	if err != nil {
		return &entities.GameList{}, fmt.Errorf("failed on parsing raw log: %w", err)
	}

	return &entities.GameList{
		Games: games,
	}, nil
}
