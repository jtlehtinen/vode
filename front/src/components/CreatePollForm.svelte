<script>
  import PollOption from './PollOption.svelte';

  const MAX_OPTIONS = 5;

  let nextId = 2;
  let options = [
    {id: 0, text: ''},
    {id: 1, text: ''},
  ];

  function addOption() {
    if (options.length >= MAX_OPTIONS) return;
    options = [...options, {id: nextId++, text: ''}]
  }

  function removeOption(id) {
    const idx = options.findIndex(option => option.id === id);
    if (idx !== -1) {
      options.splice(idx, 1);
      options = options; // happy svelte...
    }
  }

  function submit(event) {
    event.preventDefault();
  }
</script>

<form on:submit={submit}>
  <fieldset>
    <legend>New Poll</legend>
    <br>

    <div class='row'>
      <div class='col-8'>
        <input type='text' name='question' autocomplete='off' placeholder='Type your question here'>
      </div>
      <div class='col-2'>
        <button type='submit' class='btn-success'>
          Create Poll
        </button>
      </div>
    </div>
    <br>

    <div class='row'>
      <div class='col-2'>
        <button type='button' class='btn-secondary' disabled={options.length >= MAX_OPTIONS} on:click={addOption}>
          Add Option
        </button>
      </div>
    </div>

    <div class='row'>
      {#each options as option (option.id)}
      <div class='col-12'>
        <PollOption on:remove={() => removeOption(option.id)}/>
      </div>
      {/each}
    </div>
    <br>
  </fieldset>
</form>

<style>
  input {
    width: 100%;
    height: 100%;
  }

  button {
    width: 100%;
  }

  button[type=submit] {
    margin-left: 1.0em;
  }
</style>
