package godog

// import (
// 	"fmt"
// 	"strconv"
// 	"testing"
// 	"taxi_service/models"
// 	"taxi_service/services"

// 	"github.com/cucumber/godog"
// )

// var corrida models.Corrida
// var avaliacaoMensagem string

// func existeUmaCorridaConcluidaComOsDados(tabela *godog.Table) error {
// 	corrida = models.Corrida{ID: 101, MotoristaID: 123, Status: "CONCLUIDA_NO_TEMPO"}
// 	services.GetCorridas() // ou services.AdicionarCorrida(corrida)
// 	return nil
// }

// func estouAutenticadoComoMotoristaID(id int) error {
// 	if id != corrida.MotoristaID {
// 		return fmt.Errorf("motorista não autenticado")
// 	}
// 	return nil
// }

// func estouNaPagina(pagina string) error {
// 	return nil // simulado
// }

// func avalioACorridaDeIDComNota(id, nota int) error {
// 	err := services.AvaliarCorrida(id, nota)
// 	if err != nil {
// 		avaliacaoMensagem = err.Error()
// 		return err
// 	}
// 	avaliacaoMensagem = "Avaliação registrada com sucesso"
// 	return nil
// }

// func vejoAMensagemEsperada(msg string) error {
// 	if avaliacaoMensagem != msg {
// 		return fmt.Errorf("mensagem esperada '%s', mas foi '%s'", msg, avaliacaoMensagem)
// 	}
// 	return nil
// }

// func aCorridaTemAvaliacaoEsperada(id, nota int) error {
// 	for _, c := range services.GetCorridas() {
// 		if c.ID == id && c.Avaliacao != nil && *c.Avaliacao == nota {
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("avaliação da corrida %d não foi encontrada", id)
// }

// func avalioCorridaComNota(id, nota int) error {
// 	err := services.AvaliarCorrida(id, nota)
// 	if err != nil {
// 		avaliacaoMensagem = err.Error()
// 		return nil // não falha agora, para permitir verificação no Then
// 	}
// 	avaliacaoMensagem = "Avaliação registrada com sucesso"
// 	return nil
// }

// func vejoMensagemErroEsperada(msg string) error {
// 	if avaliacaoMensagem != msg {
// 		return fmt.Errorf("esperava mensagem: '%s', mas recebeu: '%s'", msg, avaliacaoMensagem)
// 	}
// 	return nil
// }



// //Registrar os steps
// func InitializeScenario(ctx *godog.ScenarioContext) {
// 	ctx.Step(`^existe uma corrida concluída com os dados:$`, existeUmaCorridaConcluidaComOsDados)
// 	ctx.Step(`^estou autenticado como motorista ID (\d+)$`, estouAutenticadoComoMotoristaID)
// 	ctx.Step(`^estou na página "([^"]*)"$`, estouNaPagina)
// 	ctx.Step(`^avalio a corrida de ID (\d+) com nota (\d+)$`, avalioCorridaComNota)
// 	ctx.Step(`^vejo a mensagem "([^"]*)"$`, vejoAMensagemEsperada)
// 	ctx.Step(`^a corrida de ID (\d+) tem avaliação (\d+)$`, aCorridaTemAvaliacaoEsperada)
// 	ctx.Step(`^vejo a mensagem de erro "([^"]*)"$`, vejoMensagemErroEsperada)
// }
