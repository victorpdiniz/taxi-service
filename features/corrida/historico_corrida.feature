# language: pt
Funcionalidade: Histórico de Corridas
  Como um motorista
  Eu quero ver o histórico das minhas corridas
  Para acompanhar meu desempenho e atividades

  Contexto:
    Dado existem as seguintes corridas concluídas:
      | ID  | MotoristaID | Status                  | Preco | TempoEstimado | TempoDecorrido |
      | 201 | 321         | CONCLUIDA_NO_TEMPO      | 50.0  | 15             | 15              |
      | 202 | 321         | CONCLUIDA_ANTECEDENCIA  | 60.0  | 20             | 15              |
      | 203 | 321         | CONCLUIDA_NO_TEMPO      | 40.0  | 10             | 10              |

  Cenário: Visualizar todas as corridas concluídas
    Dado estou autenticado como motorista ID 321
    E estou na página "Histórico de Corridas"
    Quando solicito o histórico de corridas
    Então vejo a lista com 3 corridas
    E todas têm o status "concluída"

  Cenário: Visualizar histórico vazio
    Dado estou autenticado como motorista ID 999
    E estou na página "Histórico de Corridas"
    Quando solicito o histórico de corridas
    Então vejo a mensagem "Nenhuma corrida encontrada"

  Cenário: Visualizar detalhes de uma corrida específica
    Dado estou autenticado como motorista ID 321
    E estou na página "Histórico de Corridas"
    Quando visualizo os detalhes da corrida de ID 202
    Então vejo que o status é "concluída com antecedência"
    E o preço final é 60.0
