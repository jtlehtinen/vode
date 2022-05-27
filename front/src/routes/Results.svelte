<script>
  let totalVotes = 10;
  let question = 'Which one?';
  let pollID = 123;

  let options = ['A', 'B', 'C'];
</script>

<div class="field-row" style="width: 100%">
  <p style="font-weight: bold;">{question}</p>
</div>

{#each options as option}
<div class="field-row" style="justify-content: space-between;">
  <p>{option}</p>
  <p>1 votes / 25%</p>
</div>
<hr>
{/each}

<div class="field-row" style="justify-content: space-between;">
  <p style="font-weight: bold;">{totalVotes} Votes Total</p>
  <button type="button" onclick="window.location.href='/poll/{pollID}';">Vote</button>
</div>
