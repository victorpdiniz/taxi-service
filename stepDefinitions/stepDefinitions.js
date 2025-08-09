Given("o usuário está na página de login", () => {
    cy.visit("/login");
});

When("o usuário digita seu nome de usuário e senha", () => {
    cy.get("#username").type("admin");
    cy.get("#password").type("secret");
});

And("clica no botão \"Login\"", () => {
    cy.get("#loginButton").click();
});

Then("o usuário é direcionado para a página principal", () => {
    cy.url().should("include", "/");
});