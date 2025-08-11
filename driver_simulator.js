const axios = require('axios');

// --- Funções de utilidade para parsing de argumentos ---
function getArg(argName, defaultValue) {
  const arg = process.argv.find(a => a.startsWith(`--${argName}=`));
  if (arg) {
    return arg.split('=')[1];
  }
  return defaultValue;
}

// --- Configuração ---
const API_BASE_URL = 'http://localhost:3000/api';
const CORRIDA_ID = getArg('corrida-id', null);
const MOTORISTA_ID = getArg('motorista-id', 456);
const SCENARIO = getArg('cenario', 'ON_TIME');
// --------------------

const UPDATE_INTERVAL_MS = 2000;
const TOTAL_STEPS = 15;

if (!CORRIDA_ID) {
  console.error('Erro: O parâmetro --corrida-id é obrigatório.');
  console.log('Uso: node driver_simulator.js --corrida-id=<ID> [--cenario=<NOME>]');
  process.exit(1);
}

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
    console.log('CYPRESS_TASK_RIDE_ACCEPTED'); // Signal for Cypress task
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
  let currentLng = lngDestino;

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

    switch (SCENARIO.toUpperCase()) {
        case 'EARLY':
            return tempoEstimadoMs - 30000;
        case 'LATE':
            return tempoEstimadoMs + 30000;
        case 'AUTO_CANCEL':
            return tempoEstimadoMs * 100; 
        case 'ON_TIME':
        default:
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
    setTimeout(() => {
        iniciarViagem(rideDetails.Origem, rideDetails.Destino, rideDetails.TempoEstimado);
    }, 2000);
  }
}

iniciarSimulacao();