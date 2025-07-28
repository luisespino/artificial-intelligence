package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// bot responses
var botResponses = map[string]string{
	"greeting":  "¡Hola! ¿Cómo puedo ayudarte hoy?",
	"how":       "Estoy bien, gracias por preguntar. ¿Y tú?",
	"name":      "Soy un chatbot creado para ayudarte.",
	"weather":   "No sé nada del clima, pero puedo ayudarte con preguntas.",
	"bye":       "¡Adiós! ¡Que tengas un buen día!",
	"thanks":    "¡De nada! Si necesitas algo más, pregúntame.",
	"questions": "Puedo responderte a palabras clave.",
	"chatbot":   "Es software que simula mantener una conversación.",
	"ai":        "Es un campo de ciencias de la computación que resuelve\n un problema de manera automática mediante la lógica.",
}

// synonymGroups define groups of synonyms for different intents
var synonymGroups = map[string][]string{
	"greeting":  {"hola", "hey", "buenas"},
	"how":       {"como estas", "cómo estás", "que tal", "qué tal"},
	"name":      {"cómo te llamas", "quién eres", "como te llamas", "quien eres"},
	"weather":   {"cómo está el clima", "clima", "qué tiempo hace", "tiempo"},
	"bye":       {"adiós", "hasta luego", "adios"},
	"thanks":    {"gracias", "muchas gracias", "ok"},
	"questions": {"qué preguntas", "qué tipo de preguntas", "qué puedes hacer", "que puedes hacer", "preguntas"},
	"chatbot":   {"chatbot", "chat"},
	"ai":        {"ia", "inteligencia artificial"},
}

// getBotResponse returns a response based on the input string
func getBotResponse(input string) string {
	input = strings.ToLower(input)
	for key, synonyms := range synonymGroups {
		for _, synonym := range synonyms {
			if strings.Contains(input, synonym) {
				return botResponses[key]
			}
		}
	}
	return "Lo siento, no entendí eso. ¿Puedes preguntar de otra manera?"
}

// main function initializes the Fyne app and sets up the UI
func main() {
	a := app.New()
	w := a.NewWindow("ChatBot")

	// Label multiline para mensajes
	chatLabel := widget.NewLabel("")
	chatLabel.Wrapping = fyne.TextWrapWord
	chatLabel.TextStyle = fyne.TextStyle{Monospace: false}

	// Scrollable container for chat messages
	scroll := container.NewVScroll(chatLabel)
	scroll.SetMinSize(fyne.NewSize(480, 300))

	// Input entry for user messages
	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Escribe tu mensaje...")

	// messages send button
	sendButton := widget.NewButton("Enviar", func() {
		userInput := strings.TrimSpace(inputEntry.Text)
		if userInput == "" {
			return
		}

		chatLabel.SetText(chatLabel.Text + "Tú: " + userInput + "\n")
		response := getBotResponse(userInput)
		chatLabel.SetText(chatLabel.Text + "Bot: " + response + "\n")
		inputEntry.SetText("")

		// Forzar actualización y scroll automático
		chatLabel.Refresh()
		scroll.ScrollToBottom()
	})

	// enter key submits the input
	inputEntry.OnSubmitted = func(string) {
		sendButton.OnTapped()
	}

	// Layout for input area
	inputArea := container.NewBorder(nil, nil, nil, sendButton, inputEntry)

	// Main content layout
	content := container.NewVBox(
		scroll,
		layout.NewSpacer(),
		inputArea,
	)

	// Main window content
	w.SetContent(content)
	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun()
}
