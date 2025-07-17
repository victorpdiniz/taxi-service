# language: pt
Funcionalidade: Gerenciamento de Perfil do Motorista
  Como um motorista autenticado
  Eu quero gerenciar meu perfil e dados pessoais
  Para manter minhas informações sempre atualizadas

  Contexto:
    Dado existe um motorista cadastrado com os dados:
      | dado            | valor                        |
      | nome            | João Silva                   |
      | cpf             | 221.623.340-46               |
      | cnh             | 12345678901                  |
      | categoria_cnh   | B                            |
      | validade_cnh    | 15/03/2030                   |
      | placa_veiculo   | ABC1234                      |
      | modelo_veiculo  | Honda Civic 2020             |
      | telefone        | (11) 99999-9999              |
      | email           | joao.silva@email.com         |
    E estou autenticado como "joao.silva@email.com"
    E eu estou na página "Painel do Motorista"

  Cenário: Visualizar perfil completo
    Então vejo meu dados:
      | dado            | valor                        |
      | nome            | João Silva                   |
      | cpf             | 221.623.340-46               |
      | cnh             | 12345678901                  |
      | categoria_cnh   | B                            |
      | validade_cnh    | 15/03/2030                   |
      | placa_veiculo   | ABC1234                      |
      | modelo_veiculo  | Honda Civic 2020             |
      | telefone        | (11) 99999-9999              |
      | email           | joao.silva@email.com         |
    E vejo minha foto de perfil

  Esquema do Cenário: Editar dados pessoais básicos com sucesso
    Quando preencho o campo "<campo>" com o valor "<valor>"
    E tento salvar as alterações
    Então vejo a mensagem "Perfil atualizado com sucesso"
    E vejo o campo "<campo>" com valor "<valor>"

    Exemplos:
      | campo    | valor                  |
      | telefone | (11) 88888-8888        |
      | email    | joao.s.silva@email.com |

  # TODO: Editar dados de documentos

  Esquema do Cenário: Validação de campos editáveis
    Quando preencho o campo "<campo>" com o valor "<valor_invalido>"
    E tento salvar as alterações
    Então vejo a mensagem de erro "<mensagem_erro>"

    Exemplos:
      | campo    | valor_invalido | mensagem_erro                 |
      | telefone | 999999999      | Formato de telefone inválido. |
      | telefone | (11) 9999-999  | Formato de telefone inválido. |
      | telefone |                | Telefone é obrigatório.       |
      | email    | email_invalido | Formato de email inválido.    |
      | email    |                | E-mail é obrigatório.         |

  Cenário: Alterar senha com sucesso
    Dado vejo a mensagem "Alteração de Senha"
    Quando preencho os campos:
      | senha_atual | MinhaSenh@123 |
      | nova_senha  | NovaSenha@456 |
      | confirmacao | NovaSenha@456 |
    E tento salvar as alterações
    Então vejo a mensagem "Senha alterada com sucesso"
    E existe um motorista cadastrado com os dados:
      | dado   | valor                |
      | email  | joao.silva@email.com |
      | senha  | NovaSenha@456        |

  Esquema do Cenário: Validação na alteração de senha
    Dado vejo a mensagem "Alteração de Senha"
    Quando preencho os campos:
      | senha_atual | <senha_atual> |
      | nova_senha  | <nova_senha>  |
      | confirmacao | <confirmacao> |
    E tento salvar as alterações
    Então vejo a mensagem de erro "<mensagem_erro>"

    Exemplos:
      | senha_atual    | nova_senha     | confirmacao    | mensagem_erro                                                                            |
      | senha_errada   | NovaSenha@456  | NovaSenha@456  | Senha atual incorreta.                                                                   |
      | MinhaSenh@123  | 123456         | 123456         | Senha deve ter pelo menos 8 caracteres, incluindo maiúscula, minúscula, número e símbolo |
      | MinhaSenh@123  | NovaSenha@456  | NovaSenha@789  | Nova senha e confirmação não correspondem.                                               |
      |                | NovaSenha@456  | NovaSenha@789  | Senha atual é obrigatória.                                                               |
      | MinhaSenh@123  |                | NovaSenha@789  | Nova senha é obrigatória.                                                                |  

  Cenário: Upload de nova foto de perfil válida
    Dado vejo a mensagem "Upload de nova foto de perfil"
    Quando tento fazer upload de arquivo com tamanho "1MB" e formato "JPG"
    Então vejo a mensagem "Foto de perfil atualizada com sucesso"

  Esquema do Cenário: Upload de foto inválida
    Dado vejo a mensagem "Upload de nova foto de perfil"
    Quando tento fazer upload de arquivo com tamanho "<tamanho>" e formato "<formato>"
    Então vejo a mensagem de erro "<mensagem>"

    Exemplos:
      | formato | tamanho         | mensagem                                     |
      | JPG     | 8MB             | Foto muito grande. Tamanho máximo: 5MB       |
      | TXT     | 1MB             | Formato não suportado. Use JPG, PNG ou WEBP  |

  Cenário: Solicitar exclusão permanente da conta
    Quando solicito excluir minha conta permanentemente
    Então um email é enviado para "joao.silva@email.com" com o assunto "Solicitação de Exclusão de Conta" 

  Cenário: Aceitar solicitação de exclusão
    Dado que o status de ""joao.silva@email.com" é "aguardando_exclusao"
    Quando eu estou na página "Confirmar Solicitação de Exclusão" através de link válido para "jose.santos@email.com"
    Então o status de ""joao.silva@email.com" é "encerrado"
    E eu vejo a mensagem "Sua conta foi encerrada e será excluida permanentemente em 72h"
    E não estou autenticado como motorista