<script>
  import { onMount, onDestroy } from 'svelte';
  import { navigate } from 'svelte-routing';
  import NotificacaoPopup from './NotificacaoPopup.svelte';
  import axios from 'axios';

  let motoristaId = 1; // ID do motorista logado
  let currentNotificacao = null;
  let pollingInterval = null;
  let historicoNotificacoes = [];
  let isOnline = true;

  const api = axios.create({
    baseURL: 'http://localhost:3000/api',
  });

  onMount(() => {
    if (isOnline) {
      startPolling();
    }
    loadHistorico();
  });

  onDestroy(() => {
    stopPolling();
  });

  function startPolling() {
    stopPolling(); // Garante que não há polling duplicado
    
    pollingInterval = setInterval(async () => {
      if (isOnline && !currentNotificacao) {
        await checkForNewNotifications();
      }
    }, 2000); // Verifica a cada 2 segundos
  }

  function stopPolling() {
    if (pollingInterval) {
      clearInterval(pollingInterval);
      pollingInterval = null;
    }
  }

  async function checkForNewNotifications() {
    try {
      const response = await api.get(`/notificacoes/motorista/${motoristaId}/pending`);
      const { pending_notificacoes } = response.data;
      
      if (pending_notificacoes && pending_notificacoes.length > 0) {
        // Pega a primeira notificação pendente
        currentNotificacao = pending_notificacoes[0];
      }
    } catch (error) {
      console.error('Erro ao verificar notificações:', error);
    }
  }

  async function loadHistorico() {
    try {
      const response = await api.get(`/notificacoes/motorista/${motoristaId}/historico`);
      historicoNotificacoes = response.data.historico || [];
    } catch (error) {
      console.error('Erro ao carregar histórico:', error);
    }
  }

  function handleNotificationAccepted(event) {
    const { corridaId } = event.detail;
    alert(`Corrida aceita! Redirecionando para a corrida ${corridaId}`);
    
    // Redirecionar para a página da corrida
    navigate(`/corrida/${corridaId}`);
    
    currentNotificacao = null;
    loadHistorico();
  }

  function handleNotificationRefused(event) {
    const { notificacaoId } = event.detail;
    console.log('Notificação recusada:', notificacaoId);
    
    currentNotificacao = null;
    loadHistorico();
  }

  function handleNotificationExpired(event) {
    const { notificacaoId } = event.detail;
    console.log('Notificação expirada:', notificacaoId);
    
    currentNotificacao = null;
    loadHistorico();
  }

  function toggleOnlineStatus() {
    isOnline = !isOnline;
    
    if (isOnline) {
      startPolling();
    } else {
      stopPolling();
      currentNotificacao = null;
    }
  }

  async function forceExpireNotifications() {
    try {
      await api.post('/notificacoes/expire');
      console.log('Notificações expiradas processadas');
    } catch (error) {
      console.error('Erro ao expirar notificações:', error);
    }
  }

  function getStatusColor(status) {
    switch (status) {
      case 'aceita': return '#10b981';
      case 'recusada': return '#ef4444';
      case 'expirada': return '#f59e0b';
      default: return '#6b7280';
    }
  }

  function formatTimestamp(timestamp) {
    return new Date(timestamp).toLocaleString('pt-BR');
  }
</script>

<style>
  .painel-container {
    padding: 2rem;
    max-width: 1200px;
    margin: 0 auto;
  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
    background: white;
    padding: 1.5rem;
    border-radius: 12px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }

  .status-indicator {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    font-weight: 600;
  }

  .status-dot {
    width: 12px;
    height: 12px;
    border-radius: 50%;
    animation: pulse 2s infinite;
  }

  .online { background-color: #10b981; }
  .offline { background-color: #ef4444; }

  .toggle-btn {
    padding: 0.75rem 1.5rem;
    border: none;
    border-radius: 8px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
  }

  .btn-online {
    background-color: #ef4444;
    color: white;
  }

  .btn-offline {
    background-color: #10b981;
    color: white;
  }

  .historico-section {
    background: white;
    padding: 2rem;
    border-radius: 12px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }

  .historico-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    border-left: 4px solid;
    background: #f8fafc;
    margin-bottom: 0.5rem;
    border-radius: 0 8px 8px 0;
  }

  .historico-info {
    flex: 1;
  }

  .historico-status {
    font-weight: 600;
    padding: 0.25rem 0.75rem;
    border-radius: 16px;
    font-size: 0.875rem;
    color: white;
  }

  .admin-actions {
    margin-top: 1rem;
    display: flex;
    gap: 1rem;
  }

  .admin-btn {
    padding: 0.5rem 1rem;
    background: #6b7280;
    color: white;
    border: none;
    border-radius: 6px;
    cursor: pointer;
    font-size: 0.875rem;
  }
</style>

<div class="painel-container">
  <div class="header">
    <div>
      <h1>Painel do Motorista #{motoristaId}</h1>
      <div class="status-indicator">
        <div class="status-dot" class:online={isOnline} class:offline={!isOnline}></div>
        Status: {isOnline ? 'Online' : 'Offline'}
      </div>
    </div>
    
    <button class="toggle-btn" 
            class:btn-online={isOnline} 
            class:btn-offline={!isOnline}
            on:click={toggleOnlineStatus}>
      {isOnline ? 'Ficar Offline' : 'Ficar Online'}
    </button>
  </div>

  <div class="historico-section">
    <h2>Histórico de Notificações</h2>
    
    {#if historicoNotificacoes.length === 0}
      <p>Nenhuma notificação no histórico.</p>
    {:else}
      {#each historicoNotificacoes as item}
        <div class="historico-item" style="border-left-color: {getStatusColor(item.Status)}">
          <div class="historico-info">
            <div><strong>Corrida #{item.CorridaID}</strong> - {item.PassageiroNome}</div>
            <div>R$ {item.Valor?.toFixed(2)} - {formatTimestamp(item.CriadoEm)}</div>
          </div>
          <div class="historico-status" style="background-color: {getStatusColor(item.Status)}">
            {item.Status}
          </div>
        </div>
      {/each}
    {/if}

    <div class="admin-actions">
      <button class="admin-btn" on:click={forceExpireNotifications}>
        Expirar Notificações Antigas
      </button>
      <button class="admin-btn" on:click={loadHistorico}>
        Recarregar Histórico
      </button>
    </div>
  </div>
</div>

<!-- Popup de Notificação -->
<NotificacaoPopup 
  {motoristaId}
  notificacao={currentNotificacao}
  on:accepted={handleNotificationAccepted}
  on:refused={handleNotificationRefused}
  on:expired={handleNotificationExpired}
/>