package godog
// var (
//     motoristaLogadoID int
//     historicoCorridas []models.Corrida
//     corridaSelecionada models.Corrida
//     mensagemRetornada string
// )
// func existemAsSeguintesCorridasConcluidas(tabela *godog.Table) error {
//     historicoCorridas = []models.Corrida{}
//     for _, row := range tabela.Rows[1:] {
//         id, _ := strconv.Atoi(row.Cells[0].Value)
//         motoristaID, _ := strconv.Atoi(row.Cells[1].Value)
//         preco, _ := strconv.ParseFloat(row.Cells[3].Value, 64)
//         tempoEstimado, _ := strconv.Atoi(row.Cells[4].Value)
//         tempoDecorrido, _ := strconv.Atoi(row.Cells[5].Value)

//         historicoCorridas = append(historicoCorridas, models.Corrida{
//             ID:              id,
//             MotoristaID:     motoristaID,
//             Status:          row.Cells[2].Value,
//             Preco:           preco,
//             TempoEstimado:   tempoEstimado,
//             TempoDecorrido:  tempoDecorrido,
//         })
//     }
//     return nil
// }

// func estouAutenticadoComoMotorista(id int) error {
//     motoristaLogadoID = id
//     return nil
// }

// func estouNaPaginaHistoricoDeCorridas() error {
//     return nil // simula que está na página
// }

// func solicitoOHistoricoDeCorridas() error {
//     var corridasFiltradas []models.Corrida
//     for _, c := range historicoCorridas {
//         if c.MotoristaID == motoristaLogadoID && strings.HasPrefix(c.Status, "CONCLUIDA") {
//             corridasFiltradas = append(corridasFiltradas, c)
//         }
//     }

//     if len(corridasFiltradas) == 0 {
//         mensagemRetornada = "Nenhuma corrida encontrada"
//     } else {
//         historicoCorridas = corridasFiltradas
//         mensagemRetornada = fmt.Sprintf("Encontradas %d corridas", len(corridasFiltradas))
//     }

//     return nil
// }

// func vejoAListaComNCorridas(quantidade int) error {
//     if len(historicoCorridas) != quantidade {
//         return fmt.Errorf("esperava %d corridas, mas encontrou %d", quantidade, len(historicoCorridas))
//     }
//     return nil
// }

// func todasTemStatusConcluida() error {
//     for _, c := range historicoCorridas {
//         if !strings.HasPrefix(c.Status, "CONCLUIDA") {
//             return fmt.Errorf("corrida com ID %d tem status %s", c.ID, c.Status)
//         }
//     }
//     return nil
// }

// func vejoAMensagemEsperada(msg string) error {
//     if mensagemRetornada != msg {
//         return fmt.Errorf("esperava mensagem: '%s', mas recebeu: '%s'", msg, mensagemRetornada)
//     }
//     return nil
// }

// func visualizoDetalhesDaCorridaID(id int) error {
//     for _, c := range historicoCorridas {
//         if c.ID == id {
//             corridaSelecionada = c
//             return nil
//         }
//     }
//     return fmt.Errorf("corrida com ID %d não encontrada", id)
// }

// func statusDaCorridaSelecionadaEh(status string) error {
//     if corridaSelecionada.Status != status {
//         return fmt.Errorf("esperava status '%s', recebeu '%s'", status, corridaSelecionada.Status)
//     }
//     return nil
// }

// func precoFinalDaCorridaSelecionadaEh(valor float64) error {
//     if corridaSelecionada.Preco != valor {
//         return fmt.Errorf("esperava preço %.2f, recebeu %.2f", valor, corridaSelecionada.Preco)
//     }
//     return nil
// }
// func InitializeScenario(ctx *godog.ScenarioContext) {
//     ctx.Step(`^existem as seguintes corridas concluídas:$`, existemAsSeguintesCorridasConcluidas)
//     ctx.Step(`^estou autenticado como motorista ID (\d+)$`, estouAutenticadoComoMotorista)
//     ctx.Step(`^estou na página "Histórico de Corridas"$`, estouNaPaginaHistoricoDeCorridas)
//     ctx.Step(`^solicito o histórico de corridas$`, solicitoOHistoricoDeCorridas)
//     ctx.Step(`^vejo a lista com (\d+) corridas$`, vejoAListaComNCorridas)
//     ctx.Step(`^todas têm o status "concluída"$`, todasTemStatusConcluida)
//     ctx.Step(`^vejo a mensagem "([^"]*)"$`, vejoAMensagemEsperada)
//     ctx.Step(`^visualizo os detalhes da corrida de ID (\d+)$`, visualizoDetalhesDaCorridaID)
//     ctx.Step(`^vejo que o status é "([^"]*)"$`, statusDaCorridaSelecionadaEh)
//     ctx.Step(`^o preço final é ([\d.]+)$`, precoFinalDaCorridaSelecionadaEh)
// }
