Feature: Notificação de novas corridas por perto

  Scenario: Receber notificação de corrida próxima quando estiver disponível
    Given que estou com o status "disponível"
    And estou logado como "Motorista"
    When uma nova corrida é solicitada por um passageiro próximo
    Then devo receber uma notificação com os detalhes da corrida
    And devo ter a opção de aceitar ou recusar a corrida

  Scenario: Notificação expira após tempo limite
    Given que estou com o status "disponível"
    And recebi uma notificação de corrida
    When não interajo com a notificação em até 20 segundos
    Then a notificação deve desaparecer

  Scenario: Motorista recusa a corrida e continua disponível para novas notificações
    Given que estou logado como "Motorista"
    And meu status está como "disponível"
    And recebi uma notificação de uma nova corrida próxima
    When eu recuso a corrida
    Then meu status deve permanecer como "disponível"
    And não deve ser registrada nenhuma corrida no meu histórico
    And devo continuar recebendo notificações de novas corridas próximas

  Scenario: Corrida aceita deve ser cadastrada no histórico e impedir novas notificações
    Given que estou logado como "Motorista"
    And meu status está como "disponível"
    And recebi uma notificação de uma nova corrida próxima
    When eu aceito a corrida
    Then meu status deve ser alterado para "ocupado"
    And a corrida deve ser registrada no meu histórico como "aceita"
    And enquanto meu status estiver como "ocupado", não devo receber notificações de novas corridas

  Scenario: Motorista desativa notificações de novas corridas
    Given que estou logado como "Motorista"
    And estou na tela de configurações do meu perfil
    When eu desativo a opção "Receber notificações de novas corridas"
    Then o sistema deve registrar que as notificações estão ativadas
    And eu não devo receber notificações de novas corridas enquanto a opção estiver desativada

 Scenario: Motorista ativa notificações de novas corridas
    Given que estou logado como "Motorista"
    And estou na tela de configurações do meu perfil
    When eu ativo novamente a opção "Receber notificações de novas corridas"
    Then o sistema deve registrar que as notificações estão ativadas
    And eu devo voltar a receber notificações de novas corridas enquanto estiver com status "disponível"
