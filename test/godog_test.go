package test

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../features"},
			Output:   colors.Colored(os.Stdout),
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a conta é bloqueada temporariamente por (\d+) minutos$`, aContaBloqueadaTemporariamentePorMinutos)
	ctx.Step(`^a corrida deve ser registrada no histórico como "([^"]*)"$`, aCorridaDeveSerRegistradaNoHistricoComo)
	ctx.Step(`^a corrida for carregada na tela$`, aCorridaForCarregadaNaTela)
	ctx.Step(`^a corrida é marcada como avaliada no histórico\.$`, aCorridaMarcadaComoAvaliadaNoHistrico)
	ctx.Step(`^a mensagem de erro desaparece$`, aMensagemDeErroDesaparece)
	ctx.Step(`^a minha localização atual coincide com o destino da corrida$`, aMinhaLocalizaoAtualCoincideComODestinoDaCorrida)
	ctx.Step(`^a notificação deve desaparecer$`, aNotificaoDeveDesaparecer)
	ctx.Step(`^aceita a corrida$`, aceitaACorrida)
	ctx.Step(`^aceitei uma corrida com (\d+) km de distância estimada até o local de embarque$`, aceiteiUmaCorridaComKmDeDistnciaEstimadaAtOLocalDeEmbarque)
	ctx.Step(`^aceitei uma corrida que está em andamento$`, aceiteiUmaCorridaQueEstEmAndamento)
	ctx.Step(`^acessei o link de recuperação válido$`, acesseiOLinkDeRecuperaoVlido)
	ctx.Step(`^acesso a tela de histórico de corridas$`, acessoATelaDeHistricoDeCorridas)
	ctx.Step(`^acesso histórico de corridas$`, acessoHistricoDeCorridas)
	ctx.Step(`^acesso meu histórico de corridas$`, acessoMeuHistricoDeCorridas)
	ctx.Step(`^acesso meu perfil$`, acessoMeuPerfil)
	ctx.Step(`^acesso o link de recuperação expirado$`, acessoOLinkDeRecuperaoExpirado)
	ctx.Step(`^ativa a opção "([^"]*)"$`, ativaAOpo)
	ctx.Step(`^atualizo as seguintes informações:$`, atualizoAsSeguintesInformaes)
	ctx.Step(`^cada item da lista exibe destino, data, hora e valor da corrida\.$`, cadaItemDaListaExibeDestinoDataHoraEValorDaCorrida)
	ctx.Step(`^cancelei uma corrida que havia aceitado com destino de (\d+) km e valor estimado de R\$ (\d+),(\d+)$`, canceleiUmaCorridaQueHaviaAceitadoComDestinoDeKmEValorEstimadoDeR)
	ctx.Step(`^clico em "([^"]*)"$`, clicoEm)
	ctx.Step(`^clico em "([^"]*)" ao lado da corrida pendente$`, clicoEmAoLadoDaCorridaPendente)
	ctx.Step(`^clico na opção "([^"]*)"$`, clicoNaOpo)
	ctx.Step(`^confirmo a desativação$`, confirmoADesativao)
	ctx.Step(`^confirmo a exclusão fornecendo a justificativa "([^"]*)"$`, confirmoAExclusoFornecendoAJustificativa)
	ctx.Step(`^confirmo a nova senha "([^"]*)"$`, confirmoANovaSenha)
	ctx.Step(`^confirmo a reativação$`, confirmoAReativao)
	ctx.Step(`^deve continuar recebendo notificações de novas corridas próximas$`, deveContinuarRecebendoNotificaesDeNovasCorridasPrximas)
	ctx.Step(`^deve receber uma notificação com os detalhes da corrida$`, deveReceberUmaNotificaoComOsDetalhesDaCorrida)
	ctx.Step(`^deve ter a opção de aceitar ou recusar a corrida$`, deveTerAOpoDeAceitarOuRecusarACorrida)
	ctx.Step(`^deve voltar a receber notificações de novas corridas$`, deveVoltarAReceberNotificaesDeNovasCorridas)
	ctx.Step(`^enquanto o status de “João” estiver como "([^"]*)", não deve receber notificações de novas corridas$`, enquantoOStatusDeJooEstiverComoNoDeveReceberNotificaesDeNovasCorridas)
	ctx.Step(`^escolho uma nota de (\d+) a (\d+) estrelas e adiciono um comentário opcional$`, escolhoUmaNotaDeAEstrelasEAdicionoUmComentrioOpcional)
	ctx.Step(`^essa entrada apresenta a data e hora do cancelamento, distância estimada e valor previsto da corrida$`, essaEntradaApresentaADataEHoraDoCancelamentoDistnciaEstimadaEValorPrevistoDaCorrida)
	ctx.Step(`^essa notificação aparece com as opções "([^"]*)" e "([^"]*)"$`, essaNotificaoApareceComAsOpesE)
	ctx.Step(`^está com o status "([^"]*)"$`, estComOStatus)
	ctx.Step(`^estou na tela de avaliação de corridas$`, estouNaTelaDeAvaliaoDeCorridas)
	ctx.Step(`^estou na tela de histórico de corridas$`, estouNaTelaDeHistricoDeCorridas)
	ctx.Step(`^estou na tela de menu principal$`, estouNaTelaDeMenuPrincipal)
	ctx.Step(`^estou visualizando o tempo estimado de chegada de (\d+) minutos$`, estouVisualizandoOTempoEstimadoDeChegadaDeMinutos)
	ctx.Step(`^eu acesso meu perfil no navegador web$`, euAcessoMeuPerfilNoNavegadorWeb)
	ctx.Step(`^eu acesso o histórico de alterações$`, euAcessoOHistricoDeAlteraes)
	ctx.Step(`^eu acesso o link de cancelamento no email recebido$`, euAcessoOLinkDeCancelamentoNoEmailRecebido)
	ctx.Step(`^eu altero a senha para "([^"]*)"$`, euAlteroASenhaPara)
	ctx.Step(`^eu cancelo a solicitação$`, euCanceloASolicitao)
	ctx.Step(`^eu clico "([^"]*)",$`, euClico)
	ctx.Step(`^eu clico em "([^"]*)"$`, euClicoEm)
	ctx.Step(`^eu confirmo a desativação$`, euConfirmoADesativao)
	ctx.Step(`^eu confirmo a reativação$`, euConfirmoAReativao)
	ctx.Step(`^eu confirmo digitando "([^"]*)"$`, euConfirmoDigitando)
	ctx.Step(`^eu desativo a opção "([^"]*)"$`, euDesativoAOpo)
	ctx.Step(`^eu estou na minha área do motorista$`, euEstouNaMinhaReaDoMotorista)
	ctx.Step(`^eu estou na página de upload de documentos$`, euEstouNaPginaDeUploadDeDocumentos)
	ctx.Step(`^eu faço upload da foto da CNH em formato JPG com (\d+)MB$`, euFaoUploadDaFotoDaCNHEmFormatoJPGComMB)
	ctx.Step(`^eu faço upload da foto do CRLV em formato PNG com (\d+)\.(\d+)MB$`, euFaoUploadDaFotoDoCRLVEmFormatoPNGComMB)
	ctx.Step(`^eu faço upload de uma selfie com CNH em formato JPG com (\d+)MB$`, euFaoUploadDeUmaSelfieComCNHEmFormatoJPGComMB)
	ctx.Step(`^eu fico inativo por mais de (\d+) minutos$`, euFicoInativoPorMaisDeMinutos)
	ctx.Step(`^eu não recebo novas solicitações de corrida$`, euNoReceboNovasSolicitaesDeCorrida)
	ctx.Step(`^eu perco a conexão com a internet por (\d+) minutos$`, euPercoAConexoComAInternetPorMinutos)
	ctx.Step(`^eu permaneço logado$`, euPermaneoLogado)
	ctx.Step(`^eu posso fazer login no sistema$`, euPossoFazerLoginNoSistema)
	ctx.Step(`^eu posso fazer upload da nova CNH$`, euPossoFazerUploadDaNovaCNH)
	ctx.Step(`^eu posso reenviar os documentos corrigidos$`, euPossoReenviarOsDocumentosCorrigidos)
	ctx.Step(`^eu preencho a confirmação de senha "([^"]*)"$`, euPreenchoAConfirmaoDeSenha)
	ctx.Step(`^eu preencho a senha "([^"]*)"$`, euPreenchoASenha)
	ctx.Step(`^eu preencho o motivo "([^"]*)"$`, euPreenchoOMotivo)
	ctx.Step(`^eu realizo login com email "([^"]*)" e senha "([^"]*)"$`, euRealizoLoginComEmailESenha)
	ctx.Step(`^eu recebo um email "([^"]*)"$`, euReceboUmEmail)
	ctx.Step(`^eu recebo um email com instruções para confirmar a exclusão$`, euReceboUmEmailComInstruesParaConfirmarAExcluso)
	ctx.Step(`^eu recebo um email confirmando o recebimento dos documentos$`, euReceboUmEmailConfirmandoORecebimentoDosDocumentos)
	ctx.Step(`^eu recebo um email detalhando os problemas encontrados$`, euReceboUmEmailDetalhandoOsProblemasEncontrados)
	ctx.Step(`^eu recupero a conexão com a internet$`, euRecuperoAConexoComAInternet)
	ctx.Step(`^eu seleciono o motivo "([^"]*)"$`, euSelecionoOMotivo)
	ctx.Step(`^eu sou direcionado para a página de upload de documentos$`, euSouDirecionadoParaAPginaDeUploadDeDocumentos)
	ctx.Step(`^eu sou redirecionado para a página de login$`, euSouRedirecionadoParaAPginaDeLogin)
	ctx.Step(`^eu sou redirecionado para o meu próprio perfil$`, euSouRedirecionadoParaOMeuPrprioPerfil)
	ctx.Step(`^eu submeto o formulário$`, euSubmetoOFormulrio)
	ctx.Step(`^eu tento acessar diretamente a URL do perfil do motorista "([^"]*)"$`, euTentoAcessarDiretamenteAURLDoPerfilDoMotorista)
	ctx.Step(`^eu tento acessar o dashboard$`, euTentoAcessarODashboard)
	ctx.Step(`^eu tento acessar uma funcionalidade do sistema$`, euTentoAcessarUmaFuncionalidadeDoSistema)
	ctx.Step(`^eu tento fazer upload de um arquivo de (\d+)MB$`, euTentoFazerUploadDeUmArquivoDeMB)
	ctx.Step(`^eu tento fazer upload de um arquivo em formato TXT$`, euTentoFazerUploadDeUmArquivoEmFormatoTXT)
	ctx.Step(`^eu vejo "([^"]*)"$`, euVejo)
	ctx.Step(`^eu vejo a data\/hora da última sincronização$`, euVejoADatahoraDaLtimaSincronizao)
	ctx.Step(`^eu vejo a mensagem "([^"]*)"$`, euVejoAMensagem)
	ctx.Step(`^eu vejo a mensagem de erro "([^"]*)"$`, euVejoAMensagemDeErro)
	ctx.Step(`^eu vejo a opção "([^"]*)"$`, euVejoAOpo)
	ctx.Step(`^eu vejo apenas as alterações dos últimos (\d+) dias$`, euVejoApenasAsAlteraesDosLtimosDias)
	ctx.Step(`^eu vejo o botão "([^"]*)"$`, euVejoOBoto)
	ctx.Step(`^eu vejo o indicador de força da senha como "([^"]*)"$`, euVejoOIndicadorDeForaDaSenhaComo)
	ctx.Step(`^eu vejo o status "([^"]*)"$`, euVejoOStatus)
	ctx.Step(`^eu vejo o telefone atualizado também no navegador$`, euVejoOTelefoneAtualizadoTambmNoNavegador)
	ctx.Step(`^eu vejo o tutorial inicial do motorista$`, euVejoOTutorialInicialDoMotorista)
	ctx.Step(`^eu vejo quais documentos específicos foram rejeitados$`, euVejoQuaisDocumentosEspecficosForamRejeitados)
	ctx.Step(`^eu vejo um alerta "([^"]*)"$`, euVejoUmAlerta)
	ctx.Step(`^eu vejo um aviso detalhado sobre as consequências da exclusão$`, euVejoUmAvisoDetalhadoSobreAsConsequnciasDaExcluso)
	ctx.Step(`^eu vejo um formulário para informar o motivo da exclusão$`, euVejoUmFormulrioParaInformarOMotivoDaExcluso)
	ctx.Step(`^eu vejo uma lista cronológica das modificações:$`, euVejoUmaListaCronolgicaDasModificaes)
	ctx.Step(`^eu vejo uma notificação "([^"]*)"$`, euVejoUmaNotificao)
	ctx.Step(`^eu volto a receber solicitações de corrida$`, euVoltoAReceberSolicitaesDeCorrida)
	ctx.Step(`^existe uma corrida que já foi avaliada anteriormente$`, existeUmaCorridaQueJFoiAvaliadaAnteriormente)
	ctx.Step(`^faço upload de uma nova foto em formato JPG com (\d+)MB$`, faoUploadDeUmaNovaFotoEmFormatoJPGComMB)
	ctx.Step(`^faço upload dos documentos obrigatórios:$`, faoUploadDosDocumentosObrigatrios)
	ctx.Step(`^fico inativo por mais de (\d+) minutos$`, ficoInativoPorMaisDeMinutos)
	ctx.Step(`^foi ao local de embarque "([^"]*)" e chegou ao destino "([^"]*)"$`, foiAoLocalDeEmbarqueEChegouAoDestino)
	ctx.Step(`^foi ao local de embarque "([^"]*)" e iniciou o trajeto para o destino "([^"]*)"$`, foiAoLocalDeEmbarqueEIniciouOTrajetoParaODestino)
	ctx.Step(`^meu status inicial é "([^"]*)"$`, meuStatusInicial)
	ctx.Step(`^meu status muda para "([^"]*)"$`, meuStatusMudaPara)
	ctx.Step(`^meus dados são atualizados no sistema$`, meusDadosSoAtualizadosNoSistema)
	ctx.Step(`^meus dados são sincronizados automaticamente$`, meusDadosSoSincronizadosAutomaticamente)
	ctx.Step(`^minha conta está com status "([^"]*)"$`, minhaContaEstComStatus)
	ctx.Step(`^minha conta fica com status "([^"]*)"$`, minhaContaFicaComStatus)
	ctx.Step(`^minha conta volta ao status "([^"]*)"$`, minhaContaVoltaAoStatus)
	ctx.Step(`^motorista demora mais do que "([^"]*)" minutos para chegar ao local$`, motoristaDemoraMaisDoQueMinutosParaChegarAoLocal)
	ctx.Step(`^nenhum email de recuperação é enviado$`, nenhumEmailDeRecuperaoEnviado)
	ctx.Step(`^nenhum email é enviado$`, nenhumEmailEnviado)
	ctx.Step(`^não consigo concluir o cadastro$`, noConsigoConcluirOCadastro)
	ctx.Step(`^não consigo fazer login mesmo com credenciais corretas$`, noConsigoFazerLoginMesmoComCredenciaisCorretas)
	ctx.Step(`^não deve receber notificações de novas corridas$`, noDeveReceberNotificaesDeNovasCorridas)
	ctx.Step(`^não deve ser registrada nenhuma corrida no histórico$`, noDeveSerRegistradaNenhumaCorridaNoHistrico)
	ctx.Step(`^não interage com a notificação em até “(\d+)” segundos$`, noInterageComANotificaoEmAtSegundos)
	ctx.Step(`^não sou autenticado$`, noSouAutenticado)
	ctx.Step(`^não vejo o botão "([^"]*)"$`, noVejoOBoto)
	ctx.Step(`^o email contém um link válido por (\d+) hora$`, oEmailContmUmLinkVlidoPorHora)
	ctx.Step(`^o link de recuperação é invalidado$`, oLinkDeRecuperaoInvalidado)
	ctx.Step(`^o motorista de nome "([^"]*)" aceitou uma corrida$`, oMotoristaDeNomeAceitouUmaCorrida)
	ctx.Step(`^o motorista de nome “João” está logado$`, oMotoristaDeNomeJooEstLogado)
	ctx.Step(`^o motorista está a "([^"]*)" minutos do local de embarque "([^"]*)"$`, oMotoristaEstAMinutosDoLocalDeEmbarque)
	ctx.Step(`^o motorista recebe a mensagem de confirmação na tela "([^"]*)"$`, oMotoristaRecebeAMensagemDeConfirmaoNaTela)
	ctx.Step(`^o motorista recebe a mensagem de erro na tela "([^"]*)"$`, oMotoristaRecebeAMensagemDeErroNaTela)
	ctx.Step(`^o motorista seleciona a opção "([^"]*)"$`, oMotoristaSelecionaAOpo)
	ctx.Step(`^o motorista tenta cancelar a corrida selecionando a opção "([^"]*)"$`, oMotoristaTentaCancelarACorridaSelecionandoAOpo)
	ctx.Step(`^o pop-up exibe os botões "([^"]*)" e  "([^"]*)"$`, oPopupExibeOsBotesE)
	ctx.Step(`^o pop-up é fechado\.$`, oPopupFechado)
	ctx.Step(`^o serviço de validação de documentos está disponível$`, oServioDeValidaoDeDocumentosEstDisponvel)
	ctx.Step(`^o sistema cancela automaticamente a corrida$`, oSistemaCancelaAutomaticamenteACorrida)
	ctx.Step(`^o sistema cancela o progresso da corrida$`, oSistemaCancelaOProgressoDaCorrida)
	ctx.Step(`^o sistema de envio de emails está funcionando$`, oSistemaDeEnvioDeEmailsEstFuncionando)
	ctx.Step(`^o sistema de verificação automática está funcionando$`, oSistemaDeVerificaoAutomticaEstFuncionando)
	ctx.Step(`^o sistema exibe uma entrada no histórico com o status "([^"]*)"$`, oSistemaExibeUmaEntradaNoHistricoComOStatus)
	ctx.Step(`^o sistema exibe uma notificação na tela com a mensagem:$`, oSistemaExibeUmaNotificaoNaTelaComAMensagem)
	ctx.Step(`^o sistema não permite o cancelamento$`, oSistemaNoPermiteOCancelamento)
	ctx.Step(`^o sistema ou analista verifica meus documentos$`, oSistemaOuAnalistaVerificaMeusDocumentos)
	ctx.Step(`^o sistema registra a tentativa de acesso não autorizado$`, oSistemaRegistraATentativaDeAcessoNoAutorizado)
	ctx.Step(`^o sistema valida meus documentos$`, oSistemaValidaMeusDocumentos)
	ctx.Step(`^o status deve estar “disponível”$`, oStatusDeveEstarDisponvel)
	ctx.Step(`^o status deve estar “ocupado”$`, oStatusDeveEstarOcupado)
	ctx.Step(`^o status deve permanecer como "([^"]*)"$`, oStatusDevePermanecerComo)
	ctx.Step(`^o status deve ser alterado para "([^"]*)"$`, oStatusDeveSerAlteradoPara)
	ctx.Step(`^o upload é cancelado$`, oUploadCancelado)
	ctx.Step(`^o upload não é concluído$`, oUploadNoConcludo)
	ctx.Step(`^o valor total e a distância percorrida são exibidos$`, oValorTotalEADistnciaPercorridaSoExibidos)
	ctx.Step(`^os campos de senha são destacados$`, osCamposDeSenhaSoDestacados)
	ctx.Step(`^os dados da avaliação anterior são exibidos$`, osDadosDaAvaliaoAnteriorSoExibidos)
	ctx.Step(`^posso fazer login com a nova senha "([^"]*)"$`, possoFazerLoginComANovaSenha)
	ctx.Step(`^preencho:$`, preencho)
	ctx.Step(`^preencho a nova senha "([^"]*)"$`, preenchoANovaSenha)
	ctx.Step(`^preencho o campo "([^"]*)" com "([^"]*)"$`, preenchoOCampoCom)
	ctx.Step(`^preencho o campo "([^"]*)" com o valor "([^"]*)"$`, preenchoOCampoComOValor)
	ctx.Step(`^preencho os campos de senha:$`, preenchoOsCamposDeSenha)
	ctx.Step(`^que aceitei uma corrida disponível$`, queAceiteiUmaCorridaDisponvel)
	ctx.Step(`^que completei o cadastro básico com sucesso$`, queCompleteiOCadastroBsicoComSucesso)
	ctx.Step(`^que estou autenticado como motorista:$`, queEstouAutenticadoComoMotorista)
	ctx.Step(`^que estou logado como motorista$`, queEstouLogadoComoMotorista)
	ctx.Step(`^que estou logado como o motorista João$`, queEstouLogadoComoOMotoristaJoo)
	ctx.Step(`^que estou logado no sistema como motorista$`, queEstouLogadoNoSistemaComoMotorista)
	ctx.Step(`^que estou logado no sistema como o motorista João$`, queEstouLogadoNoSistemaComoOMotoristaJoo)
	ctx.Step(`^que estou na página "([^"]*)"$`, queEstouNaPgina)
	ctx.Step(`^que estou na página "([^"]*)" através de link válido$`, queEstouNaPginaAtravsDeLinkVlido)
	ctx.Step(`^que estou na página de alteração de senha$`, queEstouNaPginaDeAlteraoDeSenha)
	ctx.Step(`^que estou na página de edição do perfil$`, queEstouNaPginaDeEdioDoPerfil)
	ctx.Step(`^que estou na página de upload de documentos$`, queEstouNaPginaDeUploadDeDocumentos)
	ctx.Step(`^que estou na página do meu perfil$`, queEstouNaPginaDoMeuPerfil)
	ctx.Step(`^que eu altero meu telefone no dispositivo mobile$`, queEuAlteroMeuTelefoneNoDispositivoMobile)
	ctx.Step(`^que eu completei o cadastro básico com sucesso$`, queEuCompleteiOCadastroBsicoComSucesso)
	ctx.Step(`^que eu enviei documentos com problemas de qualidade$`, queEuEnvieiDocumentosComProblemasDeQualidade)
	ctx.Step(`^que eu enviei todos os documentos obrigatórios$`, queEuEnvieiTodosOsDocumentosObrigatrios)
	ctx.Step(`^que eu estou logado como motorista "([^"]*)"$`, queEuEstouLogadoComoMotorista)
	ctx.Step(`^que eu estou na página de cadastro de motorista$`, queEuEstouNaPginaDeCadastroDeMotorista)
	ctx.Step(`^que eu estou na página de upload de documentos$`, queEuEstouNaPginaDeUploadDeDocumentos)
	ctx.Step(`^que eu estou na página do meu perfil$`, queEuEstouNaPginaDoMeuPerfil)
	ctx.Step(`^que eu solicitei exclusão da conta há (\d+) dias$`, queEuSoliciteiExclusoDaContaHDias)
	ctx.Step(`^que existe outro motorista com ID "([^"]*)"$`, queExisteOutroMotoristaComID)
	ctx.Step(`^que existe um motorista cadastrado:$`, queExisteUmMotoristaCadastrado)
	ctx.Step(`^que existe um motorista cadastrado com cnh "([^"]*)"$`, queExisteUmMotoristaCadastradoComCnh)
	ctx.Step(`^que existe um motorista cadastrado com cpf "([^"]*)"$`, queExisteUmMotoristaCadastradoComCpf)
	ctx.Step(`^que existe um motorista cadastrado com email "([^"]*)"$`, queExisteUmMotoristaCadastradoComEmail)
	ctx.Step(`^que existe um motorista cadastrado com os dados:$`, queExisteUmMotoristaCadastradoComOsDados)
	ctx.Step(`^que existe um motorista com status "([^"]*)":$`, queExisteUmMotoristaComStatus)
	ctx.Step(`^que finalizei uma corrida após chegar ao destino$`, queFinalizeiUmaCorridaApsChegarAoDestino)
	ctx.Step(`^que minha CNH está próxima do vencimento \((\d+) dias\)$`, queMinhaCNHEstPrximaDoVencimentoDias)
	ctx.Step(`^que minha conta está com status "([^"]*)"$`, queMinhaContaEstComStatus)
	ctx.Step(`^que o email "([^"]*)" não está bloqueado$`, queOEmailNoEstBloqueado)
	ctx.Step(`^que o motorista de nome “João” está logado$`, queOMotoristaDeNomeJooEstLogado)
	ctx.Step(`^que o sistema de cadastro está funcionando$`, queOSistemaDeCadastroEstFuncionando)
	ctx.Step(`^que redefini minha senha usando um link de recuperação$`, queRedefiniMinhaSenhaUsandoUmLinkDeRecuperao)
	ctx.Step(`^que solicitei recuperação de conta há (\d+) horas$`, queSoliciteiRecuperaoDeContaHHoras)
	ctx.Step(`^que solicitei recuperação de conta há (\d+) minutos$`, queSoliciteiRecuperaoDeContaHMinutos)
	ctx.Step(`^realizo login com email "([^"]*)" e senha "([^"]*)"$`, realizoLoginComEmailESenha)
	ctx.Step(`^realizo o cadastro com dados válidos:$`, realizoOCadastroComDadosVlidos)
	ctx.Step(`^realizo (\d+) tentativas de login com senha incorreta para "([^"]*)"$`, realizoTentativasDeLoginComSenhaIncorretaPara)
	ctx.Step(`^recebe uma notificação de corrida$`, recebeUmaNotificaoDeCorrida)
	ctx.Step(`^recebe uma notificação de uma nova corrida próxima$`, recebeUmaNotificaoDeUmaNovaCorridaPrxima)
	ctx.Step(`^recebi uma notificação de uma nova corrida próxima$`, recebiUmaNotificaoDeUmaNovaCorridaPrxima)
	ctx.Step(`^recebo um email com instruções para confirmar a exclusão$`, receboUmEmailComInstruesParaConfirmarAExcluso)
	ctx.Step(`^recebo um email de confirmação$`, receboUmEmailDeConfirmao)
	ctx.Step(`^recusa a corrida$`, recusaACorrida)
	ctx.Step(`^salvo a alteração$`, salvoAAlterao)
	ctx.Step(`^salvo a nova foto$`, salvoANovaFoto)
	ctx.Step(`^salvo as alterações$`, salvoAsAlteraes)
	ctx.Step(`^seleciona a opção "([^"]*)"$`, selecionaAOpo)
	ctx.Step(`^seleciona "([^"]*)" na interface de desencentivo a cancelamento$`, selecionaNaInterfaceDeDesencentivoACancelamento)
	ctx.Step(`^selecionei uma corrida ainda não avaliada$`, selecioneiUmaCorridaAindaNoAvaliada)
	ctx.Step(`^seleciono o motivo "([^"]*)"$`, selecionoOMotivo)
	ctx.Step(`^solicito desativar minha conta temporariamente$`, solicitoDesativarMinhaContaTemporariamente)
	ctx.Step(`^solicito excluir minha conta permanentemente$`, solicitoExcluirMinhaContaPermanentemente)
	ctx.Step(`^solicito reativar minha conta$`, solicitoReativarMinhaConta)
	ctx.Step(`^solicito recuperação novamente com o mesmo email$`, solicitoRecuperaoNovamenteComOMesmoEmail)
	ctx.Step(`^solicito recuperação para "([^"]*)"$`, solicitoRecuperaoPara)
	ctx.Step(`^sou autenticado com sucesso$`, souAutenticadoComSucesso)
	ctx.Step(`^sou redirecionado para a página de login$`, souRedirecionadoParaAPginaDeLogin)
	ctx.Step(`^sou redirecionado para a tela de avaliação de corrida$`, souRedirecionadoParaATelaDeAvaliaoDeCorrida)
	ctx.Step(`^submeto o formulário$`, submetoOFormulrio)
	ctx.Step(`^submeto o formulário com email "([^"]*)"$`, submetoOFormulrioComEmail)
	ctx.Step(`^tenho pelo menos uma corrida realizada que ainda não foi avaliada$`, tenhoPeloMenosUmaCorridaRealizadaQueAindaNoFoiAvaliada)
	ctx.Step(`^tenho uma corrida anterior ainda não avaliada$`, tenhoUmaCorridaAnteriorAindaNoAvaliada)
	ctx.Step(`^tento acessar o mesmo link novamente$`, tentoAcessarOMesmoLinkNovamente)
	ctx.Step(`^tento acessar uma funcionalidade do sistema$`, tentoAcessarUmaFuncionalidadeDoSistema)
	ctx.Step(`^tento fazer upload de arquivo "([^"]*)"$`, tentoFazerUploadDeArquivo)
	ctx.Step(`^tento fazer upload de arquivo com "([^"]*)"$`, tentoFazerUploadDeArquivoCom)
	ctx.Step(`^tento realizar cadastro com CNH válida até "([^"]*)"$`, tentoRealizarCadastroComCNHVlidaAt)
	ctx.Step(`^tento realizar cadastro com cnh "([^"]*)"$`, tentoRealizarCadastroComCnh)
	ctx.Step(`^tento realizar cadastro com cpf "([^"]*)"$`, tentoRealizarCadastroComCpf)
	ctx.Step(`^tento realizar cadastro com data de nascimento "([^"]*)"$`, tentoRealizarCadastroComDataDeNascimento)
	ctx.Step(`^tento realizar cadastro com email "([^"]*)"$`, tentoRealizarCadastroComEmail)
	ctx.Step(`^tento realizar cadastro sem preencher o campo "([^"]*)"$`, tentoRealizarCadastroSemPreencherOCampo)
	ctx.Step(`^tento salvar a alteração$`, tentoSalvarAAlterao)
	ctx.Step(`^tento salvar as alterações$`, tentoSalvarAsAlteraes)
	ctx.Step(`^todos os documentos são aprovados automaticamente$`, todosOsDocumentosSoAprovadosAutomaticamente)
	ctx.Step(`^um email de recuperação é enviado para "([^"]*)"$`, umEmailDeRecuperaoEnviadoPara)
	ctx.Step(`^uma mensagem de erro aparece na tela "([^"]*)"$`, umaMensagemDeErroApareceNaTela)
	ctx.Step(`^uma notificação é enviada para o email do motorista sobre tentativa de acesso$`, umaNotificaoEnviadaParaOEmailDoMotoristaSobreTentativaDeAcesso)
	ctx.Step(`^uma nova corrida é solicitada por um passageiro próximo$`, umaNovaCorridaSolicitadaPorUmPassageiroPrximo)
	ctx.Step(`^vejo a corrida com data e hora$`, vejoACorridaComDataEHora)
	ctx.Step(`^vejo a mensagem "([^"]*)"$`, vejoAMensagem)
	ctx.Step(`^vejo a mensagem de erro "([^"]*)"$`, vejoAMensagemDeErro)
	ctx.Step(`^vejo a notificação "([^"]*)"$`, vejoANotificao)
	ctx.Step(`^vejo a opção de solicitar um novo link$`, vejoAOpoDeSolicitarUmNovoLink)
	ctx.Step(`^vejo "([^"]*)" como motorista logado$`, vejoComoMotoristaLogado)
	ctx.Step(`^vejo estatísticas sobre minhas corridas$`, vejoEstatsticasSobreMinhasCorridas)
	ctx.Step(`^vejo meus dados pessoais cadastrados$`, vejoMeusDadosPessoaisCadastrados)
	ctx.Step(`^vejo minha foto de perfil$`, vejoMinhaFotoDePerfil)
	ctx.Step(`^vejo o dashboard do motorista$`, vejoODashboardDoMotorista)
	ctx.Step(`^vejo o tempo estimado de chegada até o local de embarque$`, vejoOTempoEstimadoDeChegadaAtOLocalDeEmbarque)
	ctx.Step(`^vejo os campos de nota e comentário disponíveis para preenchimento\.$`, vejoOsCamposDeNotaEComentrioDisponveisParaPreenchimento)
	ctx.Step(`^vejo que devo aguardar antes de solicitar novamente$`, vejoQueDevoAguardarAntesDeSolicitarNovamente)
	ctx.Step(`^vejo um pop-up escrito "([^"]*)"$`, vejoUmPopupEscrito)
	ctx.Step(`^vejo uma lista contendo as últimas corridas realizadas$`, vejoUmaListaContendoAsLtimasCorridasRealizadas)
	ctx.Step(`^vejo uma mensagem "([^"]*)"$`, vejoUmaMensagem)
	ctx.Step(`^visualizo os detalhes dessa corrida$`, visualizoOsDetalhesDessaCorrida)
}

