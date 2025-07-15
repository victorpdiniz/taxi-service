Feature: Cancelamento de Corrida
    Como um sistema de taxi
    Eu quero gerenciar o cancelamento de corridas
    Para garantir o funcionamento adequado do serviço

Scenario: Motorista cancela corrida antes de chegar ao local
    
    Given o motorista de nome "João" aceitou uma corrida
    And o motorista está a "5" minutos do local de embarque "abc"
    When o motorista seleciona a opção "Cancelar"
    And seleciona "Confirmar" na interface de desencentivo a cancelamento
    Then o sistema cancela o progresso da corrida
    And o motorista recebe a mensagem de confirmação na tela "Corrida cancelada com sucesso"

Scenario: Motorista conclui corrida antes do tempo estimado
Given o tempo estimado de chegada ao destino é de "20" minutos
And o tempo decorrido é de "15" minutos
When o motorista chega ao destino
Then o status da corrida é "concluída com antecedência"
And uma mensagem de parabéns é exibida ao motorista
And um bônus é aplicado ao motorista

Scenario: Motorista conclui corrida dentro do tempo estimado
Given o tempo estimado de chegada ao destino é de "20" minutos
And o tempo decorrido é de "19" minutos
When o motorista chega ao destino
Then o status da corrida é "concluída no tempo previsto"
And nenhuma penalização ou bônus é aplicado

Scenario: Corrida cancelada automaticamente por excesso de tempo
Given o tempo estimado de chegada é de "20" minutos
And o tempo decorrido é de "36" minutos
When o sistema detecta que o tempo ultrapassou o limite de "15" minutos extras
Then a corrida é automaticamente cancelada
And o motorista recebe uma notificação por pop-up com a mensagem "Corrida cancelada por excesso de tempo"