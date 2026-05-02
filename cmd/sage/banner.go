package sage

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

func printBanner() {
	startColor, _ := colorful.Hex("#225A40")
	endColor, _ := colorful.Hex("#67E0A8")

	logo := `

   ▄▄▄  ▀▀█                          ▄▄▄▄
 ▄▀   ▀   █     ▄▄▄   ▄ ▄▄    ▄▄▄   █▀   ▀  ▄▄▄    ▄▄▄▄   ▄▄▄
 █        █    █▀ ▀█  █▀  █  █▀  █  ▀█▄▄▄  ▀   █  █▀ ▀█  █▀  █
 █        █    █   █  █   █  █▀▀▀▀      ▀█ ▄▀▀▀█  █   █  █▀▀▀▀
  ▀▄▄▄▀   ▀▄▄  ▀█▄█▀  █   █  ▀█▄▄▀  ▀▄▄▄█▀ ▀▄▄▀█  ▀█▄▀█  ▀█▄▄▀
                                                   ▄  █
                                                    ▀▀
`

	lines := strings.Split(logo, "\n")
	totalChars := 0
	for _, line := range lines {
		totalChars += len([]rune(line))
	}

	charIndex := 0
	for _, line := range lines {
		var sb strings.Builder
		for _, ch := range line {
			ratio := float64(charIndex) / float64(totalChars)
			color := startColor.BlendLuv(endColor, ratio).Hex()
			style := lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Bold(true)
			sb.WriteString(style.Render(string(ch)))
			charIndex++
		}
		fmt.Println(sb.String())
	}
}