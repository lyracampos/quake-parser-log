package parser

import (
	"errors"
	"regexp"
	"strings"

	"github.com/lyracampos/quake-log-parser/internal/entities"
)

const (
	INIT_GAME                = "InitGame:"
	KILL                     = "Kill:"
	CLIENT_USER_INFO_CHANGED = "ClientUserinfoChanged:"
)

var ErrFileLogIsEmpty = errors.New("could not parse an empty file")

func ParseLog(rawLog string) (map[string]*entities.Game, error) {
	gameList := entities.NewGameList()
	lines := strings.Split(rawLog, "\n")

	if len(lines) <= 0 || lines[0] == "" {
		return map[string]*entities.Game{}, ErrFileLogIsEmpty
	}

	for lineNumber, line := range lines {
		line = strings.TrimSpace(line)
		chunks := strings.Split(line, " ")

		if isInitGame(chunks) {
			game := entities.NewGame()
			gameList.AddGame(game)
			parseGameData(game, lines, lineNumber+1)
		}
	}

	return gameList.Games, nil
}

func isInitGame(chunks []string) bool {
	if len(chunks) <= 2 {
		return false
	}

	return chunks[1] == INIT_GAME
}

func parseGameData(game *entities.Game, lines []string, lineNumber int) {
	for lineNumber < len(lines) {
		var line string = strings.TrimSpace(lines[lineNumber])
		var chunks []string = strings.Split(line, " ")

		if len(chunks) > 1 {
			switch chunks[1] {
			case CLIENT_USER_INFO_CHANGED:
				handlePlayers(game, chunks)
			case KILL:
				handleKills(game, chunks)
			case INIT_GAME:
				handlePlayerRanking(game)
				return
			}
		}

		lineNumber++
	}
}

func handlePlayers(game *entities.Game, chunks []string) {
	regex := regexp.MustCompile(`[^\\n](\w*|\w* )*`)
	player := regex.FindString(strings.Join(chunks[3:], " "))

	if len(player) > 1 {
		if !game.IsPlayerOnList(player) {
			game.AddPlayer(player)
		}
	}
}

func handleKills(game *entities.Game, chunks []string) {
	game.AddTotalKills()

	killRegexp := regexp.MustCompile(`Kill: \d+ \d+ \d+: (.+) killed (.+) by (.+)`)

	line := strings.Join(chunks, " ")

	matches := killRegexp.FindStringSubmatch(line)
	if len(matches) == 4 {
		killer := matches[1]
		victim := matches[2]
		meansOfDeath := matches[3]

		if killer != "<world>" {
			game.AddPlayerKill(killer)
		} else {
			game.RemovePlayerKill(victim)
		}

		game.AddKillMeans(meansOfDeath)
	}
}

func handlePlayerRanking(game *entities.Game) {
	playersRanking := make([]string, len(game.Players))
	copy(playersRanking, game.Players)

	for i := 1; i < len(playersRanking); i++ {
		key := playersRanking[i]
		j := i - 1
		for j >= 0 && game.PlayerKills[playersRanking[j]] < game.PlayerKills[key] {
			playersRanking[j+1] = playersRanking[j]
			j--
		}
		playersRanking[j+1] = key
	}

	for i, player := range playersRanking {
		game.PlayerRanking[i+1] = player
	}
}
