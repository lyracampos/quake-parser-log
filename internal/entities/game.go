package entities

type Game struct {
	TotalKills    int            `json:"total_kills"`
	Players       []string       `json:"players"`
	PlayerKills   map[string]int `json:"kills"`
	KillByMeans   map[string]int `json:"kills_by_means"`
	PlayerRanking map[int]string `json:"player_ranking"`
}

func NewGame() *Game {
	return &Game{
		TotalKills:    0,
		Players:       make([]string, 0),
		PlayerKills:   make(map[string]int),
		KillByMeans:   make(map[string]int),
		PlayerRanking: make(map[int]string),
	}
}

func (g *Game) AddTotalKills() {
	g.TotalKills++
}

func (g *Game) AddPlayer(player string) {
	g.Players = append(g.Players, player)
	g.PlayerKills[player] = 0
}

func (g *Game) IsPlayerOnList(player string) bool {
	for _, p := range g.Players {
		if p == player {
			return true
		}
	}
	return false
}

func (g *Game) AddPlayerKill(player string) {
	g.PlayerKills[player]++
}

func (g *Game) RemovePlayerKill(player string) {
	g.PlayerKills[player]--
}

func (g *Game) AddKillMeans(killMean string) {
	g.KillByMeans[killMean]++
}
