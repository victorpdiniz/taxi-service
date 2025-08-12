<script>
  import { createEventDispatcher } from 'svelte';
  
  export let timeLeft = 20;
  export let loading = false;
  
  const dispatch = createEventDispatcher();
</script>

<div class="notification-actions">
  <button class="btn btn-refuse" 
          on:click={() => dispatch('refuse')}
          disabled={loading}>
    ❌ Recusar
  </button>
  <button class="btn btn-accept" 
          on:click={() => dispatch('accept')}
          class:pulse-animation={timeLeft <= 10}
          disabled={loading}>
    ✅ Aceitar
  </button>
</div>

<style>
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

  .btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .btn-accept {
    background: linear-gradient(135deg, #10b981 0%, #059669 100%);
    color: white;
    box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
  }

  .btn-accept:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
  }

  .btn-refuse {
    background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
    color: white;
    box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3);
  }

  .btn-refuse:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(239, 68, 68, 0.4);
  }

  .pulse-animation {
    animation: pulse 2s infinite;
  }

  @keyframes pulse {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.05); }
  }
</style>