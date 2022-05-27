<script>
  let question = 'Which one?';

  let options = [
    {id: 0, text: 'Option A'},
    {id: 1, text: 'Option B'},
    {id: 2, text: 'Option C'},
  ];

  function submit(event) {
    event.preventDefault();
  }
</script>

<form on:submit={submit}>
  <fieldset>
    <legend>{question}</legend>
    <br>

    {#each options as option}
      <div class='row'>
        <input type='radio' id='{option.id}' name='option' value='{option.id}'>
        <label for='{option.id}' class='paper-radio'>{option.text}</label>
      </div>
    {/each}
    <br>

    <div class='row'>
      <div class='col-2'>
        <button type='submit' class='btn-success'>Vote</button>
      </div>
    </div>
  </fieldset>
</form>

<style>
  label {
    margin-left: 0.5em;
  }

  button {
    width: 100%;
  }
</style>
