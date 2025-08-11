// cypress/e2e/corrida_cancelada_auto.cy.js

describe('Cenário: Corrida Cancelada Automaticamente por Excesso de Tempo', () => {
  before(function () {
    cy.request({
      method: 'POST',
      url: 'http://localhost:3000/api/corrida',
      body: {
        PassageiroID: 123,
        Origem: '-23.550520,-46.633308',
        Destino: '-23.561353,-46.656489',
      },
    }).then((response) => {
      expect(response.status).to.eq(201);
      expect(response.body).to.have.property('ID');
      this.ride = response.body;
    });
  });

  const ESTIMATED_TIME_MS = 60000; // 1 minuto
  const AUTO_CANCEL_THRESHOLD_MS = 15 * 60 * 1000; // 15 minutos em milissegundos

  it('deve criar uma corrida e verificar o cancelamento automático por excesso de tempo', function () {
    expect(this.ride).to.not.be.undefined;
    const RIDE_ID = this.ride.ID;

    cy.intercept({
      method: 'GET',
      url: `http://localhost:3000/api/corrida/${RIDE_ID}`,
    }, (req) => {
      req.headers['cache-control'] = 'no-cache';
    }).as('getRideStatus');

    cy.visit(`/corrida/${RIDE_ID}`);
    cy.wait('@getRideStatus');

    cy.task('runDriverSimulator', { rideId: RIDE_ID, scenario: 'AUTO_CANCEL' }).then(() => {
      cy.get('.status', { timeout: 10000 }).should('contain', 'Motorista Encontrado');

      cy.log('Aguardando o cancelamento automático da corrida por excesso de tempo...');
      // O cenário AUTO_CANCEL do simulador não finaliza a corrida, esperando o backend cancelar.
      // O backend cancela após o tempo estimado + 15 minutos.
      cy.wait(ESTIMATED_TIME_MS + AUTO_CANCEL_THRESHOLD_MS + 5000); // Espera o tempo de cancelamento + 5 segundos

      cy.get('.status', { timeout: 35000 }).should('contain', 'Cancelada Por Excesso De Tempo');

      cy.log('Teste concluído! O status foi atualizado para Cancelada Por Excesso De Tempo como esperado.');
    });
  });
});