<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import axios from 'axios';

  let mapElement;
  let map;
  let rideStatus = 'Procurando motorista...';

  onMount(() => {
    map = L.map(mapElement).setView([-23.55052, -46.633308], 14);

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);

    // TODO: Get ride data from a store or API
    // and draw the route on the map.

    // Periodically check for ride status
    const interval = setInterval(async () => {
      try {
        // This endpoint should return the current status of the ride
        const response = await axios.post('http://localhost:3000/api/corrida/monitorar'); 
        rideStatus = response.data.status || 'Em andamento';
        
        if (response.data.status === 'finalizada' || response.data.status === 'cancelada') {
          clearInterval(interval);
          navigate('/');
        }
      } catch (error) {
        console.error('Erro ao monitorar a corrida:', error);
      }
    }, 5000); // Check every 5 seconds
  });

  async function cancelRide() {
    try {
      await axios.post('http://localhost:3000/api/corrida/cancelar-por-excesso-tempo');
      alert('Sua corrida foi cancelada.');
      navigate('/');
    } catch (error) {
      console.error('Erro ao cancelar a corrida:', error);
      alert('Não foi possível cancelar a corrida.');
    }
  }
  
  async function finishRide() {
    try {
      await axios.post('http://localhost:3000/api/corrida/finalizar');
      alert('Corrida finalizada com sucesso!');
      navigate('/');
    } catch (error) {
      console.error('Erro ao finalizar a corrida:', error);
      alert('Não foi possível finalizar a corrida.');
    }
  }

</script>

<style>
  .container {
    padding: 2rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
  #map {
    height: 400px;
    width: 100%;
    background-color: #eee;
  }
  .status {
    font-size: 1.2rem;
    font-weight: bold;
    text-align: center;
  }
  .actions {
      display: flex;
      gap: 1rem;
  }
  button {
    padding: 0.75rem;
    font-size: 1rem;
    cursor: pointer;
    flex: 1;
  }
  .cancel-btn {
      background-color: #f44336;
      color: white;
  }
</style>

<div class="container">
  <h1>Sua Corrida</h1>

  <div id="map" bind:this={mapElement}></div>

  <div class="status">{rideStatus}</div>

  <div class="actions">
    <button class="cancel-btn" on:click={cancelRide}>Cancelar Corrida</button>
    <button on:click={finishRide}>Finalizar Corrida</button>
  </div>
</div>
