# language: pt
Funcionalidade: Autenticação de Motorista
  Como um motorista cadastrado
  Eu quero fazer login no sistema
  Para acessar minhas funcionalidades de trabalho

  Contexto:
    Dado existe um motorista cadastrado com os dados:
      | dado  | valor                |
      | email | joao.silva@email.com |
      | senha | MinhaSenh@123        |

  Cenário: Login bem-sucedido de motorista
    Dado estou na página "Login do Motorista"
    E não estou autenticado como motorista
    E o status de "joao.silva@email.com" é "ativo"
    Quando realizo login com email "joao.silva@email.com" e senha "MinhaSenh@123"
    Então estou autenticado como "joao.silva@email.com"
    E estou na página "Painel do Motorista"

  Esquema do Cenário: Tentativa de login com credenciais inválidas
    Dado estou na página "Login do Motorista"
    E não estou autenticado como motorista
    E o status de "joao.silva@email.com" é "ativo"
    Quando realizo login com email "<email>" e senha "<senha>"
    Então vejo a mensagem de erro "<mensagem>"
    E estou na página "Login do Motorista"

    Exemplos:
      | email                    | senha         | mensagem                   |
      | joao.silva@email.com     | senhaerrada   | Credenciais inválidas      |
      | inexistente@email.com    | qualquersenha | Credenciais inválidas      |
      | email_invalido           | MinhaSenh@123 | Formato de email inválido  |
      | joao.silva@email.com     |               | Senha é obrigatória        |
      |                          | MinhaSenh@123 | Email é obrigatório        |

  Esquema do Cenário: Login de motorista sem documentos aprovados
    Dado estou na página "Login do Motorista"
    E não estou autenticado como motorista
    Dado o status de "joao.silva@email.com" é "<status>":
    Quando realizo login com email "joao.silva@email.com" e senha "MinhaSenh@123"
    Então estou na página "Upload de Documentos"

    Exemplos:
      | status                |
      | aguardando_documentos | 
      | documentos_rejeitados |

  Esquema do Cenário: Tentativa de login com status inválido
    Dado estou na página "Login do Motorista"
    E não estou autenticado como motorista
    E o status de "joao.silva@email.com" é "<status>"
    Quando realizo login com email "joao.silva@email.com" e senha "MinhaSenh@123"
    Então vejo a mensagem de erro "<mensagem>"
    E estou na página "Login do Motorista"

    Exemplos:
      | status                | mensagem                                                                                                |
      | documentos_em_analise | Seus documentos ainda estão em análise.                                        | 
      | aguardando_exclusao   | Sua conta está em processo de exclusão. Contate o suporte se isso for um erro. |
      | encerrado             | Sua conta foi encerrada. Em caso de dúvidas, contate o suporte.                |

  Esquema do Cenário: Redirecionamento de motorista autenticado
    Dado estou autenticado como "joao.silva@email.com"
    Quando estou na página "<pagina>"
    Então estou na página "Painel do Motorista"

    Exemplos:
    | pagina                |
    | Cadastro do Motorista |
    | Login do Motorista    |
    | Upload de Documentos  |

  
  Esquema do Cenário: Tentativa de acesso sem estar autenticado
    Dado não estou autenticado como motorista
    Quando estou na página "<pagina>"
    Então estou na página "Login do Motorista"

    Exemplos:
    | pagina              |
    | Painel do Motorista |
 # TODO: Inserir todas as outras páginas aqui