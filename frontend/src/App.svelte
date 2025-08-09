<script>
  import { Router, Route } from "svelte-routing";
  import Home from "./lib/Home.svelte";
  import Corrida from "./lib/Corrida.svelte";
</script>

<Router>
  <Route path="/" component={Home} />
  <Route path="/corrida" component={Corrida} />
</Router>
