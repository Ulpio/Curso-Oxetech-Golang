// Exemplo simples de uso de canais com buffer em Go.
package main

import "fmt"

// main demonstra como enviar e receber dados de um canal com buffer.
func main() {
	canal := make(chan string, 2) // Canal comporta até dois elementos sem bloquear.

	canal <- "Ulpio" // Envia primeira mensagem.
	canal <- "Netto" // Envia segunda mensagem; ainda não bloqueia pois há buffer.

	fmt.Println(<-canal) // Lê o primeiro valor enviado (FIFO): "Ulpio".
	fmt.Println(<-canal) // Lê o segundo valor: "Netto".
}