func aContaBloqueadaTemporariamentePorMinutos(arg1 int) error {
	return godog.ErrPending
}

func aCorridaDeveSerRegistradaNoHistricoComo(arg1 string) error {
	return godog.ErrPending
}

func aCorridaForCarregadaNaTela() error {
	return godog.ErrPending
}

func aCorridaMarcadaComoAvaliadaNoHistrico() error {
	return godog.ErrPending
}

func aMensagemDeErroDesaparece() error {
	return godog.ErrPending
}

func aMinhaLocalizaoAtualCoincideComODestinoDaCorrida() error {
	return godog.ErrPending
}

func aNotificaoDeveDesaparecer() error {
	return godog.ErrPending
}

func aceitaACorrida() error {
	return godog.ErrPending
}

func aceiteiUmaCorridaComKmDeDistnciaEstimadaAtOLocalDeEmbarque(arg1 int) error {
	return godog.ErrPending
}

func aceiteiUmaCorridaQueEstEmAndamento() error {
	return godog.ErrPending
}

func acesseiOLinkDeRecuperaoVlido() error {
	return godog.ErrPending
}

func acessoATelaDeHistricoDeCorridas() error {
	return godog.ErrPending
}

func acessoHistricoDeCorridas() error {
	return godog.ErrPending
}

