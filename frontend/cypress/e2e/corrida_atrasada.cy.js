// cypress/e2e/corrida_atrasada.cy.js

describe('Cenário: Corrida Atrasada', () => {
  // Usamos o hook 'before' para garantir que a corrida seja criada uma vez
  // antes de todos os testes neste bloco. Isso torna o teste independente e robusto.
  before(function () {
    // Cria uma nova corrida antes do teste começar
    cy.request({
      method: 'POST',
      url: 'http://localhost:3000/api/corrida',
      body: {
        PassageiroID: 123, // ID de um passageiro de teste
        Origem: '-23.550520,-46.633308', // Ex: Centro de São Paulo
        Destino: '-23.561353,-46.656489', // Ex: Av. Paulista
      },
    }).then((response) => {
      // Verificamos se a resposta tem o corpo esperado e o ID da corrida
      expect(response.status).to.eq(201);
      expect(response.body).to.have.property('ID');
      
      // Armazenamos os detalhes da corrida para usar em outros testes
      // cy.wrap(response.body).as('ride');
      // Cypress não compartilha o contexto do 'as' entre 'before' e 'it' da mesma forma
      // que compartilha entre 'beforeEach' e 'it'.
      // Portanto, usamos uma variável 'this' do Mocha.
      this.ride = response.body;
    });
  });

  // O tempo estimado é de 1 minuto (60000 ms), conforme configurado no backend.
  const ESTIMATED_TIME_MS = 60000;

  it('deve criar uma corrida, aguardar o atraso e verificar o status', function () {
    // Verificamos se a corrida foi criada no hook 'before'
    expect(this.ride).to.not.be.undefined;
    const RIDE_ID = this.ride.ID;

    // Intercepta a chamada GET para a API da corrida para garantir dados frescos.
    cy.intercept({
      method: 'GET',
      url: `http://localhost:3000/api/corrida/${RIDE_ID}`,
    }, (req) => {
      req.headers['cache-control'] = 'no-cache';
    }).as('getRideStatus');

    // Visita a página da corrida específica
    cy.visit(`/corrida/${RIDE_ID}`);

    // Aumentamos o timeout para dar tempo ao simulador de aceitar a corrida.
    // Esperamos pela primeira chamada de API ser completada.
    cy.wait('@getRideStatus');

    // Executa o simulador de motorista via Cypress task
    cy.task('runDriverSimulator', { rideId: RIDE_ID, scenario: 'LATE' }).then(() => {
      // Agora que o simulador aceitou a corrida, o status deve ser "Motorista Encontrado"
      cy.get('.status', { timeout: 10000 }).should('contain', 'Motorista Encontrado');

      // O cenário LATE do simulador vai demorar 1m30s.
      // Vamos esperar um pouco mais que 1 minuto para que o backend marque a corrida como atrasada.
      cy.log('Aguardando o tempo estimado da corrida ser ultrapassado...');
      cy.wait(ESTIMATED_TIME_MS + 5000); // Espera 1 minuto e 5 segundos

      // Após a espera, o status na tela DEVE ter sido atualizado para "Atrasado".
      // O monitor do backend roda a cada 30s, então a atualização pode não ser instantânea.
      cy.get('.status', { timeout: 35000 }).should('contain', 'Atrasado');

      cy.log('Teste concluído! O status foi atualizado para Atrasado como esperado.');
    });
  });
});