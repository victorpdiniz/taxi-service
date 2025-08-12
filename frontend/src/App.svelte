<script>
  import { Router, Route } from "svelte-routing";
  import Home from "./lib/Home.svelte";
  import Corrida from "./lib/Corrida.svelte";
  import PainelMotorista from "./lib/PainelMotorista.svelte";
  import DemoNotificacao from "./lib/DemoNotificacao.svelte";
</script>

<Router>
  <Route path="/" component={Home} />
  <Route path="/corrida/:id" component={Corrida} />
  <Route path="/motorista" component={PainelMotorista} />
  <Route path="/demo-notificacao" component={DemoNotificacao} />
</Router>