func acessoMeuHistricoDeCorridas() error {
	return godog.ErrPending
}

func acessoMeuPerfil() error {
	return godog.ErrPending
}

func acessoOLinkDeRecuperaoExpirado() error {
	return godog.ErrPending
}

func ativaAOpo(arg1 string) error {
	return godog.ErrPending
}

func atualizoAsSeguintesInformaes(arg1 *godog.Table) error {
	return godog.ErrPending
}

func cadaItemDaListaExibeDestinoDataHoraEValorDaCorrida() error {
	return godog.ErrPending
}

func canceleiUmaCorridaQueHaviaAceitadoComDestinoDeKmEValorEstimadoDeR(arg1, arg2, arg3 int) error {
	return godog.ErrPending
}

func clicoEm(arg1 string) error {
	return godog.ErrPending
}

func clicoEmAoLadoDaCorridaPendente(arg1 string) error {
	return godog.ErrPending
}

func clicoNaOpo(arg1 string) error {
	return godog.ErrPending
}

func confirmoADesativao() error {
	return godog.ErrPending
}

func confirmoAExclusoFornecendoAJustificativa(arg1 string) error {
	return godog.ErrPending
}

func confirmoANovaSenha(arg1 string) error {
	return godog.ErrPending
}

func confirmoAReativao() error {
	return godog.ErrPending
}

