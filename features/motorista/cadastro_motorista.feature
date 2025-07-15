# language: pt
Funcionalidade: Cadastro de Motorista
  Como um futuro motorista
  Eu quero me cadastrar no sistema
  Para poder aceitar corridas e trabalhar como motorist

  Cenário: Cadastro bem-sucedido com todos os dados válidos
    Dado eu estou na página "Cadastro de Motorista"
    Quando eu preencho o formulário com dados válidos:
      | nome           | João Silva                   |
      | data_nascimento| 15/03/1990                   |
      | cpf            | 221.623.340-46               |
      | cnh            | 12345678901                  |
      | categoria_cnh  | B                            |
      | validade_cnh   | 15/03/2030                   |
      | placa_veiculo  | ABC1234                      |
      | modelo_veiculo | Honda Civic 2020             |
      | telefone       | (11) 99999-9999              |
      | email          | joao.silva@email.com         |
      | senha          | MinhaSenh@123                |
    E eu submeto o formulário
    Então vejo a mensagem "Cadastro realizado com sucesso"
    E recebo um email de confirmação
    E meu status é "aguardando_aprovacao"

  Esquema do Cenário: Validação de campos obrigatórios
    Dado eu estou na página "Cadastro de Motorista"
    Quando preencho o campo "<campo_faltante>" com o valor ""
    E eu submeto o formulário
    Então vejo a mensagem de erro "<mensagem_erro>"
    E não consigo concluir o cadastro

    Exemplos:
      | campo_faltante  | mensagem_erro                       |
      | nome            | Nome é obrigatório                  |
      | data_nascimento | CPF é obrigatório                   |
      | cpf             | CNH é obrigatória                   |
      | cnh             | Email é obrigatório                 |
      | categoria_cnh   | Senha é obrigatória                 |
      | validade_cnh    | Telefone é obrigatório              |
      | placa_veiculo   | Placa do veículo é obrigatória      |
      | modelo_veiculo  | Modelo do Veículo é obrigatório     |
      | telefone        | Telefone é obrigatório              |
      | email           | E-mail é obrigatório                |
      | senha           | Senha é obrigatória                 |

  Esquema do Cenário: Validação de formatos de documentos
    Dado eu estou na página "Cadastro de Motorista"
    Quando preencho o campo "<campo>" com o valor "<valor_invalido>"
    Então vejo a mensagem de erro "<mensagem_erro>"
    E não consigo concluir o cadastro

    Exemplos:
      | campo            | valor_invalido   | mensagem_erro                                   |
      | cpf              | 123.456.789-99   | CPF inválido                                    |
      | cpf              | 111.111.111-11   | CPF inválido                                    |
      | cnh              | 1234567890       | CNH deve ter 11 dígitos                         |
      | email            | email_invalido   | Formato de email inválido                       |
      | telefone         | 999999999        | Formato de telefone inválido                    |
      | placa_veiculo    | ABC12345         | Formato de placa inválido                       |
      | data_nascimento  | 01/01/2099       | Motorista deve ter pelo menos 18 anos           |
      | validade_cnh     | 01/01/1970       | CNH vencida. Renove sua CNH para prosseguir     |

  Esquema do Cenário: Tentativa de cadastro com dados já existentes
    Dado eu estou na página "Cadastro de Motorista"
    E que existe um motorista cadastrado com "<campo>" "<valor_existente>"
    Quando preencho o campo "<campo>" com o valor "<valor_existente>"
    E eu submeto o formulário
    Então vejo a mensagem de erro "<mensagem_erro>"
    E não consigo concluir o cadastro

    Exemplos:
      | campo | valor_existente      | mensagem_erro              |
      | cpf   | 123.456.789-00       | CPF já cadastrado          |
      | cnh   | 12345678901          | CNH já cadastrada          |
      | email | joao.silva@email.com | Email já cadastrado        |

  Cenário: Upload de documentos obrigatórios
    Dado eu estou na página "Upload de Documentos"
    E completei o cadastro básico com sucesso
    Quando faço upload dos documentos obrigatórios:
      | documento        | formato | tamanho |
      | CNH              | JPG     | 2MB     |
      | CRLV             | PNG     | 1.5MB   |
      | selfie com CNH   | JPG     | 1MB     |
    Então vejo a mensagem "Documentos enviados com sucesso"
    E meu status é "documentos_em_analise"
    E eu recebo um email confirmando o recebimento dos documentos

  Esquema do Cenário: Rejeição de upload de documento inválido
    Dado eu estou na página "Upload de Documentos"
    E completei o cadastro básico com sucesso
    Quando tento fazer upload de arquivo com "<problema>"
    Então vejo a mensagem de erro "<mensagem_erro>"
    E o upload não é concluído

    Exemplos:
      | problema           | mensagem_erro                                 |
      | tamanho de 6MB     | Arquivo muito grande. Tamanho máximo: 5MB     |
      | formato TXT        | Formato não suportado. Use JPG, PNG ou PDF    |

  Cenário: Senha fraca
    Dado eu estou na página "Cadastro de Motorista"
    Quando eu preencho a senha "123456"
    Então eu vejo o indicador de força da senha como "Fraca"
    E eu vejo a mensagem de erro "Senha deve ter pelo menos 8 caracteres, incluindo maiúscula, minúscula, número e símbolo"
    E não consigo concluir o cadastro

  Cenário: Senha suficientemente forte
    Dado eu estou na página "Cadastro de Motorista"
    Quando eu preencho a senha "Senh@123"
    Então eu vejo o indicador de força da senha como "Média"
    E consigo concluir o cadastro

  Cenário: Confirmação de senha diferente
    Dado que eu estou na página de cadastro de motorista
    Quando eu preencho a senha "MinhaSenh@123"
    E eu preencho a confirmação de senha "MinhaSenh@456"
    E eu submeto o formulário
    Então eu vejo a mensagem de erro "Senhas não conferem"

  Cenário: Verificação de documentos rejeitada
    Dado que eu enviei documentos com problemas de qualidade
    Quando o sistema ou analista verifica meus documentos
    Então meu status é "rejeitado"
    E eu recebo um email detalhando os problemas encontrados 
    # TODO: Ser específico

  Cenário: Verificação de documentos rejeitada
    Dado que eu enviei documentos válidos
    Quando o sistema ou analista verifica meus documentos
    Então meu status é "ativo"

