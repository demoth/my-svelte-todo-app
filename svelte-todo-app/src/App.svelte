<script>
  import TodoList from './lib/TodoList.svelte';
  import { todos } from './lib/stores.js';
  
  // Calculate total items
  $: totalItems = $todos.length;
</script>

<main>
  <h1>Daniil's Svelte Todo App</h1>
  <TodoList />
</main>

<footer>
  <p>Total items in list: {totalItems}</p>
</footer>

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
  footer {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: #f4f4f4;
    padding: 15px;
    text-align: center;
    border-top: 1px solid #ddd;
    font-family: Arial, sans-serif;
  }
  footer p {
    margin: 0;
    color: #666;
  }
</style>
