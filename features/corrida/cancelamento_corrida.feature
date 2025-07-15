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


Scenario: Sistema cancela corrida por demora do motorista

    Given o motorista de nome "João" aceitou uma corrida
    And o motorista está a "3" minutos do local de embarque "abc"
    When motorista demora mais do que "15" minutos para chegar ao local
    Then o sistema cancela automaticamente a corrida
    And o motorista recebe a mensagem de erro na tela "Corrida cancelada, limite de tempo atingido" 


Scenario: Tentativa cancelamento de corrida em andamento
    Given o motorista de nome "João" aceitou uma corrida
    And foi ao local de embarque "123" e iniciou o trajeto para o destino "abc"
    When o motorista tenta cancelar a corrida selecionando a opção "Cancelar"
    Then o sistema não permite o cancelamento
    And uma mensagem de erro aparece na tela "Não foi possível cancelar, corrida em andamento"

Scenario: Tentativa cancelamento após chegar ao destino
    Given o motorista de nome "João" aceitou uma corrida
    And foi ao local de embarque "123" e chegou ao destino "abc"
    When o motorista tenta cancelar a corrida selecionando a opção "Cancelar"
    Then o sistema não permite o cancelamento
    And uma mensagem de erro aparece na tela "Destino alcançado"
