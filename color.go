package lipbalm

import "strings"

const (
	ansi_esc   byte   = '\x1b'
	ansi_end   byte   = 'm'
	ansi_fg    uint8  = 38
	ansi_bg    uint8  = 48
	ansi_reset string = "\x1b[0m"
	ansi_256   byte   = '5'
	ansi_tc    byte   = '2'
)

// sets the color of the string to the color specified
// can be background/foreground
func SetColor(color string, str string) string {
	var (
		lines       = strings.Split(str, "\n")
		lastLineIdx = len(lines) - 1

		clr string = color
		sb  strings.Builder
	)

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		if line[0] == ansi_esc {
			clr = color[:len(color)-1] + ";" // remove end
			line = line[2 : len(line)-4]     // remove esc & reset
		}

		sb.WriteString(clr)
		sb.WriteString(line)
		sb.WriteString(ansi_reset)

		if i == lastLineIdx {
			break
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}

// 256 color code
func Color(code uint8) string {
	var sb strings.Builder

	sb.WriteByte(ansi_esc)
	sb.WriteByte('[')
	writeBytes(&sb, getCodeBytes(ansi_fg))
	sb.WriteByte(';')
	sb.WriteByte(ansi_256)
	sb.WriteByte(';')
	writeBytes(&sb, getCodeBytes(code))
	sb.WriteByte(ansi_end)

	return sb.String()
}

// 256 color code Background
func ColorBg(code uint8) string {
	var sb strings.Builder

	sb.WriteByte(ansi_esc)
	sb.WriteByte('[')
	writeBytes(&sb, getCodeBytes(ansi_bg))
	sb.WriteByte(';')
	sb.WriteByte(ansi_256)
	sb.WriteByte(';')
	writeBytes(&sb, getCodeBytes(code))
	sb.WriteByte(ansi_end)

	return sb.String()
}

// true color / RGB
func ColorRGB(R, G, B uint8) string {
	var sb strings.Builder

	sb.WriteByte(ansi_esc)
	sb.WriteByte('[')
	writeBytes(&sb, getCodeBytes(ansi_fg))
	sb.WriteByte(';')
	sb.WriteByte(ansi_tc)
	sb.WriteByte(';')
	writeBytes(&sb, getCodeBytes(R))
	sb.WriteByte(';')
	writeBytes(&sb, getCodeBytes(G))
	sb.WriteByte(';')
	writeBytes(&sb, getCodeBytes(B))
	sb.WriteByte(ansi_end)

	return sb.String()
}

// true color / RGB Background
func ColorBgRGB(R, G, B uint8) string {
	var sb strings.Builder

	sb.WriteByte(ansi_esc)
	sb.WriteByte('[')
	writeBytes(&sb, getCodeBytes(ansi_bg))
	sb.WriteByte(';')
	sb.WriteByte(ansi_tc)
	sb.WriteByte(';')
	writeBytes(&sb, getCodeBytes(R))
	sb.WriteByte(';')
	writeBytes(&sb, getCodeBytes(G))
	sb.WriteByte(';')
	writeBytes(&sb, getCodeBytes(B))
	sb.WriteByte(ansi_end)

	return sb.String()
}

func getCodeBytes(code uint8) []byte {
	if code == 0 {
		return []byte{'0'}
	}

	var digits []byte

	for code > 0 {
		digits = append(digits, byte('0'+code%10))
		code /= 10
	}

	return digits
}

func writeBytes(sb *strings.Builder, bytes []byte) {
	for i := len(bytes) - 1; i >= 0; i-- {
		sb.WriteByte(bytes[i])
	}
}
