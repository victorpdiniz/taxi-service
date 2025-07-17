# language: pt
Funcionalidade: Recuperação de Conta do Motorista
  Como um motorista que esqueceu suas credenciais
  Eu quero recuperar o acesso à minha conta
  Para voltar a trabalhar e acessar o sistema

  Contexto:
    Dado existe um motorista cadastrado com os dados:
      | email  | jose.santos@email.com |
      | status | ativo                 |
    E estou na página "Recuperação de Conta"

  Cenário: Solicitar recuperação com email válido
    Quando preencho o campo "email" com o valor "jose.santos@email.com"
    E submeto o formulário
    Então vejo a mensagem "Instruções de recuperação enviadas para jose.santos@email.com"
    E um email é enviado para "joao.silva@email.com" com o assunto "Recuperação de Conta" 

  Cenário: Tentativa de recuperação com email não cadastrado
    Dado não existe um motorista cadastrado com os dados:
      | email | nao.existe@email.com |
    Quando preencho o campo "email" com o valor "nao.existe@email.com"
    E submeto o formulário
    Então vejo a mensagem de erro "O E-mail informado não está cadastrado."

  Esquema do Cenário: Tentativa de recuperação com email inválido
    Dado não existe um motorista cadastrado com os dados:
      | email | nao.existe@email.com |
    Quando preencho o campo "email" com o valor "<valor>"
    E submeto o formulário
    Então vejo a mensagem de erro "<mensagem>"

    Exemplos:
      | valor                | mensagem                                |
      |                      | Email é obrigatório.                    |
      | email_invalido       | Formato de email inválido.              |
      | nao.existe@email.com | O E-mail informado não está cadastrado. |

  Cenário: Redefinir senha com sucesso
    Dado que estou na página "Redefinir Senha" através de link válido para "jose.santos@email.com"
    Quando preencho os campos:
      | nova_senha  | NovaSenha@123 |
      | confirmacao | NovaSenha@123 |
    E submeto o formulário
    Então vejo a mensagem "Senha redefinida com sucesso"
    E existe um motorista cadastrado com os dados:
      | email | jose.santos@email.com |
      | senha | NovaSenha@123         |

  Esquema do Cenário: Validação na redefinição de senha
    Dado que estou na página "Redefinir Senha" através de link válido para "jose.santos@email.com"
    Quando preencho os campos:
      | nova_senha  | <nova_senha>  |
      | confirmacao | <confirmacao> |
    E submeto o formulário
    Então vejo a mensagem de erro "<mensagem_erro>"

    Exemplos:
      Exemplos:
      | nova_senha     | confirmacao    | mensagem_erro                                                                            |
      | 123456         | 123456         | Senha deve ter pelo menos 8 caracteres, incluindo maiúscula, minúscula, número e símbolo |
      | NovaSenha@456  | NovaSenha@789  | Nova senha e confirmação não correspondem.                                               |
      |                | NovaSenha@789  | Nova senha é obrigatória.                                                                |  
  
  Cenário: Acesso à redefinição de senha com link inválido
    Dado que estou na página "Redefinir Senha" através de link válido para ""
    Então vejo a mensagem de erro "Este link não é válido."