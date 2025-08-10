<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import axios from 'axios';

  // O ID da corrida será passado pelo router
  export let id;

  let mapElement;
  let map;
  let ride = null; // Objeto completo da corrida
  let rideStatus = 'Carregando informações...';
  let motoristaMarker;

  const api = axios.create({
    baseURL: 'http://localhost:3000/api',
  });

  onMount(() => {
    map = L.map(mapElement).setView([-23.55052, -46.633308], 14);

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);

    // Busca os dados da corrida e inicia o monitoramento
    fetchRideData();
    const interval = setInterval(fetchRideData, 3000); // Verifica a cada 3 segundos

    return () => clearInterval(interval); // Limpa o intervalo quando o componente é destruído
  });

  async function fetchRideData() {
    try {
      const response = await api.get(`/corrida/${id}`);
      ride = response.data;
      rideStatus = formatStatus(ride.Status);

      // Atualiza a posição do motorista no mapa
      if (ride.MotoristaLat && ride.MotoristaLng) {
        const latLng = [ride.MotoristaLat, ride.MotoristaLng];
        if (!motoristaMarker) {
          motoristaMarker = L.marker(latLng).addTo(map).bindPopup('Motorista');
          map.setView(latLng, 15);
        } else {
          motoristaMarker.setLatLng(latLng);
        }
      }

      // Se a corrida foi finalizada ou cancelada, para de monitorar e redireciona
      if (ride.Status.startsWith('concluída') || ride.Status.startsWith('cancelada')) {
        setTimeout(() => navigate('/'), 3000);
      }

    } catch (error) {
      console.error('Erro ao buscar dados da corrida:', error);
      rideStatus = 'Erro ao carregar dados.';
    }
  }

  async function cancelRide() {
    if (confirm('Tem certeza que deseja cancelar a corrida?')) {
      try {
        await api.post(`/corrida/${id}/cancelar`);
        alert('Sua corrida foi cancelada.');
        navigate('/');
      } catch (error) {
        console.error('Erro ao cancelar a corrida:', error);
        alert('Não foi possível cancelar a corrida.');
      }
    }
  }
  
  async function finishRide() {
    if (confirm('Confirmar a finalização da corrida?')) {
      try {
        await api.post(`/corrida/${id}/finalizar`);
        alert('Corrida finalizada com sucesso!');
        navigate('/');
      } catch (error) {
        console.error('Erro ao finalizar a corrida:', error);
        alert('Não foi possível finalizar a corrida.');
      }
    }
  }

  function formatStatus(status) {
    return status.replace(/_/g, ' ').replace(/\b\w/g, l => l.toUpperCase());
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
    padding: 1rem;
    background-color: #f0f0f0;
    border-radius: 8px;
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
    border: none;
    border-radius: 5px;
    color: white;
  }
  .cancel-btn {
      background-color: #f44336;
  }
  .finish-btn {
      background-color: #4CAF50;
  }
</style>

<div class="container">
  <h1>Sua Corrida (ID: {id})</h1>

  <div id="map" bind:this={mapElement}></div>

  <div class="status">Status: {rideStatus}</div>

  <div class="actions">
    <button class="cancel-btn" on:click={cancelRide}>Cancelar Corrida</button>
    <button class="finish-btn" on:click={finishRide}>Finalizar Corrida</button>
  </div>
</div>
