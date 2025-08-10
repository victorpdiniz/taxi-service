/**
 * Script de Simulação de Motorista para a API de Serviço de Táxi
 *
 * Este script simula um motorista aceitando uma corrida, dirigindo até o destino
 * e finalizando a corrida.
 * É necessário ter o `axios` instalado (`npm install axios`).
 *
 * Como usar:
 * 1. Inicie o backend e o frontend.
 * 2. Use a aplicação frontend para criar uma nova corrida. Anote o ID da corrida.
 * 3. Configure as variáveis `CORRIDA_ID` e `MOTORISTA_ID` abaixo.
 * 4. Execute o script com `node driver_simulator.js`.
 */

const axios = require('axios');

// --- Configuração ---
const API_BASE_URL = 'http://localhost:3000/api';
const CORRIDA_ID = 1; // <-- IMPORTANTE: Altere para o ID da corrida que você quer simular
const MOTORISTA_ID = 123; // ID do motorista que está simulando
const UPDATE_INTERVAL_MS = 2000; // Intervalo de atualização da posição em milissegundos (2 segundos)
const TOTAL_STEPS = 20; // Número de "passos" para ir da origem ao destino
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
    console.log('Detalhes da corrida obtidos com sucesso.');
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
function iniciarViagem(origem, destino) {
  console.log(`Iniciando viagem da origem ${origem} para o destino ${destino}.`);

  const [latOrigem, lngOrigem] = origem.split(',').map(Number);
  const [latDestino, lngDestino] = destino.split(',').map(Number);

  let currentLat = latOrigem;
  let currentLng = lngOrigem;

  const latStep = (latDestino - latOrigem) / TOTAL_STEPS;
  const lngStep = (lngDestino - lngOrigem) / TOTAL_STEPS;

  let stepCount = 0;

  const interval = setInterval(async () => {
    if (stepCount >= TOTAL_STEPS) {
      clearInterval(interval);
      console.log('Motorista chegou ao destino!');
      finalizarCorrida();
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
      console.log(`[Passo ${stepCount}/${TOTAL_STEPS}] Posição atualizada: Lat ${currentLat.toFixed(6)}, Lng ${currentLng.toFixed(6)}`);
    } catch (error) {
      console.error('Erro ao atualizar a posição:', error.response ? error.response.data : error.message);
    }
  }, UPDATE_INTERVAL_MS);
}

/**
 * Chama o endpoint para finalizar a corrida.
 */
async function finalizarCorrida() {
  try {
    console.log(`Finalizando a corrida ${CORRIDA_ID}...`);
    await api.post(`/corrida/${CORRIDA_ID}/finalizar`);
    console.log('Corrida finalizada com sucesso no backend!');
  } catch (error) {
    console.error('Erro ao finalizar a corrida:', error.response ? error.response.data : error.message);
  }
}

/**
 * Função principal que orquestra a simulação.
 */
async function iniciarSimulacao() {
  console.log('--- Iniciando Simulação de Motorista ---');
  
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
    iniciarViagem(rideDetails.Origem, rideDetails.Destino);
  }
}

iniciarSimulacao();