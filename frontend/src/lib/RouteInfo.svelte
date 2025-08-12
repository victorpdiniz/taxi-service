<script>
  export let origem = '';
  export let destino = '';
</script>

<div class="route-info">
  <div class="route-item">
    <span class="route-icon">📍</span>
    <div>
      <div class="info-label">Origem</div>
      <div class="info-value">{origem || 'Local de partida'}</div>
    </div>
  </div>
  
  <div class="route-item">
    <span class="route-icon">🎯</span>
    <div>
      <div class="info-label">Destino</div>
      <div class="info-value">{destino || 'Local de destino'}</div>
    </div>
  </div>
</div>

<style>
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

  .route-item:last-child {
    margin-bottom: 0;
  }

  .route-icon {
    font-size: 1.2rem;
    margin-right: 0.75rem;
    width: 24px;
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
</style>