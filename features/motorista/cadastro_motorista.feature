# language: pt
Funcionalidade: Cadastro de Motorista
  Como um futuro motorista
  Eu quero me cadastrar no sistema
  Para poder aceitar corridas e trabalhar como Motorista

  Cenário: Cadastro bem-sucedido
    Dado estou na página "Cadastro de Motorista"
    E submeto o cadastro com os dados:
      | campo           | valor                |
      | nome            | João Silva           |
      | data_nascimento | 15/03/1990           |
      | cpf             | 221.623.340-46       |
      | cnh             | 12345678901          |
      | categoria_cnh   | B                    |
      | validade_cnh    | 15/03/2030           |
      | placa_veiculo   | ABC1234              |
      | modelo_veiculo  | Honda Civic 2020     |
      | telefone        | (11) 99999-9999      |
      | email           | joao.silva@email.com |
      | senha           | MinhaSenh@123        |
    E não existe um motorista cadastrado com os dados:
      | campo         | valor                |
      | email         | joao.silva@email.com |
      | cpf           | 221.623.340-46       |
      | placa_veiculo | ABC1234              |
    Então vejo a mensagem "Cadastro realizado com sucesso"
    E o status de "joao.silva@email.com" é "aguardando_documentos"
    E estou autenticado como "joao.silva@email.com"
    E estou na página "Upload de Documentos"

  Esquema do Cenário: Tentativa de cadastro com informações já cadastradas
    Dado estou na página "Cadastro de Motorista"
    E submeto o cadastro com os dados:
      | campo           | valor                |
      | nome            | João Silva           |
      | data_nascimento | 15/03/1990           |
      | cpf             | 221.623.340-46       |
      | cnh             | 12345678901          |
      | categoria_cnh   | B                    |
      | validade_cnh    | 15/03/2030           |
      | placa_veiculo   | ABC1234              |
      | modelo_veiculo  | Honda Civic 2020     |
      | telefone        | (11) 99999-9999      |
      | email           | joao.silva@email.com |
      | senha           | MinhaSenh@123        |
    E existe um motorista cadastrado com o dado "<dado>" no valor "<valor>"
    Então vejo a mensagem de erro "<mensagem>"
    E estou na página "Cadastro de Motorista"

    Exemplos:
    | dado           | valor                | mensagem                |
    | telefone       | (11) 99999-9999      | Telefone já cadastrado. |
    | email          | joao.silva@email.com | E-mail já cadastrado.   |
    | cpf            | 221.623.340-46       | CPF já cadastrado.      |
    | cnh            | 12345678901          | CNH já cadastrada.      |
    | placa_veiculo  | ABC1234              | Veículo já cadastrado.  |

  Esquema do Cenário: Validação de campos obrigatórios
    Dado eu estou na página "Cadastro de Motorista"
    Quando preencho o campo "<campo>" com o valor ""
    E submeto o formulário
    Então vejo a mensagem de erro "<mensagem>"
    E estou na página "Cadastro do Motorista"

    Exemplos:
      | campo           | mensagem                             |
      | nome            | Nome é obrigatório.                  |
      | data_nascimento | CPF é obrigatório.                   |
      | cpf             | CNH é obrigatória.                   |
      | cnh             | E-mail é obrigatório.                |
      | categoria_cnh   | Senha é obrigatória.                 |
      | validade_cnh    | Telefone é obrigatório.              |
      | placa_veiculo   | Placa do veículo é obrigatória.      |
      | modelo_veiculo  | Modelo do Veículo é obrigatório.     |
      | telefone        | Telefone é obrigatório.              |
      | email           | E-mail é obrigatório.                |
      | senha           | Senha é obrigatória.                 |

  Esquema do Cenário: Validação dos campos de cadastro
    Dado eu estou na página "Cadastro de Motorista"
    Quando preencho o campo "<campo>" com o valor "<valor_invalido>"
    E submeto o formulário
    Então vejo a mensagem de erro "<mensagem>"
    E estou na página "Cadastro de Motorista"

    Exemplos:
      | campo            | valor_invalido  | mensagem                                                                                 |
      | cpf              | 123.456.789-99  | CPF inválido.                                                                            |
      | cpf              | 111.111.111-11  | CPF inválido.                                                                            |
      | cpf              | 22162334046     | Formato de CPF inválido.                                                                 |
      | cnh              | 1234567890      | CNH inválida.                                                                            |
      | validade_cnh     | 01/01/1970      | CNH vencida. Renove sua CNH para prosseguir                                              |
      | email            | email_invalido  | Formato de email inválido.                                                               |
      | telefone         | 999999999       | Formato de telefone inválido                                                             |
      | telefone         | (11) 99999-9999 | Formato de telefone inválido                                                             |
      | placa_veiculo    | ABC12345        | Placa inválida                                                                           |
      | data_nascimento  | 01/01/2099      | Você deve ter 18 anos ou mais para dirigir.                                              |
      | senha            | 123456          | Senha deve ter pelo menos 8 caracteres, incluindo maiúscula, minúscula, número e símbolo |

  Cenário: Upload de documentos com sucesso
    Dado eu estou na página "Upload de Documentos"
    E o status de "joao.silva@email.com" é "aguardando_documentos"
    E estou autenticado como "joao.silva@email.com"
    Quando faço upload dos documentos obrigatórios:
      | documento        | formato | tamanho |
      | CNH              | PDF     | 2MB     |
      | CRLV             | PNG     | 1.5MB   |
      | selfie com CNH   | JPG     | 1MB     |
    Então vejo a mensagem "Documentos enviados com sucesso"
    E o status de "joao.silva@email.com" é "documentos_em_analise"
    E um email é enviado para "joao.silva@email.com" com o assunto "Documentos recebidos"

  Esquema do Cenário: Tentativa de upload de documento inválido
    Dado eu estou na página "Upload de Documentos"
    E o status de "joao.silva@email.com" é "aguardando_documentos"
    E estou autenticado como "joao.silva@email.com"
    Quando tento fazer upload de arquivo com tamanho "<tamanho>" e formato "<formato>"
    Então vejo a mensagem de erro "<mensagem>"
    E o status de "joao.silva@email.com" é "aguardando_documentos"

    Exemplos:
      | tamanho | formato | mensagem                                   |
      | 10MB    | PNG     | Arquivo muito grande. Tamanho máximo: 5MB  |
      | 1MB     | TXT     | Formato não suportado. Use JPG, PNG ou PDF |

  Cenário: Verificação de documentos rejeitada
    Dado "joao.silva@email.com" enviou documentos com problemas de qualidade
    Quando o sistema ou analista verifica os documentos de "joao.silva@email.com"
    Então o status de "joao.silva@email.com" é "documentos_rejeitados"
    E um email é enviado para "joao.silva@email.com" com o assunto "Documentos rejeitados" 

  Cenário: Verificação de documentos aceita
    Dado "joao.silva@email.com" enviou documentos válidos
    Quando o sistema ou analista verifica os documentos de "joao.silva@email.com"
    Então o status de "joao.silva@email.com" é "ativo"
    E um email é enviado para "joao.silva@email.com" com o assunto "Parabéns! Seu cadastro foi aprovado" 


