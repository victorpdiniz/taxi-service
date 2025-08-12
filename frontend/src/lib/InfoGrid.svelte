<script>
  export let items = [];
</script>

<div class="info-grid">
  {#each items as item}
    <div class="info-item">
      <div class="info-label">{item.label}</div>
      <div class="info-value">{item.value}</div>
    </div>
  {/each}
</div>

<style>
  .info-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
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
</style>