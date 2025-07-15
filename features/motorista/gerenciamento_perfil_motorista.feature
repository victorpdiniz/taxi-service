# language: pt
Funcionalidade: Gerenciamento de Perfil do Motorista
  Como um motorista autenticado
  Eu quero gerenciar meu perfil e dados pessoais
  Para manter minhas informações sempre atualizadas

  Contexto:
    Dado que estou autenticado como motorista:
      | nome           | João Silva               |
      | email          | joao.silva@email.com     |
      | cpf            | 123.456.789-00           |
      | telefone       | (11) 99999-9999          |
      | data_nascimento| 15/03/1990               |
      | cnh            | 12345678901              |
      | status         | ativo                    |
      | avaliacao      | 4.8                      |
    E eu estou na página "Perfil"

  Cenário: Visualizar perfil completo
    Então vejo meu dado "<dado>" com o valor "<valor>":
      | dado           | valor                    |
      | nome           | João Silva               |
      | email          | joao.silva@email.com     |
      | cpf            | 123.456.789-00           |
      | telefone       | (11) 99999-9999          |
      | data_nascimento| 15/03/1990               |
      | cnh            | 12345678901              |
      | status         | ativo                    |
      | avaliacao      | 4.8                      |
    E vejo minha foto de perfil
    E vejo estatísticas sobre minhas corridas

  Esquema do Cenário: Editar dados pessoais básicos com sucesso
    Quando atualizo o dado "<dado>" com o valor "<exemplo>"
    E tento salvar as alterações
    Então vejo a mensagem "Perfil atualizado com sucesso"
    E meus dados são atualizados no sistema

    Exemplos:
      | dado     | exemplo               |
      | nome     | João Santos Silva     |
      | telefone | (11) 88888-8888       |

  Esquema do Cenário: Validação de campos editáveis
    Quando preencho o campo "<campo>" com "<valor_invalido>"
    E tento salvar as alterações
    Então vejo a mensagem de erro "<mensagem_erro>"

    Exemplos:
      | campo    | valor_invalido   | mensagem_erro                |
      | nome     |                  | Nome é obrigatório           |
      | telefone | 999999999        | Formato de telefone inválido |
      | telefone | (11) 9999-999    | Telefone incompleto          |

  Cenário: Alterar senha com sucesso
    Quando preencho o dado "<campo>" com o valor "valor":
      | senha_atual   | MinhaSenh@123  |
      | nova_senha    | NovaSenha@456  |
      | confirmacao   | NovaSenha@456  |
    E salvo a alteração
    Então vejo a mensagem "Senha alterada com sucesso"
    E posso fazer login com a nova senha "NovaSenha@456"

  Esquema do Cenário: Validação na alteração de senha
    Dado que estou na página de alteração de senha
    Quando preencho os campos de senha:
      | senha_atual   | <senha_atual>  |
      | nova_senha    | <nova_senha>   |
      | confirmacao   | <confirmacao>  |
    E tento salvar a alteração
    Então vejo a mensagem de erro "<mensagem_erro>"

    Exemplos:
      | senha_atual    | nova_senha     | confirmacao    | mensagem_erro                                |
      | senha_errada   | NovaSenha@456  | NovaSenha@456  | Senha atual incorreta                        |
      | MinhaSenh@123  | 123456         | 123456         | Nova senha deve atender aos critérios mínimos|
      | MinhaSenh@123  | NovaSenha@456  | NovaSenha@789  | Confirmação de senha não confere             |

  Cenário: Upload de nova foto de perfil válida
    Quando faço upload de "uma nova foto" em formato "JPG" com "1MB" de tamanho
    E tento salvar as alterações
    Então vejo a mensagem "Foto de perfil atualizada com sucesso"

  Esquema do Cenário: Upload de foto inválida
    Dado que estou na página de edição do perfil
    Quando faço upload de "uma nova foto" em formato "<formato>" com "<tamanho>" de tamanho
    Então vejo a mensagem de erro "<mensagem_erro>"

    Exemplos:
      | formato | tamanho         | mensagem_erro                                    |
      | JPG     | 8MB             | Foto muito grande. Tamanho máximo: 5MB           |
      | TXT     | 1MB             | Formato não suportado. Use JPG, PNG ou WEBP      |


  Cenário: Solicitar exclusão permanente da conta
    Quando solicito excluir minha conta permanentemente
    Então recebo um email com instruções para confirmar a exclusão
    E meu status é "aguardando_exclusao"

  Cenário: Cancelar solicitação de exclusão
    Dado que meu status é "aguardando_exclusao"
    Quando eu estou na página "Cancelar Solicitação de Exclusão"
    Então meu status é "ativo"
    E eu vejo a mensagem "Solicitação de exclusão cancelada"

  Cenário: Aceitar solicitação de exclusão
    Dado que meu status é "aguardando_exclusao"
    Quando eu estou na página "Aceitar Solicitação de Exclusão"
    Então meu status é "encerrado"
    E eu vejo a mensagem "Sua conta foi excluída permanentemente"