func deveContinuarRecebendoNotificaesDeNovasCorridasPrximas() error {
	return godog.ErrPending
}

func deveReceberUmaNotificaoComOsDetalhesDaCorrida() error {
	return godog.ErrPending
}

func deveTerAOpoDeAceitarOuRecusarACorrida() error {
	return godog.ErrPending
}

func deveVoltarAReceberNotificaesDeNovasCorridas() error {
	return godog.ErrPending
}

func enquantoOStatusDeJooEstiverComoNoDeveReceberNotificaesDeNovasCorridas(arg1 string) error {
	return godog.ErrPending
}

func escolhoUmaNotaDeAEstrelasEAdicionoUmComentrioOpcional(arg1, arg2 int) error {
	return godog.ErrPending
}

func essaEntradaApresentaADataEHoraDoCancelamentoDistnciaEstimadaEValorPrevistoDaCorrida() error {
	return godog.ErrPending
}

func essaNotificaoApareceComAsOpesE(arg1, arg2 string) error {
	return godog.ErrPending
}

func estComOStatus(arg1 string) error {
	return godog.ErrPending
}

func estouNaTelaDeAvaliaoDeCorridas() error {
	return godog.ErrPending
}

func estouNaTelaDeHistricoDeCorridas() error {
	return godog.ErrPending
}

