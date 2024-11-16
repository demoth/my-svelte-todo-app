<script>
  import TodoList from './lib/TodoList.svelte';
</script>

<main>
  <h1>Todo App</h1>
  <TodoList />
</main>

<style>
  main {
    max-width: 500px;
    margin: 0 auto;
    padding: 20px;
    font-family: Arial, sans-serif;
  }
  h1 {
    text-align: center;
    color: #333;
  }
</style>
