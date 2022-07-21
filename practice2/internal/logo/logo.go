package logo

import "strings"

func GenerateLogo() string {
	builder := strings.Builder{}

	builder.WriteString("\n")
	builder.WriteString("  _____                _   _          ___  \n")
	builder.WriteString(" |  __ \\              | | (_)        |__ \\ \n")
	builder.WriteString(" | |__) | __ __ _  ___| |_ _  ___ ___   ) |\n")
	builder.WriteString(" |  ___/ '__/ _` |/ __| __| |/ __/ _ \\ / / \n")
	builder.WriteString(" | |   | | | (_| | (__| |_| | (_|  __// /_ \n")
	builder.WriteString(" |_|   |_|  \\__,_|\\___|\\__|_|\\___\\___|____|\n")
	builder.WriteString("                                           ")
	return builder.String()
}