func estouNaTelaDeMenuPrincipal() error {
	return godog.ErrPending
}

func estouVisualizandoOTempoEstimadoDeChegadaDeMinutos(arg1 int) error {
	return godog.ErrPending
}

func euAcessoMeuPerfilNoNavegadorWeb() error {
	return godog.ErrPending
}

func euAcessoOHistricoDeAlteraes() error {
	return godog.ErrPending
}

func euAcessoOLinkDeCancelamentoNoEmailRecebido() error {
	return godog.ErrPending
}

func euAlteroASenhaPara(arg1 string) error {
	return godog.ErrPending
}

func euCanceloASolicitao() error {
	return godog.ErrPending
}

func euClico(arg1 string) error {
	return godog.ErrPending
}

func euClicoEm(arg1 string) error {
	return godog.ErrPending
}

func euConfirmoADesativao() error {
	return godog.ErrPending
}

func euConfirmoAReativao() error {
	return godog.ErrPending
}

func euConfirmoDigitando(arg1 string) error {
	return godog.ErrPending
}

func euDesativoAOpo(arg1 string) error {
	return godog.ErrPending
}

func euEstouNaMinhaReaDoMotorista() error {
	return godog.ErrPending
}

func euEstouNaPginaDeUploadDeDocumentos() error {
	return godog.ErrPending
}

