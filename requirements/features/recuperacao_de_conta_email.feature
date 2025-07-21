Scenario: Usuário vai para a página de recuperação de conta
Given estou na página “Login”
And seleciono a opção “Esqueceu sua conta?”
Then estou na página “Recuperação de Conta”
And há um formulário de título “qual o email usado na sua conta?”

Scenario: Usuário tenta recuperar conta com um email que existe
Given estou na página “Recuperação de Conta”
And o email “jose@bol.com.br” está cadastrado na plataforma
When submeto o formulário “qual o email usado na sua conta?” com o email “jose@bol.com.br”
Then um email é disparado para “jose@bol.com.br” contendo um link para a página de recuperação de conta

Scenario: Usuário tenta recuperar conta com um email que não existe
Given estou na página “Recuperação de Conta”
And o email “jose@bol.com.br” não está cadastrado na plataforma
When submeto o formulário “qual o email usado na sua conta?” com o email “jose@bol.com.br”
Then o texto de erro do formulário “qual o email usado na sua conta?” diz “Esse email não está registrado”
And nenhum email é enviado

Scenario: Usuário clica em um link expirado de recuperação de conta
Given o usuário “José” solicitou recuperação de conta há mais de 24h
When estou na página de recuperação de conta para o usuário “José”
Then a plataforma exibe a mensagem “Esse link de recuperação está expirado. Por favor, solicite a recuperação de conta novamente”