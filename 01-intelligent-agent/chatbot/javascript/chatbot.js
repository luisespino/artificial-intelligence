const botResponses = {
    "greeting": "¡Hola! ¿Cómo puedo ayudarte hoy?",
    "how": "Estoy bien, gracias por preguntar. ¿Y tú?",
    "name": "Soy un chatbot creado para ayudarte.",
    "weather": "No sé nada del clima, pero puedo ayudarte con preguntas.",
    "bye": "¡Adiós! ¡Que tengas un buen día!",
    "thanks": "¡De nada! Si necesitas algo más, pregúntame.",
    "questions": "Puedo responderte a palabras clave.",
    "chatbot": "Es software que simula mantener una conversación.",
    "ai": "Es un campo de ciencias de la computación que resuelve un problema de manera automática mediante la lógica."
};

const synonymGroups = {
    greeting: ["hola", "hey", "buenas"],
    how: ["como estas", "cómo estás", "que tal", "qué tal"],
    name: ["cómo te llamas", "quién eres", "como te llamas", "quien eres"],
    weather: ["cómo está el clima", "clima", "qué tiempo hace", "tiempo"],
    bye: ["adiós", "hasta luego", "adios",],
    thanks: ["gracias", "muchas gracias", "ok"],
    questions: ["qué preguntas", "qué tipo de preguntas", "qué puedes hacer", "que puedes hacer", "preguntas"],
    chatbot: ["chatbot", "chat"],
    ai: ["ia", "inteligencia artificial"]
};

function processUserInput() {
    const userInput = document.getElementById("user-input").value.trim().toLowerCase();

    if (userInput === "") {
        return;
    }

    displayMessage(userInput, "user");

    const response = getBotResponse(userInput);
    
    displayMessage(response, "bot");

    document.getElementById("user-input").value = "";
}

function displayMessage(message, sender) {
    const chatBox = document.getElementById("chat-box");

    const messageElement = document.createElement("div");
    messageElement.classList.add("message");

    if (sender === "user") {
        messageElement.classList.add("user-message");
        messageElement.textContent = "Tú: " + message;
    } else {
        messageElement.classList.add("bot-message");
        messageElement.textContent = "Bot: " + message;
    }

    chatBox.appendChild(messageElement);

    chatBox.scrollTop = chatBox.scrollHeight;
}

function getBotResponse(userInput) {
    const input = userInput.toLowerCase();
    
    for (const key in synonymGroups) {
        for (const synonym of synonymGroups[key]) {
            if (input.includes(synonym)) {
                return botResponses[key];
            }
        }
    }

    return "Lo siento, no entendí eso. ¿Puedes preguntar de otra manera?";
}

document.getElementById("user-input").addEventListener("keypress", function(event) {
    if (event.key === "Enter") {
        processUserInput();
    }
});

