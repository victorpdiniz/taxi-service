package test

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestMotorista(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenarioMotorista,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../features/motorista/"},
			Output:   colors.Colored(os.Stdout),
			TestingT: t,
			Strict:   true,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func acesseiOLinkDeRecuperaoVlido() error {
	return godog.ErrPending
}


func acessoOLinkDeRecuperaoExpirado() error {
	return godog.ErrPending
}


func atualizoODadoComOValor(arg1, arg2 string) error {
	return godog.ErrPending
}


func completeiOCadastroBsicoComSucesso() error {
	return godog.ErrPending
}


func confirmoANovaSenha(arg1 string) error {
	return godog.ErrPending
}


func consigoConcluirOCadastro() error {
	return godog.ErrPending
}


func estouNaPgina(arg1 string) error {
	return godog.ErrPending
}


func euEstouNaPgina(arg1 string) error {
	return godog.ErrPending
}


func euPreenchoAConfirmaoDeSenha(arg1 string) error {
	return godog.ErrPending
}


func euPreenchoASenha(arg1 string) error {
	return godog.ErrPending
}


func euPreenchoOFormulrioComDadosVlidos(arg1 *godog.Table) error {
	return godog.ErrPending
}


func euReceboUmEmailConfirmandoORecebimentoDosDocumentos() error {
	return godog.ErrPending
}


func euReceboUmEmailDetalhandoOsProblemasEncontrados() error {
	return godog.ErrPending
}


func euSubmetoOFormulrio() error {
	return godog.ErrPending
}


func euVejoAMensagem(arg1 string) error {
	return godog.ErrPending
}


func euVejoAMensagemDeErro(arg1 string) error {
	return godog.ErrPending
}


func euVejoOIndicadorDeForaDaSenhaComo(arg1 string) error {
	return godog.ErrPending
}


func faoLoginComE(arg1, arg2 string) error {
	return godog.ErrPending
}


func faoUploadDeEmFormatoComDeTamanho(arg1, arg2, arg3 string) error {
	return godog.ErrPending
}


func faoUploadDosDocumentosObrigatrios(arg1 *godog.Table) error {
	return godog.ErrPending
}


func meuStatus(arg1 string) error {
	return godog.ErrPending
}


func meusDadosSoAtualizadosNoSistema() error {
	return godog.ErrPending
}


func nenhumEmailEnviado() error {
	return godog.ErrPending
}


func noConsigoConcluirOCadastro() error {
	return godog.ErrPending
}


func oLinkDeRecuperaoInvalidado() error {
	return godog.ErrPending
}


func oLinkExpiraEmHora(arg1 int) error {
	return godog.ErrPending
}


func oSistemaOuAnalistaVerificaMeusDocumentos() error {
	return godog.ErrPending
}


func oUploadNoConcludo() error {
	return godog.ErrPending
}


func possoFazerLoginComANovaSenha(arg1 string) error {
	return godog.ErrPending
}


func preenchoANovaSenha(arg1 string) error {
	return godog.ErrPending
}


func preenchoOCampoCom(arg1, arg2 string) error {
	return godog.ErrPending
}


func preenchoOCampoComOValor(arg1, arg2 string) error {
	return godog.ErrPending
}


func preenchoODadoComOValor(arg1, arg2 string, arg3 *godog.Table) error {
	return godog.ErrPending
}


func preenchoOsCamposDeSenha(arg1 *godog.Table) error {
	return godog.ErrPending
}


func queEstouAutenticadoComoMotorista(arg1 *godog.Table) error {
	return godog.ErrPending
}


func queEstouNaPginaAtravsDeLinkVlido(arg1 string) error {
	return godog.ErrPending
}


func queEstouNaPginaDeAlteraoDeSenha() error {
	return godog.ErrPending
}


func queEstouNaPginaDeEdioDoPerfil() error {
	return godog.ErrPending
}


func queEuEnvieiDocumentosComProblemasDeQualidade() error {
	return godog.ErrPending
}


func queEuEnvieiDocumentosVlidos() error {
	return godog.ErrPending
}


func queEuEstouNaPginaDeCadastroDeMotorista() error {
	return godog.ErrPending
}


func queExisteUmMotoristaCadastradoCom(arg1 string, arg2 int) error {
	return godog.ErrPending
}


func queExisteUmMotoristaCadastradoComJoaosilvaemailcom(arg1 string) error {
	return godog.ErrPending
}


func queExisteUmMotoristaCadastradoComOsDados(arg1 *godog.Table) error {
	return godog.ErrPending
}


func queMeuStatus(arg1 string) error {
	return godog.ErrPending
}


func queRedefiniMinhaSenhaUsandoUmLinkDeRecuperao() error {
	return godog.ErrPending
}


func queSoliciteiRecuperaoDeContaHHoras(arg1 int) error {
	return godog.ErrPending
}


func queSoliciteiRecuperaoDeContaHMinutos(arg1 int) error {
	return godog.ErrPending
}


func receboUmEmailComInstruesParaConfirmarAExcluso() error {
	return godog.ErrPending
}


func receboUmEmailDeConfirmao() error {
	return godog.ErrPending
}


func salvoAAlterao() error {
	return godog.ErrPending
}


func solicitoExcluirMinhaContaPermanentemente() error {
	return godog.ErrPending
}


func solicitoRecuperaoNovamenteComOMesmoEmail() error {
	return godog.ErrPending
}


func submetoOFormulrio() error {
	return godog.ErrPending
}


func submetoOFormulrioComEmail(arg1 string) error {
	return godog.ErrPending
}


func tentoAcessarOMesmoLinkNovamente() error {
	return godog.ErrPending
}


func tentoFazerUploadDeArquivoCom(arg1 string) error {
	return godog.ErrPending
}


func tentoSalvarAAlterao() error {
	return godog.ErrPending
}


func tentoSalvarAsAlteraes() error {
	return godog.ErrPending
}


func umEmailDeRecuperaoEnviado() error {
	return godog.ErrPending
}


func vejo(arg1 string) error {
	return godog.ErrPending
}


func vejoAMensagem(arg1 string) error {
	return godog.ErrPending
}


func vejoAMensagemDeErro(arg1 string) error {
	return godog.ErrPending
}


func vejoAOpoDeSolicitarUmNovoLink() error {
	return godog.ErrPending
}


func vejoErro(arg1 string) error {
	return godog.ErrPending
}


func vejoEstatsticasSobreMinhasCorridas() error {
	return godog.ErrPending
}


func vejoMeuDadoComOValor(arg1, arg2 string, arg3 *godog.Table) error {
	return godog.ErrPending
}


func vejoMinhaFotoDePerfil() error {
	return godog.ErrPending
}


func vejoQueDevoAguardarAntesDeSolicitarNovamente() error {
	return godog.ErrPending
}


func InitializeScenarioMotorista(ctx *godog.ScenarioContext) {
		ctx.Step(`^acessei o link de recuperação válido$`, acesseiOLinkDeRecuperaoVlido)
		ctx.Step(`^acesso o link de recuperação expirado$`, acessoOLinkDeRecuperaoExpirado)
		ctx.Step(`^atualizo o dado "([^"]*)" com o valor "([^"]*)"$`, atualizoODadoComOValor)
		ctx.Step(`^completei o cadastro básico com sucesso$`, completeiOCadastroBsicoComSucesso)
		ctx.Step(`^confirmo a nova senha "([^"]*)"$`, confirmoANovaSenha)
		ctx.Step(`^consigo concluir o cadastro$`, consigoConcluirOCadastro)
		ctx.Step(`^estou na página "([^"]*)"$`, estouNaPgina)
		ctx.Step(`^eu estou na página "([^"]*)"$`, euEstouNaPgina)
		ctx.Step(`^eu preencho a confirmação de senha "([^"]*)"$`, euPreenchoAConfirmaoDeSenha)
		ctx.Step(`^eu preencho a senha "([^"]*)"$`, euPreenchoASenha)
		ctx.Step(`^eu preencho o formulário com dados válidos:$`, euPreenchoOFormulrioComDadosVlidos)
		ctx.Step(`^eu recebo um email confirmando o recebimento dos documentos$`, euReceboUmEmailConfirmandoORecebimentoDosDocumentos)
		ctx.Step(`^eu recebo um email detalhando os problemas encontrados$`, euReceboUmEmailDetalhandoOsProblemasEncontrados)
		ctx.Step(`^eu submeto o formulário$`, euSubmetoOFormulrio)
		ctx.Step(`^eu vejo a mensagem "([^"]*)"$`, euVejoAMensagem)
		ctx.Step(`^eu vejo a mensagem de erro "([^"]*)"$`, euVejoAMensagemDeErro)
		ctx.Step(`^eu vejo o indicador de força da senha como "([^"]*)"$`, euVejoOIndicadorDeForaDaSenhaComo)
		ctx.Step(`^faço login com "([^"]*)" e "([^"]*)"$`, faoLoginComE)
		ctx.Step(`^faço upload de "([^"]*)" em formato "([^"]*)" com "([^"]*)" de tamanho$`, faoUploadDeEmFormatoComDeTamanho)
		ctx.Step(`^faço upload dos documentos obrigatórios:$`, faoUploadDosDocumentosObrigatrios)
		ctx.Step(`^meu status é "([^"]*)"$`, meuStatus)
		ctx.Step(`^meus dados são atualizados no sistema$`, meusDadosSoAtualizadosNoSistema)
		ctx.Step(`^nenhum email é enviado$`, nenhumEmailEnviado)
		ctx.Step(`^não consigo concluir o cadastro$`, noConsigoConcluirOCadastro)
		ctx.Step(`^o link de recuperação é invalidado$`, oLinkDeRecuperaoInvalidado)
		ctx.Step(`^o link expira em (\d+) hora$`, oLinkExpiraEmHora)
		ctx.Step(`^o sistema ou analista verifica meus documentos$`, oSistemaOuAnalistaVerificaMeusDocumentos)
		ctx.Step(`^o upload não é concluído$`, oUploadNoConcludo)
		ctx.Step(`^posso fazer login com a nova senha "([^"]*)"$`, possoFazerLoginComANovaSenha)
		ctx.Step(`^preencho a nova senha "([^"]*)"$`, preenchoANovaSenha)
		ctx.Step(`^preencho o campo "([^"]*)" com "([^"]*)"$`, preenchoOCampoCom)
		ctx.Step(`^preencho o campo "([^"]*)" com o valor "([^"]*)"$`, preenchoOCampoComOValor)
		ctx.Step(`^preencho o dado "([^"]*)" com o valor "([^"]*)":$`, preenchoODadoComOValor)
		ctx.Step(`^preencho os campos de senha:$`, preenchoOsCamposDeSenha)
		ctx.Step(`^que estou autenticado como motorista:$`, queEstouAutenticadoComoMotorista)
		ctx.Step(`^que estou na página "([^"]*)" através de link válido$`, queEstouNaPginaAtravsDeLinkVlido)
		ctx.Step(`^que estou na página de alteração de senha$`, queEstouNaPginaDeAlteraoDeSenha)
		ctx.Step(`^que estou na página de edição do perfil$`, queEstouNaPginaDeEdioDoPerfil)
		ctx.Step(`^que eu enviei documentos com problemas de qualidade$`, queEuEnvieiDocumentosComProblemasDeQualidade)
		ctx.Step(`^que eu enviei documentos válidos$`, queEuEnvieiDocumentosVlidos)
		ctx.Step(`^que eu estou na página de cadastro de motorista$`, queEuEstouNaPginaDeCadastroDeMotorista)
		ctx.Step(`^que existe um motorista cadastrado com "([^"]*)" "(\d+)"$`, queExisteUmMotoristaCadastradoCom)
		ctx.Step(`^que existe um motorista cadastrado com "([^"]*)" "(\d+)\.(\d+)\.(\d+)-(\d+)"$`, queExisteUmMotoristaCadastradoCom)
		ctx.Step(`^que existe um motorista cadastrado com "([^"]*)" "joao\.silva@email\.com"$`, queExisteUmMotoristaCadastradoComJoaosilvaemailcom)
		ctx.Step(`^que existe um motorista cadastrado com os dados:$`, queExisteUmMotoristaCadastradoComOsDados)
		ctx.Step(`^que meu status é "([^"]*)"$`, queMeuStatus)
		ctx.Step(`^que redefini minha senha usando um link de recuperação$`, queRedefiniMinhaSenhaUsandoUmLinkDeRecuperao)
		ctx.Step(`^que solicitei recuperação de conta há (\d+) horas$`, queSoliciteiRecuperaoDeContaHHoras)
		ctx.Step(`^que solicitei recuperação de conta há (\d+) minutos$`, queSoliciteiRecuperaoDeContaHMinutos)
		ctx.Step(`^recebo um email com instruções para confirmar a exclusão$`, receboUmEmailComInstruesParaConfirmarAExcluso)
		ctx.Step(`^recebo um email de confirmação$`, receboUmEmailDeConfirmao)
		ctx.Step(`^salvo a alteração$`, salvoAAlterao)
		ctx.Step(`^solicito excluir minha conta permanentemente$`, solicitoExcluirMinhaContaPermanentemente)
		ctx.Step(`^solicito recuperação novamente com o mesmo email$`, solicitoRecuperaoNovamenteComOMesmoEmail)
		ctx.Step(`^submeto o formulário com email "([^"]*)"$`, submetoOFormulrioComEmail)
		ctx.Step(`^submeto o formulário$`, submetoOFormulrio)
		ctx.Step(`^tento acessar o mesmo link novamente$`, tentoAcessarOMesmoLinkNovamente)
		ctx.Step(`^tento fazer upload de arquivo com "([^"]*)"$`, tentoFazerUploadDeArquivoCom)
		ctx.Step(`^tento salvar a alteração$`, tentoSalvarAAlterao)
		ctx.Step(`^tento salvar as alterações$`, tentoSalvarAsAlteraes)
		ctx.Step(`^um email de recuperação é enviado$`, umEmailDeRecuperaoEnviado)
		ctx.Step(`^vejo "([^"]*)"$`, vejo)
		ctx.Step(`^vejo a mensagem "([^"]*)"$`, vejoAMensagem)
		ctx.Step(`^vejo a mensagem de erro "([^"]*)"$`, vejoAMensagemDeErro)
		ctx.Step(`^vejo a opção de solicitar um novo link$`, vejoAOpoDeSolicitarUmNovoLink)
		ctx.Step(`^vejo erro "([^"]*)"$`, vejoErro)
		ctx.Step(`^vejo estatísticas sobre minhas corridas$`, vejoEstatsticasSobreMinhasCorridas)
		ctx.Step(`^vejo meu dado "([^"]*)" com o valor "([^"]*)":$`, vejoMeuDadoComOValor)
		ctx.Step(`^vejo minha foto de perfil$`, vejoMinhaFotoDePerfil)
		ctx.Step(`^vejo que devo aguardar antes de solicitar novamente$`, vejoQueDevoAguardarAntesDeSolicitarNovamente)
}
