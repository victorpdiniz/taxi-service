/**
 * Script de Simulação de Motorista para a API de Serviço de Táxi
 *
 * Este script simula um motorista executando diferentes cenários de uma corrida.
 * É necessário ter o `axios` instalado (`npm install axios`).
 *
 * Como usar:
 * 1. Inicie o backend e o frontend.
 * 2. Use a aplicação frontend para criar uma nova corrida. Anote o ID da corrida.
 * 3. Configure as variáveis `CORRIDA_ID`, `MOTORISTA_ID` e `SCENARIO` abaixo.
 * 4. Execute o script com `node driver_simulator.js`.
 */

const axios = require('axios');

// --- Configuração ---
const API_BASE_URL = 'http://localhost:3000/api';
const CORRIDA_ID = 1;     // <-- Altere para o ID da corrida que você quer simular
const MOTORISTA_ID = 456;   // ID do motorista que está simulando

// --- SELECIONE O CENÁRIO ---
// Cenários disponíveis: 'ON_TIME', 'EARLY', 'LATE', 'AUTO_CANCEL'
const SCENARIO = 'LATE'; 
// ---------------------------

const UPDATE_INTERVAL_MS = 2000; // Intervalo de atualização da posição (2 segundos)
const TOTAL_STEPS = 15; // Número de "passos" para ir da origem ao destino
// ---------------------

const api = axios.create({
  baseURL: API_BASE_URL,
});

/**
 * Busca os detalhes da corrida na API.
 */
async function getRideDetails() {
  try {
    console.log(`Buscando detalhes da corrida ${CORRIDA_ID}...`);
    const response = await api.get(`/corrida/${CORRIDA_ID}`);
    return response.data;
  } catch (error) {
    console.error('Erro ao buscar detalhes da corrida:', error.response ? error.response.data : error.message);
    return null;
  }
}

/**
 * Simula o motorista aceitando a corrida.
 */
async function aceitarCorrida() {
  try {
    console.log(`Motorista ${MOTORISTA_ID} tentando aceitar a corrida ${CORRIDA_ID}...`);
    await api.put(`/corrida/${CORRIDA_ID}/aceitar`, { motoristaId: MOTORISTA_ID });
    console.log(`Corrida ${CORRIDA_ID} aceita com sucesso!`);
    return true;
  } catch (error) {
    console.error('Erro ao aceitar a corrida:', error.response ? error.response.data : error.message);
    return false;
  }
}

/**
 * Inicia o envio periódico de atualizações de posição, movendo-se em direção ao destino.
 */
function iniciarViagem(origem, destino, tempoEstimadoMin) {
  console.log(`Iniciando viagem da origem ${origem} para o destino ${destino}.`);
  console.log(`Tempo estimado pelo backend: ${tempoEstimadoMin} minuto(s).`);

  const [latOrigem, lngOrigem] = origem.split(',').map(Number);
  const [latDestino, lngDestino] = destino.split(',').map(Number);

  let currentLat = latOrigem;
  let currentLng = lngOrigem;

  const latStep = (latDestino - latOrigem) / TOTAL_STEPS;
  const lngStep = (lngDestino - lngOrigem) / TOTAL_STEPS;

  let stepCount = 0;

  const totalViagemMs = calcularDuracaoViagem(tempoEstimadoMin);
  const updateInterval = totalViagemMs / TOTAL_STEPS;
  
  console.log(`Cenário selecionado: ${SCENARIO}. A viagem simulada durará ${totalViagemMs / 1000} segundos.`);

  const interval = setInterval(async () => {
    if (stepCount >= TOTAL_STEPS) {
      clearInterval(interval);
      console.log('Motorista chegou ao destino!');
      // No cenário de auto-cancelamento, não finalizamos a corrida, esperamos o backend fazer isso.
      if (SCENARIO !== 'AUTO_CANCEL') {
        finalizarCorrida();
      }
      return;
    }

    currentLat += latStep;
    currentLng += lngStep;
    stepCount++;

    try {
      await api.put(`/corrida/${CORRIDA_ID}/posicao`, {
        lat: currentLat,
        lng: currentLng,
      });
      console.log(`[Passo ${stepCount}/${TOTAL_STEPS}] Posição atualizada.`);
    } catch (error) {
      console.error('Erro ao atualizar a posição:', error.response ? error.response.data : error.message);
    }
  }, updateInterval);
}

function calcularDuracaoViagem(tempoEstimadoMin) {
    const tempoEstimadoMs = tempoEstimadoMin * 60 * 1000;

    switch (SCENARIO) {
        case 'EARLY':
            // 30 segundos a menos que o estimado
            return tempoEstimadoMs - 30000;
        case 'LATE':
            // 30 segundos a mais que o estimado
            return tempoEstimadoMs + 30000;
        case 'AUTO_CANCEL':
            // Simula uma viagem muito longa que nunca termina
            // O backend deverá cancelar após 15 minutos do tempo estimado
            return tempoEstimadoMs * 100; 
        case 'ON_TIME':
        default:
            // Exatamente no tempo
            return tempoEstimadoMs;
    }
}

/**
 * Chama o endpoint para finalizar a corrida.
 */
async function finalizarCorrida() {
  try {
    console.log(`Finalizando a corrida ${CORRIDA_ID}...`);
    await api.post(`/corrida/${CORRIDA_ID}/finalizar`);
    console.log('Corrida finalizada com sucesso no backend! Verifique o status final.');
  } catch (error) {
    console.error('Erro ao finalizar a corrida:', error.response ? error.response.data : error.message);
  }
}

/**
 * Função principal que orquestra a simulação.
 */
async function iniciarSimulacao() {
  console.log(`--- Iniciando Simulação de Motorista para o cenário: ${SCENARIO} ---`);
  
  const rideDetails = await getRideDetails();
  if (!rideDetails) {
    console.log('Não foi possível obter os detalhes da corrida. Abortando simulação.');
    return;
  }

  if (!rideDetails.Origem || !rideDetails.Destino) {
      console.log('A corrida não tem uma origem ou destino definidos. Abortando simulação.');
      return;
  }

  const aceita = await aceitarCorrida();
  if (aceita) {
    // Espera 2 segundos para dar tempo de ver o status "Motorista encontrado"
    setTimeout(() => {
        iniciarViagem(rideDetails.Origem, rideDetails.Destino, rideDetails.TempoEstimado);
    }, 2000);
  }
}

iniciarSimulacao();