func euFaoUploadDaFotoDaCNHEmFormatoJPGComMB(arg1 int) error {
	return godog.ErrPending
}

func euFaoUploadDaFotoDoCRLVEmFormatoPNGComMB(arg1, arg2 int) error {
	return godog.ErrPending
}

func euFaoUploadDeUmaSelfieComCNHEmFormatoJPGComMB(arg1 int) error {
	return godog.ErrPending
}

func euFicoInativoPorMaisDeMinutos(arg1 int) error {
	return godog.ErrPending
}

func euNoReceboNovasSolicitaesDeCorrida() error {
	return godog.ErrPending
}

func euPercoAConexoComAInternetPorMinutos(arg1 int) error {
	return godog.ErrPending
}

func euPermaneoLogado() error {
	return godog.ErrPending
}

func euPossoFazerLoginNoSistema() error {
	return godog.ErrPending
}

func euPossoFazerUploadDaNovaCNH() error {
	return godog.ErrPending
}

func euPossoReenviarOsDocumentosCorrigidos() error {
	return godog.ErrPending
}

func euPreenchoAConfirmaoDeSenha(arg1 string) error {
	return godog.ErrPending
}

func euPreenchoASenha(arg1 string) error {
	return godog.ErrPending
}

func euPreenchoOMotivo(arg1 string) error {
	return godog.ErrPending
}

func euRealizoLoginComEmailESenha(arg1, arg2 string) error {
	return godog.ErrPending
}

func euReceboUmEmail(arg1 string) error {
	return godog.ErrPending
}

func euReceboUmEmailComInstruesParaConfirmarAExcluso() error {
	return godog.ErrPending
}

func euReceboUmEmailConfirmandoORecebimentoDosDocumentos() error {
	return godog.ErrPending
}

func euReceboUmEmailDetalhandoOsProblemasEncontrados() error {
	return godog.ErrPending
}

func euRecuperoAConexoComAInternet() error {
	return godog.ErrPending
}

func euSelecionoOMotivo(arg1 string) error {
	return godog.ErrPending
}

func euSouDirecionadoParaAPginaDeUploadDeDocumentos() error {
	return godog.ErrPending
}

func euSouRedirecionadoParaAPginaDeLogin() error {
	return godog.ErrPending
}

func euSouRedirecionadoParaOMeuPrprioPerfil() error {
	return godog.ErrPending
}

func euSubmetoOFormulrio() error {
	return godog.ErrPending
}

func euTentoAcessarDiretamenteAURLDoPerfilDoMotorista(arg1 string) error {
	return godog.ErrPending
}

func euTentoAcessarODashboard() error {
	return godog.ErrPending
}

func euTentoAcessarUmaFuncionalidadeDoSistema() error {
	return godog.ErrPending
}

func euTentoFazerUploadDeUmArquivoDeMB(arg1 int) error {
	return godog.ErrPending
}

func euTentoFazerUploadDeUmArquivoEmFormatoTXT() error {
	return godog.ErrPending
}

func euVejo(arg1 string) error {
	return godog.ErrPending
}

func euVejoADatahoraDaLtimaSincronizao() error {
	return godog.ErrPending
}

func euVejoAMensagem(arg1 string) error {
	return godog.ErrPending
}

func euVejoAMensagemDeErro(arg1 string) error {
	return godog.ErrPending
}

func euVejoAOpo(arg1 string) error {
	return godog.ErrPending
}

func euVejoApenasAsAlteraesDosLtimosDias(arg1 int) error {
	return godog.ErrPending
}

func euVejoOBoto(arg1 string) error {
	return godog.ErrPending
}

func euVejoOIndicadorDeForaDaSenhaComo(arg1 string) error {
	return godog.ErrPending
}

func euVejoOStatus(arg1 string) error {
	return godog.ErrPending
}

func euVejoOTelefoneAtualizadoTambmNoNavegador() error {
	return godog.ErrPending
}

func euVejoOTutorialInicialDoMotorista() error {
	return godog.ErrPending
}

func euVejoQuaisDocumentosEspecficosForamRejeitados() error {
	return godog.ErrPending
}

func euVejoUmAlerta(arg1 string) error {
	return godog.ErrPending
}

func euVejoUmAvisoDetalhadoSobreAsConsequnciasDaExcluso() error {
	return godog.ErrPending
}

func euVejoUmFormulrioParaInformarOMotivoDaExcluso() error {
	return godog.ErrPending
}

func euVejoUmaListaCronolgicaDasModificaes(arg1 *godog.Table) error {
	return godog.ErrPending
}

func euVejoUmaNotificao(arg1 string) error {
	return godog.ErrPending
}

func euVoltoAReceberSolicitaesDeCorrida() error {
	return godog.ErrPending
}

func existeUmaCorridaQueJFoiAvaliadaAnteriormente() error {
	return godog.ErrPending
}

func faoUploadDeUmaNovaFotoEmFormatoJPGComMB(arg1 int) error {
	return godog.ErrPending
}

func faoUploadDosDocumentosObrigatrios(arg1 *godog.Table) error {
	return godog.ErrPending
}

