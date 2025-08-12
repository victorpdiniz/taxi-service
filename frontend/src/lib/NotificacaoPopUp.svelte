<script>
  import { createEventDispatcher } from 'svelte';
  import axios from 'axios';
  import Timer from './Timer.svelte';
  import InfoGrid from './InfoGrid.svelte';
  import RouteInfo from './RouteInfo.svelte';
  import ActionButtons from './ActionButtons.svelte';

  export let notificacao = null;
  export let motoristaId = 1;

  const dispatch = createEventDispatcher();
  const api = axios.create({ baseURL: 'http://localhost:3000/api' });

  let isVisible = false;
  let timer;
  let timeLeft = 20;
  let loading = false;

  $: if (notificacao) {
    showNotification();
  }

  $: infoItems = notificacao ? [
    { label: 'Passageiro', value: notificacao.PassageiroNome || 'Usuário' },
    { label: 'Valor da Corrida', value: `R$ ${notificacao.Valor?.toFixed(2) || '0,00'}` }
  ] : [];

  function showNotification() {
    isVisible = true;
    timeLeft = 20;
    timer?.start();
  }

  async function handleAccept() {
    loading = true;
    try {
      await api.post(`/notificacoes/${notificacao.ID}/motorista/${motoristaId}/accept`);
      
      dispatch('accepted', {
        notificacaoId: notificacao.ID,
        corridaId: notificacao.CorridaID
      });
      
      closeNotification();
      
    } catch (error) {
      console.error('Erro ao aceitar notificação:', error);
      
      if (error.response?.status === 410) {
        alert('Esta notificação já expirou.');
      } else if (error.response?.status === 409) {
        alert('Esta notificação já foi processada por outro motorista.');
      } else {
        alert('Erro ao aceitar a corrida. Tente novamente.');
      }
      
      closeNotification();
    } finally {
      loading = false;
    }
  }

  async function handleRefuse() {
    loading = true;
    try {
      await api.post(`/notificacoes/${notificacao.ID}/motorista/${motoristaId}/refuse`);
      
      dispatch('refused', {
        notificacaoId: notificacao.ID,
        corridaId: notificacao.CorridaID
      });
      
      closeNotification();
      
    } catch (error) {
      console.error('Erro ao recusar notificação:', error);
      alert('Erro ao recusar a corrida.');
      closeNotification();
    } finally {
      loading = false;
    }
  }

  function handleExpire() {
    dispatch('expired', {
      notificacaoId: notificacao.ID,
      corridaId: notificacao.CorridaID
    });
    
    closeNotification();
  }

  function closeNotification() {
    isVisible = false;
    timer?.stop();
    notificacao = null;
    loading = false;
  }
</script>

{#if isVisible && notificacao}
  <div class="notification-overlay">
    <div class="notification-popup" class:urgent={timeLeft <= 5}>
      <div class="notification-header">
        <h2 class="notification-title">🚗 Nova Corrida Disponível!</h2>
        <Timer bind:this={timer} 
               bind:timeLeft 
               on:expired={handleExpire} />
      </div>

      <div class="notification-body">
        <InfoGrid items={infoItems} />
        <RouteInfo origem={notificacao.Origem} 
                   destino={notificacao.Destino} />
      </div>

      <ActionButtons {timeLeft} 
                     {loading}
                     on:accept={handleAccept}
                     on:refuse={handleRefuse} />
    </div>
  </div>
{/if}

<style>
  .notification-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.8);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999;
    animation: fadeIn 0.3s ease-in-out;
  }

  .notification-popup {
    background: white;
    border-radius: 16px;
    padding: 2rem;
    max-width: 500px;
    width: 90%;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3);
    animation: slideIn 0.3s ease-out;
  }

  .notification-header {
    text-align: center;
    margin-bottom: 1.5rem;
  }

  .notification-title {
    font-size: 1.5rem;
    font-weight: bold;
    color: #1f2937;
    margin-bottom: 0.5rem;
  }

  .notification-body {
    margin-bottom: 2rem;
  }

  .urgent {
    animation: pulse 1s infinite;
  }

  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }

  @keyframes slideIn {
    from { 
      opacity: 0;
      transform: translate(-50%, -60%) scale(0.8);
    }
    to { 
      opacity: 1;
      transform: translate(-50%, -50%) scale(1);
    }
  }

  @keyframes pulse {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.05); }
  }
</style>