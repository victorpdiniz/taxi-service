<script>
  import { createEventDispatcher, onDestroy } from 'svelte';
  
  export let initialTime = 20;
  export let autoStart = true;
  
  const dispatch = createEventDispatcher();
  
  let timeLeft = initialTime;
  let timer = null;
  
  export function start() {
    if (timer) clearInterval(timer);
    timeLeft = initialTime;
    
    timer = setInterval(() => {
      timeLeft--;
      if (timeLeft <= 0) {
        dispatch('expired');
        stop();
      }
    }, 1000);
  }
  
  export function stop() {
    if (timer) {
      clearInterval(timer);
      timer = null;
    }
  }
  
  if (autoStart) start();
  
  onDestroy(() => stop());
</script>

<div class="timer-display" 
     class:timer-urgent={timeLeft <= 5}
     class:timer-warning={timeLeft <= 10 && timeLeft > 5}
     class:timer-normal={timeLeft > 10}>
  ⏱️ {timeLeft}s restantes
</div>

<style>
  .timer-display {
    font-size: 1.2rem;
    font-weight: bold;
    padding: 0.5rem 1rem;
    border-radius: 8px;
    display: inline-block;
  }
  
  .timer-urgent { color: #dc2626; background: #fee2e2; }
  .timer-warning { color: #f59e0b; background: #fef3c7; }
  .timer-normal { color: #059669; background: #d1fae5; }
</style>