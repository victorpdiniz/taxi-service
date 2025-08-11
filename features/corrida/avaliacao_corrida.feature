# language: pt
Funcionalidade: Avaliação de Corridas
  Como um motorista
  Eu quero avaliar minhas corridas concluídas
  Para registrar minha experiência

  Contexto:
    Dado existe uma corrida concluída com os dados:
      | dado        | valor       |
      | ID          | 101         |
      | MotoristaID | 123         |
      | Status      | CONCLUIDA_NO_TEMPO |

  Cenário: Avaliação bem-sucedida
    Dado estou autenticado como motorista ID 123
    E estou na página "Histórico de Corridas"
    Quando avalio a corrida de ID 101 com nota 5
    Então vejo a mensagem "Avaliação registrada com sucesso"
    E a corrida de ID 101 tem avaliação 5

  Esquema do Cenário: Avaliação inválida
    Dado estou autenticado como motorista ID 123
    E estou na página "Histórico de Corridas"
    Quando avalio a corrida de ID 101 com nota 3
    Então vejo a mensagem de erro "Não foi possível registrar sua avaliação"

    Exemplos:
      | nota | mensagem                        |
      | -1   | Nota deve ser entre 1 e 5       |
      | 0    | Nota deve ser entre 1 e 5       |
      | 6    | Nota deve ser entre 1 e 5       |
      |      | Nota é obrigatória              |

  Cenário: Tentativa de avaliar corrida inexistente
    Dado estou autenticado como motorista ID 123
    Quando avalio a corrida de ID 999 com nota 4
    Então vejo a mensagem de erro "Corrida não encontrada"
