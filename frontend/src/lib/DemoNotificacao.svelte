<script>
  import NotificacaoPopUp from './NotificacaoPopUp.svelte';
  import CorridaCard from './CorridaCard.svelte';
  import axios from 'axios';

  const corridasMockadas = [
    { ID: 1, PassageiroNome: "Maria Silva", Origem: "Shopping Mueller", Destino: "Aeroporto", Valor: 45.50, DistanciaKm: 18.2, TempoEstimado: "25", CorridaID: 1001 },
    { ID: 2, PassageiroNome: "João Santos", Origem: "Centro", Destino: "Batel", Valor: 22.75, DistanciaKm: 8.3, TempoEstimado: "15", CorridaID: 1002 },
    { ID: 3, PassageiroNome: "Ana Costa", Origem: "UFPR", Destino: "Hospital", Valor: 18.90, DistanciaKm: 6.1, TempoEstimado: "12", CorridaID: 1003 },
    { ID: 4, PassageiroNome: "Carlos Oliveira", Origem: "Rodoviária", Destino: "Água Verde", Valor: 15.20, DistanciaKm: 4.8, TempoEstimado: "10", CorridaID: 1004 }
  ];

  let currentNotificacao = null;
  let aceitas = [];
  let recusadas = [];
  let expiradas = [];

  const api = axios.create({
    baseURL: 'http://localhost:3000'
  });

  async function solicitarCorrida(event) {
    const corrida = event.detail;
    try {
      const payload = {
        corrida_id: corrida.CorridaID,
        motorista_id: 1,
        passageiro_nome: corrida.PassageiroNome,
        origem: corrida.Origem,
        destino: corrida.Destino,
        valor: corrida.Valor,
        distancia_km: corrida.DistanciaKm,
        tempo_estimado: corrida.TempoEstimado
      };

      const response = await api.post('/api/notificacoes', payload);
      currentNotificacao = { ...corrida, timestamp: new Date().toISOString() };
      
    } catch (error) {
      console.error('Erro:', error);
      alert(`Erro ao criar notificação: ${error.response?.data?.message || error.message}`);
    }
  }

  function handleNotificationAccepted() {
    aceitas = [...aceitas, currentNotificacao];
    currentNotificacao = null;
  }

  function handleNotificationRefused() {
    recusadas = [...recusadas, currentNotificacao];
    currentNotificacao = null;
  }

  function handleNotificationExpired() {
    expiradas = [...expiradas, currentNotificacao];
    currentNotificacao = null;
  }

  function clearHistory() {
    aceitas = [];
    recusadas = [];
    expiradas = [];
  }
</script>

<div class="container">
  <div class="stats">
    <div class="stat aceitas">Aceitas: {aceitas.length}</div>
    <div class="stat recusadas">Recusadas: {recusadas.length}</div>
    <div class="stat expiradas">Expiradas: {expiradas.length}</div>
  </div>

  <div class="corridas">
    {#each corridasMockadas as corrida}
      <CorridaCard {corrida} on:solicitar={solicitarCorrida} />
    {/each}
  </div>

  {#if aceitas.length > 0 || recusadas.length > 0 || expiradas.length > 0}
    <div class="historico">
      <div class="historico-header">
        <h2>📋 Histórico</h2>
        <button on:click={clearHistory}>🗑️ Limpar</button>
      </div>
      
      {#each [...aceitas.map(c => ({...c, status: 'aceita'})), ...recusadas.map(c => ({...c, status: 'recusada'})), ...expiradas.map(c => ({...c, status: 'expirada'}))] as item}
        <div class="historico-item {item.status}">
          {item.status === 'aceita' ? '✅' : item.status === 'recusada' ? '❌' : '⏰'}
          <strong>{item.PassageiroNome}</strong> - R$ {item.Valor.toFixed(2)}
          <small>{new Date(item.timestamp).toLocaleTimeString()}</small>
        </div>
      {/each}
    </div>
  {/if}
</div>

<NotificacaoPopUp 
  motoristaId={1}
  notificacao={currentNotificacao}
  on:accepted={handleNotificationAccepted}
  on:refused={handleNotificationRefused}
  on:expired={handleNotificationExpired}
/>

<style>
  .container {
    padding: 2rem;
    max-width: 800px;
    margin: 0 auto;
    font-family: Arial, sans-serif;
  }

  .stats {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1rem;
    margin-bottom: 2rem;
  }

  .stat {
    background: white;
    padding: 1rem;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    text-align: center;
    font-weight: bold;
  }

  .stat.aceitas { border-left: 4px solid #10b981; }
  .stat.recusadas { border-left: 4px solid #ef4444; }
  .stat.expiradas { border-left: 4px solid #f59e0b; }

  .corridas {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 1rem;
    margin-bottom: 2rem;
  }

  .historico {
    background: white;
    border-radius: 8px;
    padding: 1rem;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  }

  .historico-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
  }

  .historico-header h2 {
    margin: 0;
  }

  .historico-header button {
    background: #ef4444;
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 4px;
    cursor: pointer;
  }

  .historico-item {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.5rem;
    margin-bottom: 0.5rem;
    border-radius: 4px;
  }

  .historico-item.aceita { background: #ecfdf5; }
  .historico-item.recusada { background: #fef2f2; }
  .historico-item.expirada { background: #fffbeb; }

  .historico-item small {
    margin-left: auto;
    color: #888;
  }

  @media (max-width: 600px) {
    .container { padding: 1rem; }
    .corridas { grid-template-columns: 1fr; }
    .stats { grid-template-columns: 1fr; }
  }
</style>