// cypress/e2e/corrida_no_tempo.cy.js

describe('Cenário: Corrida No Tempo Previsto', () => {
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

  it('deve criar uma corrida, concluir no tempo e verificar o status', function () {
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

    cy.task('runDriverSimulator', { rideId: RIDE_ID, scenario: 'ON_TIME' }).then(() => {
      cy.get('.status', { timeout: 10000 }).should('contain', 'Motorista Encontrado');

      cy.log('Aguardando a conclusão da corrida no tempo previsto...');
      // O cenário ON_TIME do simulador vai demorar exatamente o tempo estimado (1 minuto)
      // Esperamos um pouco mais que 1 minuto para a conclusão.
      cy.wait(ESTIMATED_TIME_MS + 5000); // Espera 1 minuto e 5 segundos

      cy.get('.status', { timeout: 35000 }).should('contain', 'Concluída No Tempo Previsto');

      cy.log('Teste concluído! O status foi atualizado para Concluída No Tempo Previsto como esperado.');
    });
  });
});