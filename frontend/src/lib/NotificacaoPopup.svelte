<script>
  import { createEventDispatcher, onMount, onDestroy } from 'svelte';
  import axios from 'axios';

  export let notificacao = null;
  export let motoristaId = 1; // ID do motorista logado

  const dispatch = createEventDispatcher();
  const api = axios.create({
    baseURL: 'http://localhost:3000/api',
  });

  let timeLeft = 20; // 20 segundos para responder
  let timer = null;
  let isVisible = false;

  $: if (notificacao) {
    showNotification();
  }

  function showNotification() {
    isVisible = true;
    timeLeft = 20;
    startTimer();
  }

  function startTimer() {
    if (timer) clearInterval(timer);
    
    timer = setInterval(() => {
      timeLeft--;
      if (timeLeft <= 0) {
        // Tempo esgotado - auto-expirar
        handleExpire();
      }
    }, 1000);
  }

  async function handleAccept() {
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
    }
  }

  async function handleRefuse() {
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
    if (timer) {
      clearInterval(timer);
      timer = null;
    }
    notificacao = null;
  }

  onDestroy(() => {
    if (timer) {
      clearInterval(timer);
    }
  });
</script>

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
    position: relative;
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

  .timer-display {
    font-size: 1.2rem;
    font-weight: bold;
    color: #dc2626;
    background: #fee2e2;
    padding: 0.5rem 1rem;
    border-radius: 8px;
    display: inline-block;
  }

  .notification-body {
    margin-bottom: 2rem;
  }

  .info-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1rem;
    margin-bottom: 1rem;
  }

  .info-item {
    background: #f8fafc;
    padding: 1rem;
    border-radius: 8px;
    border-left: 4px solid #3b82f6;
  }

  .info-label {
    font-size: 0.875rem;
    color: #6b7280;
    font-weight: 600;
    margin-bottom: 0.25rem;
  }

  .info-value {
    font-size: 1rem;
    color: #1f2937;
    font-weight: 500;
  }

  .route-info {
    background: #ecfdf5;
    border: 1px solid #d1fae5;
    border-radius: 8px;
    padding: 1rem;
    margin-bottom: 1rem;
  }

  .route-item {
    display: flex;
    align-items: center;
    margin-bottom: 0.5rem;
  }

  .route-icon {
    font-size: 1.2rem;
    margin-right: 0.75rem;
    width: 24px;
  }

  .notification-actions {
    display: flex;
    gap: 1rem;
  }

  .btn {
    flex: 1;
    padding: 1rem;
    font-size: 1.1rem;
    font-weight: 600;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s ease;
    text-transform: uppercase;
    letter-spacing: 0.5px;
  }

  .btn-accept {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    color: white;
    box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
  }

  .btn-accept:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
  }

  .btn-refuse {
    background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
    color: white;
    box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3);
  }

  .btn-refuse:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(239, 68, 68, 0.4);
  }

  .pulse-animation {
    animation: pulse 2s infinite;
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

  .urgent {
    animation: pulse 1s infinite;
  }

  .timer-urgent {
    color: #dc2626;
    background: #fee2e2;
  }

  .timer-warning {
    color: #f59e0b;
    background: #fef3c7;
  }

  .timer-normal {
    color: #059669;
    background: #d1fae5;
  }
</style>

{#if isVisible && notificacao}
  <div class="notification-overlay">
    <div class="notification-popup" class:urgent={timeLeft <= 5}>
      <div class="notification-header">
        <h2 class="notification-title">🚗 Nova Corrida Disponível!</h2>
        <div class="timer-display" 
             class:timer-urgent={timeLeft <= 5}
             class:timer-warning={timeLeft <= 10 && timeLeft > 5}
             class:timer-normal={timeLeft > 10}>
          ⏱️ {timeLeft}s restantes
        </div>
      </div>

      <div class="notification-body">
        <div class="info-grid">
          <div class="info-item">
            <div class="info-label">Passageiro</div>
            <div class="info-value">{notificacao.PassageiroNome || 'Usuário'}</div>
          </div>
          
          <div class="info-item">
            <div class="info-label">Valor da Corrida</div>
            <div class="info-value">R$ {notificacao.Valor?.toFixed(2) || '0,00'}</div>
          </div>
        </div>

        <div class="route-info">
          <div class="route-item">
            <span class="route-icon">📍</span>
            <div>
              <div class="info-label">Origem</div>
              <div class="info-value">{notificacao.Origem || 'Local de partida'}</div>
            </div>
          </div>
          
          <div class="route-item">
            <span class="route-icon">🎯</span>
            <div>
              <div class="info-label">Destino</div>
              <div class="info-value">{notificacao.Destino || 'Local de destino'}</div>
            </div>
          </div>
        </div>
      </div>

      <div class="notification-actions">
        <button class="btn btn-refuse" on:click={handleRefuse}>
          ❌ Recusar
        </button>
        <button class="btn btn-accept" on:click={handleAccept} class:pulse-animation={timeLeft <= 10}>
          ✅ Aceitar
        </button>
      </div>
    </div>
  </div>
{/if}