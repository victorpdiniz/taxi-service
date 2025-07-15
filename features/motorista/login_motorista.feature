# language: pt
Funcionalidade: Login do Motorista
  Como um motorista
  Eu quero me autenticar no sistema
  Para poder trabalhar

  Contexto:
    Dado que existe um motorista cadastrado com os dados:
      | email  | jose.santos@email.com   |
      | senha  | MinhaSenh@123           |

  Cenário: Tentativa de login com diferentes status e credenciais
    Quando faço login com "joao.silva@email.com" e "<senha>"
    E meu status é "<status>"
    Então <resultado>

  Exemplos:
    | status    | senha             | resultado                                                                                  |
    | rejeitado | MinhaSenh@123     | estou na página "Upload de Documentos"                                                     |
    | ativo     | MinhaSenh@123     | estou na página "Painel do Motorista"                                                      |
    | ativo     | SenhaErrada123    | vejo a mensagem de erro "Email ou senha inválidos."                                        |
    | encerrado | MinhaSenh@123     | vejo a mensagem de erro "Conta encerrada. Entre em contato com o suporte."                  |