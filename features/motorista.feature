Feature: Histórico e avaliação de corridas
As a usuário
I want to acessar o histórico de corridas
So that eu possa avaliar as corridas e/ou visualizar as corridas que fiz

Scenario: Exibição do histórico de corridas para o motorista
Given que estou logado como motorista,
And estou na tela de menu principal,
When clico na opção "Histórico de Corridas",
Then vejo uma lista contendo as últimas corridas realizadas,
And cada item da lista exibe destino, data, hora e valor da corrida.

Scenario: Avaliação de corrida realizada pelo motorista
Given que estou logado como motorista,
And acessei a lista do histórico de corridas,
And existe uma corrida ainda não avaliada,
When escolho uma nota de 1 a 5 estrelas,
And clico em "Enviar avaliação",
Then vejo uma mensagem "Avaliação registrada com sucesso",
And a corrida é marcada como avaliada no histórico.

Scenario: Tentativa de avaliar uma corrida já avaliada 
Given que estou logado como motorista,
And estou na tela de histórico de corridas,
And selecionei uma corrida que já foi avaliada anteriormente,
When tento clicar no botão "Avaliar",
Then vejo a mensagem "Esta corrida já foi avaliada",
And os campos de avaliação estão desabilitados.

Scenario: Visualização de corrida com avaliação pendente
Given que estou logado como motorista,
And tenho uma corrida anterior ainda não avaliada,
When estou na tela de menu principal,
Then vejo um pop-up escrito “Avalie sua corrida”,
And o pop-up exibe os botões “Avaliar agora” e “Avaliar depois”,
And eu clico “Avaliar depois”, 
Then o pop-up é fechado temporariamente
