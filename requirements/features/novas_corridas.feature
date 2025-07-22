Feature: Notificação de novas corridas por perto
  Como um motorista
  Eu quero receber notificações de novas corridas próximas
  Para que eu possa aceitar corridas 

  Background:
    Given que o motorista de nome “João” está logado
    And sua localização está em "Rua A, 123"
    And está com o status "disponível"

  Scenario: Receber notificação de corrida próxima quando estiver disponível
    Given uma corrida é solicitada dentro de um raio de "2" quilômetros
    When uma nova corrida é solicitada por um passageiro próximo
    Then deve receber uma notificação contendo:
      | tempo       | 23 minutos           |
      | distância   | 1.5 km               |
      | valor       | R$ 12,50             |
      | passageiro  | Maria Silva          |
    And deve ter botões de "Aceitar" e "Recusar" 

  Scenario: Notificação expira após tempo limite
    Given recebe uma notificação de corrida
    When passa 20 segundos sem interagir com a notificação
    Then a notificação deve desaparecer automaticamente
    And o status deve permanecer como "disponível"

  Scenario: Corrida aceita é registrada no histórico
    Given recebe uma notificação de corrida
    When clica em "Aceitar"
    Then o status deve ser alterado para "ocupado"
    And a corrida deve ser registrada no histórico como "em_andamento"

  Scenario: Motorista ocupado não recebe notificações
    Given o status está "ocupado"
    When uma nova corrida é solicitada próxima
    Then não deve receber notificação

  Scenario: Recusar corrida mantém disponibilidade
    Given recebe uma notificação de corrida
    When clica em "Recusar"
    Then o status deve permanecer "disponível"

  Scenario: Desativar notificações altera status
    Given está recebendo notificações
    When desativa a opção "Receber notificações de novas corridas"
    Then o status deve ser alterado para "ocupado"

  Scenario: Ativar notificações altera status
    Given as notificações estão desativadas
    And o status está "ocupado"
    When ativa a opção "Receber notificações de novas corridas"
    Then o status deve ser alterado para "disponível"