func ficoInativoPorMaisDeMinutos(arg1 int) error {
	return godog.ErrPending
}

func foiAoLocalDeEmbarqueEChegouAoDestino(arg1, arg2 string) error {
	return godog.ErrPending
}

func foiAoLocalDeEmbarqueEIniciouOTrajetoParaODestino(arg1, arg2 string) error {
	return godog.ErrPending
}

func meuStatusInicial(arg1 string) error {
	return godog.ErrPending
}

func meuStatusMudaPara(arg1 string) error {
	return godog.ErrPending
}

func meusDadosSoAtualizadosNoSistema() error {
	return godog.ErrPending
}

func meusDadosSoSincronizadosAutomaticamente() error {
	return godog.ErrPending
}

func minhaContaEstComStatus(arg1 string) error {
	return godog.ErrPending
}

func minhaContaFicaComStatus(arg1 string) error {
	return godog.ErrPending
}

func minhaContaVoltaAoStatus(arg1 string) error {
	return godog.ErrPending
}

func motoristaDemoraMaisDoQueMinutosParaChegarAoLocal(arg1 string) error {
	return godog.ErrPending
}

func nenhumEmailDeRecuperaoEnviado() error {
	return godog.ErrPending
}

func nenhumEmailEnviado() error {
	return godog.ErrPending
}

func noConsigoConcluirOCadastro() error {
	return godog.ErrPending
}

func noConsigoFazerLoginMesmoComCredenciaisCorretas() error {
	return godog.ErrPending
}

func noDeveReceberNotificaesDeNovasCorridas() error {
	return godog.ErrPending
}

func noDeveSerRegistradaNenhumaCorridaNoHistrico() error {
	return godog.ErrPending
}

func noInterageComANotificaoEmAtSegundos(arg1 int) error {
	return godog.ErrPending
}

func noSouAutenticado() error {
	return godog.ErrPending
}

func noVejoOBoto(arg1 string) error {
	return godog.ErrPending
}

func oEmailContmUmLinkVlidoPorHora(arg1 int) error {
	return godog.ErrPending
}

func oLinkDeRecuperaoInvalidado() error {
	return godog.ErrPending
}

func oMotoristaDeNomeAceitouUmaCorrida(arg1 string) error {
	return godog.ErrPending
}

func oMotoristaDeNomeJooEstLogado() error {
	return godog.ErrPending
}

func oMotoristaEstAMinutosDoLocalDeEmbarque(arg1, arg2 string) error {
	return godog.ErrPending
}

func oMotoristaRecebeAMensagemDeConfirmaoNaTela(arg1 string) error {
	return godog.ErrPending
}

func oMotoristaRecebeAMensagemDeErroNaTela(arg1 string) error {
	return godog.ErrPending
}

func oMotoristaSelecionaAOpo(arg1 string) error {
	return godog.ErrPending
}

func oMotoristaTentaCancelarACorridaSelecionandoAOpo(arg1 string) error {
	return godog.ErrPending
}

func oPopupExibeOsBotesE(arg1, arg2 string) error {
	return godog.ErrPending
}

func oPopupFechado() error {
	return godog.ErrPending
}

func oServioDeValidaoDeDocumentosEstDisponvel() error {
	return godog.ErrPending
}

func oSistemaCancelaAutomaticamenteACorrida() error {
	return godog.ErrPending
}

func oSistemaCancelaOProgressoDaCorrida() error {
	return godog.ErrPending
}

func oSistemaDeEnvioDeEmailsEstFuncionando() error {
	return godog.ErrPending
}

func oSistemaDeVerificaoAutomticaEstFuncionando() error {
	return godog.ErrPending
}

func oSistemaExibeUmaEntradaNoHistricoComOStatus(arg1 string) error {
	return godog.ErrPending
}

func oSistemaExibeUmaNotificaoNaTelaComAMensagem(arg1 *godog.DocString) error {
	return godog.ErrPending
}

func oSistemaNoPermiteOCancelamento() error {
	return godog.ErrPending
}

func oSistemaOuAnalistaVerificaMeusDocumentos() error {
	return godog.ErrPending
}

func oSistemaRegistraATentativaDeAcessoNoAutorizado() error {
	return godog.ErrPending
}

func oSistemaValidaMeusDocumentos() error {
	return godog.ErrPending
}

func oStatusDeveEstarDisponvel() error {
	return godog.ErrPending
}

func oStatusDeveEstarOcupado() error {
	return godog.ErrPending
}

func oStatusDevePermanecerComo(arg1 string) error {
	return godog.ErrPending
}

func oStatusDeveSerAlteradoPara(arg1 string) error {
	return godog.ErrPending
}

func oUploadCancelado() error {
	return godog.ErrPending
}

func oUploadNoConcludo() error {
	return godog.ErrPending
}

func oValorTotalEADistnciaPercorridaSoExibidos() error {
	return godog.ErrPending
}

func osCamposDeSenhaSoDestacados() error {
	return godog.ErrPending
}

func osDadosDaAvaliaoAnteriorSoExibidos() error {
	return godog.ErrPending
}

func possoFazerLoginComANovaSenha(arg1 string) error {
	return godog.ErrPending
}

