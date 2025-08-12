<script>
  import NotificacaoPopup from './NotificacaoPopup.svelte';
  import axios from 'axios';
  
  let notificacaoComponent;
  let currentNotification = null;
  let showingNotification = false;

  const api = axios.create({
    baseURL: 'http://localhost:3000/api',
  });

  // Mock data examples - Updated to match the Go model structure
  const mockNotifications = [
    {
      id: 1,
      motorista_id: 101,
      corrida_id: 1001,
      passageiro_nome: "Maria Silva",
      valor: 25.50,
      distancia_km: 5.2,
      tempo_estimado: "12 min",
      origem: "-23.5505, -46.6333",
      destino: "-23.5550, -46.6400",
      status: "pendente",
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
      expira_em: new Date(Date.now() + 20000).toISOString()
    },
    {
      id: 2,
      motorista_id: 102,
      corrida_id: 1002,
      passageiro_nome: "João Santos",
      valor: 42.80,
      distancia_km: 8.7,
      tempo_estimado: "18 min",
      origem: "-23.5600, -46.6450",
      destino: "-23.5350, -46.6200",
      status: "pendente",
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
      expira_em: new Date(Date.now() + 20000).toISOString()
    },
    {
      id: 3,
      motorista_id: 103,
      corrida_id: 1003,
      passageiro_nome: "Ana Costa",
      valor: 18.90,
      distancia_km: 3.1,
      tempo_estimado: "8 min",
      origem: "-23.5480, -46.6320",
      destino: "-23.5520, -46.6380",
      status: "pendente",
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
      expira_em: new Date(Date.now() + 20000).toISOString()
    },
    {
      id: 4,
      motorista_id: 104,
      corrida_id: 1004,
      passageiro_nome: "Carlos Oliveira",
      valor: 67.30,
      distancia_km: 15.8,
      tempo_estimado: "25 min",
      origem: "-23.5200, -46.6000",
      destino: "-23.5800, -46.6800",
      status: "pendente",
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
      expira_em: new Date(Date.now() + 20000).toISOString()
    }
  ];

  async function createNotificationViaAPI(notificationData) {
    try {
      console.log('Criando notificação via API:', notificationData);
      
      const response = await api.post('/notificacoes', notificationData);
      
      console.log('Resposta da API - Notificação criada:', response.data);
      
      return response.data;
    } catch (error) {
      console.error('Erro ao criar notificação via API:', error);
      
      if (error.response) {
        console.error('Dados do erro:', error.response.data);
        alert(`Erro ao criar notificação: ${error.response.data.message || error.response.statusText}`);
      } else {
        alert('Erro de conectividade com a API');
      }
      
      throw error;
    }
  }

  async function triggerRandomNotification() {
    if (showingNotification) {
      alert('Já existe uma notificação ativa. Aguarde ela finalizar.');
      return;
    }

    const randomNotification = mockNotifications[Math.floor(Math.random() * mockNotifications.length)];
    
    try {
      // Primeiro, cria a notificação via API
      const createdNotification = await createNotificationViaAPI(randomNotification);
      
      // Depois, mostra o popup
      currentNotification = createdNotification;
      showingNotification = true;
      
    } catch (error) {
      // Se der erro na API, mostra em modo mock
      console.log('Exibindo em modo mock devido a erro na API');
      currentNotification = { ...randomNotification, useMockData: true };
      showingNotification = true;
    }
  }

  async function triggerSpecificNotification(index) {
    if (showingNotification) {
      alert('Já existe uma notificação ativa. Aguarde ela finalizar.');
      return;
    }

    const selectedNotification = mockNotifications[index];
    
    try {
      // Primeiro, cria a notificação via API
      const createdNotification = await createNotificationViaAPI(selectedNotification);
      
      // Depois, mostra o popup
      currentNotification = createdNotification;
      showingNotification = true;
      
    } catch (error) {
      // Se der erro na API, mostra em modo mock
      console.log('Exibindo em modo mock devido a erro na API');
      currentNotification = { ...selectedNotification, useMockData: true };
      showingNotification = true;
    }
  }

  function handleNotificationAccepted(event) {
    console.log('Notificação aceita:', event.detail);
    
    if (event.detail.response) {
      console.log('Resposta da API de aceitar:', event.detail.response);
    }
    
    showingNotification = false;
    currentNotification = null;
    
    // Aqui você pode adicionar lógica adicional após aceitar
    setTimeout(() => {
      alert(`Corrida ${event.detail.corridaId} foi aceita! Redirecionando para o painel...`);
    }, 500);
  }

  function handleNotificationRefused(event) {
    console.log('Notificação recusada:', event.detail);
    
    if (event.detail.response) {
      console.log('Resposta da API de recusar:', event.detail.response);
    }
    
    showingNotification = false;
    currentNotification = null;
    
    // Aqui você pode adicionar lógica adicional após recusar
    setTimeout(() => {
      alert('Corrida recusada. Aguardando próximas oportunidades...');
    }, 500);
  }

  function handleNotificationExpired(event) {
    console.log('Notificação expirada:', event.detail);
    
    if (event.detail.response) {
      console.log('Resposta da API de expirar:', event.detail.response);
    }
    
    showingNotification = false;
    currentNotification = null;
    
    // Aqui você pode adicionar lógica adicional após expirar
    setTimeout(() => {
      alert('Tempo esgotado! A notificação expirou.');
    }, 500);
  }

  // Função para buscar notificações pendentes do motorista
  async function fetchPendingNotifications() {
    try {
      const motoristaId = 101; // ID do motorista logado
      const response = await api.get(`/notificacoes/motorista/${motoristaId}/pending`);
      
      console.log('Notificações pendentes:', response.data);
      
      if (response.data && response.data.length > 0) {
        // Se há notificações pendentes, mostra a primeira
        const notification = response.data[0];
        currentNotification = notification;
        showingNotification = true;
      } else {
        alert('Nenhuma notificação pendente encontrada.');
      }
      
    } catch (error) {
      console.error('Erro ao buscar notificações pendentes:', error);
      alert('Erro ao buscar notificações pendentes.');
    }
  }
