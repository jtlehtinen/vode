<script>
  import { Router, Link, Route } from 'svelte-routing';
  import Create from './routes/Create.svelte';
  import Vote from './routes/Vote.svelte';
  import Results from './routes/Results.svelte';

  export let url = '';
</script>

<div class='container'>
  <div class='paper col-12'>
    <Router url='{url}'>
      <Route path='/' component={Create}/>
      <Route path='/vote/:id' component={Vote}/>
      <Route path='/results/:id' component={Results}/>
    </Router>
  </div>
</div>

<style>
  .container {
    width: 100%;
    display: flex;
    justify-content: center;
  }
</style>