func preencho(arg1 *godog.Table) error {
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

func preenchoOsCamposDeSenha(arg1 *godog.Table) error {
	return godog.ErrPending
}

func queAceiteiUmaCorridaDisponvel() error {
	return godog.ErrPending
}

func queCompleteiOCadastroBsicoComSucesso() error {
	return godog.ErrPending
}

func queEstouAutenticadoComoMotorista(arg1 *godog.Table) error {
	return godog.ErrPending
}

func queEstouLogadoComoMotorista() error {
	return godog.ErrPending
}

func queEstouLogadoComoOMotoristaJoo() error {
	return godog.ErrPending
}

func queEstouLogadoNoSistemaComoMotorista() error {
	return godog.ErrPending
}

func queEstouLogadoNoSistemaComoOMotoristaJoo() error {
	return godog.ErrPending
}

func queEstouNaPgina(arg1 string) error {
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

func queEstouNaPginaDeUploadDeDocumentos() error {
	return godog.ErrPending
}

func queEstouNaPginaDoMeuPerfil() error {
	return godog.ErrPending
}

func queEuAlteroMeuTelefoneNoDispositivoMobile() error {
	return godog.ErrPending
}

func queEuCompleteiOCadastroBsicoComSucesso() error {
	return godog.ErrPending
}

func queEuEnvieiDocumentosComProblemasDeQualidade() error {
	return godog.ErrPending
}

func queEuEnvieiTodosOsDocumentosObrigatrios() error {
	return godog.ErrPending
}

func queEuEstouLogadoComoMotorista(arg1 string) error {
	return godog.ErrPending
}

func queEuEstouNaPginaDeCadastroDeMotorista() error {
	return godog.ErrPending
}

func queEuEstouNaPginaDeUploadDeDocumentos() error {
	return godog.ErrPending
}

func queEuEstouNaPginaDoMeuPerfil() error {
	return godog.ErrPending
}

func queEuSoliciteiExclusoDaContaHDias(arg1 int) error {
	return godog.ErrPending
}

func queExisteOutroMotoristaComID(arg1 string) error {
	return godog.ErrPending
}

func queExisteUmMotoristaCadastrado(arg1 *godog.Table) error {
	return godog.ErrPending
}

func queExisteUmMotoristaCadastradoComCnh(arg1 string) error {
	return godog.ErrPending
}

func queExisteUmMotoristaCadastradoComCpf(arg1 string) error {
	return godog.ErrPending
}

func queExisteUmMotoristaCadastradoComEmail(arg1 string) error {
	return godog.ErrPending
}

func queExisteUmMotoristaCadastradoComOsDados(arg1 *godog.Table) error {
	return godog.ErrPending
}

func queExisteUmMotoristaComStatus(arg1 string, arg2 *godog.Table) error {
	return godog.ErrPending
}

func queFinalizeiUmaCorridaApsChegarAoDestino() error {
	return godog.ErrPending
}

func queMinhaCNHEstPrximaDoVencimentoDias(arg1 int) error {
	return godog.ErrPending
}

func queMinhaContaEstComStatus(arg1 string) error {
	return godog.ErrPending
}

func queOEmailNoEstBloqueado(arg1 string) error {
	return godog.ErrPending
}

func queOMotoristaDeNomeJooEstLogado() error {
	return godog.ErrPending
}

func queOSistemaDeCadastroEstFuncionando() error {
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

func realizoLoginComEmailESenha(arg1, arg2 string) error {
	return godog.ErrPending
}

func realizoOCadastroComDadosVlidos(arg1 *godog.Table) error {
	return godog.ErrPending
}

func realizoTentativasDeLoginComSenhaIncorretaPara(arg1 int, arg2 string) error {
	return godog.ErrPending
}

func recebeUmaNotificaoDeCorrida() error {
	return godog.ErrPending
}

func recebeUmaNotificaoDeUmaNovaCorridaPrxima() error {
	return godog.ErrPending
}

func recebiUmaNotificaoDeUmaNovaCorridaPrxima() error {
	return godog.ErrPending
}

func receboUmEmailComInstruesParaConfirmarAExcluso() error {
	return godog.ErrPending
}

func receboUmEmailDeConfirmao() error {
	return godog.ErrPending
}

func recusaACorrida() error {
	return godog.ErrPending
}

func salvoAAlterao() error {
	return godog.ErrPending
}

func salvoANovaFoto() error {
	return godog.ErrPending
}

func salvoAsAlteraes() error {
	return godog.ErrPending
}

func selecionaAOpo(arg1 string) error {
	return godog.ErrPending
}

func selecionaNaInterfaceDeDesencentivoACancelamento(arg1 string) error {
	return godog.ErrPending
}

func selecioneiUmaCorridaAindaNoAvaliada() error {
	return godog.ErrPending
}

func selecionoOMotivo(arg1 string) error {
	return godog.ErrPending
}

func solicitoDesativarMinhaContaTemporariamente() error {
	return godog.ErrPending
}

func solicitoExcluirMinhaContaPermanentemente() error {
	return godog.ErrPending
}

func solicitoReativarMinhaConta() error {
	return godog.ErrPending
}

func solicitoRecuperaoNovamenteComOMesmoEmail() error {
	return godog.ErrPending
}

func solicitoRecuperaoPara(arg1 string) error {
	return godog.ErrPending
}

func souAutenticadoComSucesso() error {
	return godog.ErrPending
}

func souRedirecionadoParaAPginaDeLogin() error {
	return godog.ErrPending
}

func souRedirecionadoParaATelaDeAvaliaoDeCorrida() error {
	return godog.ErrPending
}

func submetoOFormulrio() error {
	return godog.ErrPending
}

func submetoOFormulrioComEmail(arg1 string) error {
	return godog.ErrPending
}

func tenhoPeloMenosUmaCorridaRealizadaQueAindaNoFoiAvaliada() error {
	return godog.ErrPending
}

func tenhoUmaCorridaAnteriorAindaNoAvaliada() error {
	return godog.ErrPending
}

func tentoAcessarOMesmoLinkNovamente() error {
	return godog.ErrPending
}

func tentoAcessarUmaFuncionalidadeDoSistema() error {
	return godog.ErrPending
}

func tentoFazerUploadDeArquivo(arg1 string) error {
	return godog.ErrPending
}

func tentoFazerUploadDeArquivoCom(arg1 string) error {
	return godog.ErrPending
}

func tentoRealizarCadastroComCNHVlidaAt(arg1 string) error {
	return godog.ErrPending
}

func tentoRealizarCadastroComCnh(arg1 string) error {
	return godog.ErrPending
}

func tentoRealizarCadastroComCpf(arg1 string) error {
	return godog.ErrPending
}

func tentoRealizarCadastroComDataDeNascimento(arg1 string) error {
	return godog.ErrPending
}

func tentoRealizarCadastroComEmail(arg1 string) error {
	return godog.ErrPending
}

func tentoRealizarCadastroSemPreencherOCampo(arg1 string) error {
	return godog.ErrPending
}

func tentoSalvarAAlterao() error {
	return godog.ErrPending
}

func tentoSalvarAsAlteraes() error {
	return godog.ErrPending
}

func todosOsDocumentosSoAprovadosAutomaticamente() error {
	return godog.ErrPending
}

func umEmailDeRecuperaoEnviadoPara(arg1 string) error {
	return godog.ErrPending
}

func umaMensagemDeErroApareceNaTela(arg1 string) error {
	return godog.ErrPending
}

func umaNotificaoEnviadaParaOEmailDoMotoristaSobreTentativaDeAcesso() error {
	return godog.ErrPending
}

func umaNovaCorridaSolicitadaPorUmPassageiroPrximo() error {
	return godog.ErrPending
}

func vejoACorridaComDataEHora() error {
	return godog.ErrPending
}

func vejoAMensagem(arg1 string) error {
	return godog.ErrPending
}

func vejoAMensagemDeErro(arg1 string) error {
	return godog.ErrPending
}

func vejoANotificao(arg1 string) error {
	return godog.ErrPending
}

func vejoAOpoDeSolicitarUmNovoLink() error {
	return godog.ErrPending
}

func vejoComoMotoristaLogado(arg1 string) error {
	return godog.ErrPending
}

func vejoEstatsticasSobreMinhasCorridas() error {
	return godog.ErrPending
}

func vejoMeusDadosPessoaisCadastrados() error {
	return godog.ErrPending
}

func vejoMinhaFotoDePerfil() error {
	return godog.ErrPending
}

func vejoODashboardDoMotorista() error {
	return godog.ErrPending
}

func vejoOTempoEstimadoDeChegadaAtOLocalDeEmbarque() error {
	return godog.ErrPending
}

func vejoOsCamposDeNotaEComentrioDisponveisParaPreenchimento() error {
	return godog.ErrPending
}

func vejoQueDevoAguardarAntesDeSolicitarNovamente() error {
	return godog.ErrPending
}

func vejoUmPopupEscrito(arg1 string) error {
	return godog.ErrPending
}

func vejoUmaListaContendoAsLtimasCorridasRealizadas() error {
	return godog.ErrPending
}

func vejoUmaMensagem(arg1 string) error {
	return godog.ErrPending
}

func visualizoOsDetalhesDessaCorrida() error {
	return godog.ErrPending
}
