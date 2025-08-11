// cypress/e2e/corrida_antecipada.cy.js

describe('Cenário: Corrida Antecipada', () => {
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

  it('deve criar uma corrida, concluir antes do tempo e verificar o status', function () {
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

    cy.task('runDriverSimulator', { rideId: RIDE_ID, scenario: 'EARLY' }).then(() => {
      cy.get('.status', { timeout: 10000 }).should('contain', 'Motorista Encontrado');

      cy.log('Aguardando a conclusão antecipada da corrida...');
      // O cenário EARLY do simulador vai demorar 30 segundos a menos que o estimado (1 minuto - 30s = 30s)
      // Esperamos um pouco mais que 30 segundos para a conclusão.
      cy.wait(ESTIMATED_TIME_MS - 25000); // Espera 35 segundos

      cy.get('.status', { timeout: 35000 }).should('contain', 'Concluída Com Antecedência');

      cy.log('Teste concluído! O status foi atualizado para Concluída Com Antecedência como esperado.');
    });
  });
});