</script>

<style>
  .demo-container {
    max-width: 800px;
    margin: 2rem auto;
    padding: 2rem;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  }

  .demo-header {
    text-align: center;
    margin-bottom: 2rem;
  }

  .demo-title {
    font-size: 2rem;
    color: #1f2937;
    margin-bottom: 0.5rem;
  }

  .demo-subtitle {
    color: #6b7280;
    font-size: 1.1rem;
  }

  .controls-section {
    background: white;
    border-radius: 12px;
    padding: 2rem;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    margin-bottom: 2rem;
  }

  .controls-title {
    font-size: 1.3rem;
    color: #1f2937;
    margin-bottom: 1rem;
  }

  .button-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 1rem;
    margin-bottom: 1.5rem;
  }

  .demo-button {
    padding: 1rem;
    border: none;
    border-radius: 8px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    font-size: 1rem;
  }

  .btn-primary {
    background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
    color: white;
    box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
  }

  .btn-primary:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(59, 130, 246, 0.4);
  }

  .btn-secondary {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    color: white;
    box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
  }

  .btn-secondary:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
  }

  .btn-warning {
    background: linear-gradient(135deg, #f59e0b 0%, #d97706 100%);
    color: white;
    box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
  }

  .btn-warning:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(245, 158, 11, 0.4);
  }

  .notification-examples {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 1rem;
  }

  .notification-card {
    background: #f8fafc;
    border: 1px solid #e2e8f0;
    border-radius: 8px;
    padding: 1rem;
    cursor: pointer;
    transition: all 0.3s ease;
  }

  .notification-card:hover {
    background: #e2e8f0;
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }

  .card-header {
    font-weight: 600;
    color: #1f2937;
    margin-bottom: 0.5rem;
  }

  .card-details {
    font-size: 0.875rem;
    color: #6b7280;
  }

  .instructions {
    background: #fef3c7;
    border: 1px solid #f59e0b;
    border-radius: 8px;
    padding: 1rem;
    margin-bottom: 2rem;
  }

  .instructions-title {
    font-weight: 600;
    color: #92400e;
    margin-bottom: 0.5rem;
  }

  .instructions-text {
    color: #92400e;
    font-size: 0.875rem;
  }

  .api-section {
    background: #f0f9ff;
    border: 1px solid #0ea5e9;
    border-radius: 8px;
    padding: 1rem;
    margin-bottom: 2rem;
  }

  .api-title {
    font-weight: 600;
    color: #0c4a6e;
    margin-bottom: 0.5rem;
  }

  .api-text {
    color: #0c4a6e;
    font-size: 0.875rem;
  }
</style>

<div class="demo-container">
  <div class="demo-header">
    <h1 class="demo-title">🚗 Demonstração de Notificações</h1>
    <p class="demo-subtitle">Teste as notificações de corrida do sistema de táxi com integração API</p>
  </div>

  <div class="instructions">
    <div class="instructions-title">📋 Instruções:</div>
    <div class="instructions-text">
      Clique nos botões abaixo para simular notificações de corrida. Cada notificação tem 20 segundos para ser aceita ou recusada.
      As notificações são criadas via API e as ações (aceitar/recusar) também são enviadas para o backend.
    </div>
  </div>

  <div class="api-section">
    <div class="api-title">🔗 Integração com API:</div>
    <div class="api-text">
      • <strong>POST</strong> /api/notificacoes - Criar notificação<br>
      • <strong>POST</strong> /api/notificacoes/:id/motorista/:motoristaID/accept - Aceitar<br>
      • <strong>POST</strong> /api/notificacoes/:id/motorista/:motoristaID/refuse - Recusar<br>
      • <strong>PUT</strong> /api/notificacoes/:id/status - Atualizar status<br>
      • <strong>GET</strong> /api/notificacoes/motorista/:motoristaID/pending - Buscar pendentes
    </div>
  </div>

  <div class="controls-section">
    <h2 class="controls-title">Controles de Demonstração</h2>
    
    <div class="button-grid">
      <button 
        class="demo-button btn-primary" 
        on:click={triggerRandomNotification}
        disabled={showingNotification}>
        🎲 Notificação Aleatória (API)
      </button>
      
      <button 
        class="demo-button btn-warning" 
        on:click={fetchPendingNotifications}
        disabled={showingNotification}>
        📥 Buscar Pendentes (API)
      </button>
    </div>

    <h3 class="controls-title">Notificações Específicas (Criadas via API):</h3>
    <div class="notification-examples">
      {#each mockNotifications as notification, index}
        <div 
          class="notification-card" 
          role="button"
          tabindex="0"
          on:click={() => triggerSpecificNotification(index)}
          on:keydown={(e) => e.key === 'Enter' || e.key === ' ' ? triggerSpecificNotification(index) : null}>
          <div class="card-header">{notification.passageiro_nome}</div>
          <div class="card-details">
            ID: {notification.id}<br>
            Valor: R$ {notification.valor.toFixed(2)}<br>
            Distância: {notification.distancia_km} km<br>
            Tempo: {notification.tempo_estimado}
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>

<!-- Componente de Notificação -->
<NotificacaoPopup 
  bind:this={notificacaoComponent}
  bind:notificacao={currentNotification}
  motoristaId={101}
  useMockData={currentNotification?.useMockData || false}
  on:accepted={handleNotificationAccepted}
  on:refused={handleNotificationRefused}
  on:expired={handleNotificationExpired}
/>