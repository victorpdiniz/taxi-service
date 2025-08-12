<script>
  import { createEventDispatcher } from 'svelte';
  
  export let corrida;
  
  const dispatch = createEventDispatcher();
  
  function handleClick() {
    dispatch('solicitar', corrida);
  }
</script>

<button class="corrida-card" on:click={handleClick}>
  <div class="header">
    <strong>{corrida.PassageiroNome}</strong>
    <span class="valor">R$ {corrida.Valor.toFixed(2)}</span>
  </div>
  <div class="rota">
    📍 {corrida.Origem} → 🎯 {corrida.Destino}
  </div>
  <div class="detalhes">
    📏 {corrida.DistanciaKm.toFixed(1)}km | ⏱️ {corrida.TempoEstimado}min | #{corrida.CorridaID}
  </div>
  <div class="action">🔔 Simular Corrida</div>
</button>

<style>
  .corrida-card {
    background: white;
    border: 1px solid #ddd;
    border-radius: 8px;
    padding: 1rem;
    cursor: pointer;
    transition: all 0.2s;
    text-align: left;
    font-family: inherit;
  }

  .corrida-card:hover {
    box-shadow: 0 4px 8px rgba(0,0,0,0.1);
    transform: translateY(-2px);
  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 0.5rem;
  }

  .valor {
    color: #10b981;
    font-weight: bold;
  }

  .rota {
    margin-bottom: 0.5rem;
    color: #666;
  }

  .detalhes {
    margin-bottom: 1rem;
    color: #888;
    font-size: 0.9rem;
  }

  .action {
    background: #667eea;
    color: white;
    padding: 0.5rem;
    border-radius: 4px;
    text-align: center;
    font-weight: bold;
  }
</style>