<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();
  function remove() {
    dispatch('remove', {});
  }
</script>

<div class='row'>
  <input type='text' name='option' class='col-8' autocomplete='off' placeholder='Enter poll option'>
  <button type='button' class='btn-danger col-2' on:click={remove}>Remove</button>
</div>

<style>
  button {
    margin-left: 1em;
  }
</style>
